# Lesson 02: slice と map の深掘り

lesson01 では slice と map を「使えるもの」として扱いました。
このレッスンでは **中身がどうなっているか** を見ます。ここを知らないと、
Go では「動いているように見えて、実は隣のデータを壊している」バグを踏みます。

実務で最も踏まれる罠が詰まっている回です。

## 1. slice は「配列へのビュー」

Go の slice は、実体としては **3 つのフィールドを持つ小さな構造体**です。

```
slice ─┬─ ptr : 実データが入っている「配列」の先頭アドレス
       ├─ len : 今見えている長さ           （len(s) で取れる）
       └─ cap : ptr から配列の末尾までの長さ（cap(s) で取れる）
```

slice 自体はデータを持っておらず、**別にある配列を指しているだけ**です。
関数に渡すとこの構造体がコピーされますが、`ptr` の指す先は同じなので、
**要素の書き換えは呼び出し側にも見えます**。

```go
func zero(s []int) { s[0] = 0 }          // 反映される（ptr が同じ配列を指している）
func grow(s []int) { s = append(s, 1) }  // 反映されない（len はコピーされた構造体の側だけ変わる）
```

「要素の変更は伝わるが、長さの変更は伝わらない」——だから `append` は
**必ず戻り値を受け取る**必要があります。

## 2. スライス式は配列を共有する

```go
arr := []int{1, 2, 3, 4, 5}
s := arr[1:3]      // s = [2 3]

len(s)  // 2
cap(s)  // 4  ← low(1) から arr の末尾まで。high は cap に影響しない

s[0] = 99
arr     // [1 99 3 4 5]  ← arr まで変わる
```

`arr[1:3]` は **コピーではありません**。同じ配列の一部を見ているだけです。

## 3. append の罠（このレッスンの核心）

`append` の挙動は cap に余りがあるかどうかで **2 通りに分岐**します。

```go
arr := []int{1, 2, 3, 4, 5}
s := arr[:2]        // len=2, cap=5 ← 余りが 3 ある

s = append(s, 99)   // 余りがあるので arr[2] にそのまま書き込む(!)
arr                 // [1 2 99 4 5]  ← 元の配列が壊れた
```

```go
s := []int{1, 2}    // len=2, cap=2 ← 余りなし
t := append(s, 3)   // 新しい配列を確保してコピー
t[0] = 99
s                   // [1 2]  ← こちらは無傷
```

**同じ `append` なのに、cap 次第で「元を壊す / 壊さない」が変わる。**
Go で最も再現しづらいバグの温床です。
「引数で受け取った slice に append して返す」関数は、この危険を常に抱えています。

対策は 2 つ。

### 対策A: 先にコピーする

```go
out := make([]int, len(s), len(s)+1)
copy(out, s)              // copy は「短い方の長さ」だけコピーし、コピーした個数を返す
out = append(out, v)

// slices.Clone(s) でも同じ（Go 1.21+）
```

### 対策B: cap を切り詰める（3インデックスのスライス式）

```go
s := arr[low:high:max]    // len = high-low, cap = max-low
s := arr[1:3:3]           // len=2, cap=2 ← 余りゼロ
```

cap に余りが無ければ `append` は必ず新しい配列を確保します。
つまり **「これ以上ここには書き込ませない」という宣言**になります。
slice を外部に渡す API では、この形にしておくと事故が減ります。

## 4. 削除のイディオム

Go に `remove` はありません。慣習はこれです。

```go
s = append(s[:i], s[i+1:]...)   // i 番目を削除
```

`slice...` は「可変長引数にスライスを展開して渡す」記法です。
これは上書きでずらすので、元の配列を破壊的に書き換えます。壊したくないなら先にコピー。

順序を保たなくていいなら、末尾と入れ替えて切り詰める方が速いです。

```go
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

> 標準の `slices` パッケージに `slices.Delete(s, i, i+1)` があります。実務ではこちらで十分ですが、
> 中で何が起きているかは知っておいてください。

## 5. nil スライス vs 空スライス

```go
var a []int        // nil スライス   a == nil → true
b := []int{}       // 空スライス     b == nil → false
```

**どちらも `len` は 0、`range` は 0 回、`append` はできる。** 実用上の差はほぼありません。
Go では **nil スライスを積極的に使う**のが慣習です（`var a []int` のまま append できる）。

差が出るのは `encoding/json` です。nil は `null`、空スライスは `[]` になります。
API のレスポンスでは `[]` を返したいことが多いので、そこだけ意識してください（Phase 3 で扱います）。

判定は `s == nil` ではなく **`len(s) == 0`** で書くのが安全です。

## 6. map は nil だと書けない

```go
var m map[string]int   // nil map
m["a"]                 // 0   ← 読むのは OK
len(m)                 // 0   ← OK
for k := range m {}    // OK（0 回）
m["a"] = 1             // panic: assignment to entry in nil map  ← これだけ死ぬ
```

読み取りは安全なのに書き込みだけ panic するので、テストをすり抜けやすい罠です。
**map は必ず `make` か `{}` で初期化**してください。

また map は slice と違い、**関数に渡した先での追加・削除・更新が呼び出し側にも見えます**
（内部的にハッシュテーブルへのポインタを持っているため）。`append` のように戻り値を受け直す作法は不要です。

## 7. map の反復順序はランダム

```go
for k, v := range m { ... }   // 順序は毎回変わる
```

これは実装の都合ではなく **意図的な仕様**です。順序に依存したコードを書かせないため、
Go はわざと毎回ランダム化しています。

出力を安定させたいなら、キーを集めてソートします。

```go
keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
slices.Sort(keys)

// Go 1.23+ なら 1 行
keys := slices.Sorted(maps.Keys(m))
```

「ローカルでは通るのに CI でたまに落ちる」テストの原因 No.1 がこれです。

## 8. map の値には直接アクセスできない

```go
type Point struct{ X, Y int }

m := map[string]Point{"a": {1, 2}}
m["a"].X = 10     // コンパイルエラー: cannot assign
&m["a"]           // コンパイルエラー: cannot take address
```

map の中身はリハッシュで移動するので、アドレスを取らせてもらえません。
書き換えたいときは **一度取り出して書き戻す**か、**値をポインタにする**。

```go
p := m["a"]; p.X = 10; m["a"] = p   // 取り出して書き戻す
m := map[string]*Point{}            // ポインタなら m["a"].X = 10 が書ける
```

一方で **値がスライスの場合は `append` が使えます**。

```go
m := map[int][]string{}
m[3] = append(m[3], "abc")   // m[3] は未登録 → nil スライスが返る → append できる
```

存在チェックが要らないのがポイントです。課題5で使います。

## 9. set は `map[T]struct{}`

Go に set 型はありません。慣習はこれです。

```go
seen := map[string]struct{}{}
seen["go"] = struct{}{}

if _, ok := seen["go"]; ok { ... }   // 存在チェックは 2 値目の ok を見る
```

`struct{}` は **サイズ 0** の型なので、値の分のメモリを一切使いません。
`map[string]bool` でも動きますが、「値に意味がない」ことを型で示すのが Go らしい書き方です。

---

## 課題

`exercise.go` の 7 つの関数を実装してください。

| # | 関数 | 学ぶこと |
|---|------|---------|
| 1 | `AppendSafe` | append の罠、`copy` / `slices.Clone` |
| 2 | `RemoveAt` | 削除イディオム、`slice...` 展開 |
| 3 | `Chunk` | 3インデックスのスライス式、cap の切り詰め |
| 4 | `Dedupe` | `map[T]struct{}` を set として使う、順序の保持 |
| 5 | `GroupByLength` | map の値がスライス、nil スライスへの append |
| 6 | `SortedKeys` | map の反復順序はランダム、`slices.Sort` |
| 7 | `MergeCounts` | map は参照的に振る舞う、nil map への書き込み |

課題1と課題3は「実装は通るのにテストが落ちる」ことがあります。
そのときこそ本題なので、失敗メッセージをよく読んでください。

### 答え合わせ

```bash
go test ./lessons/lesson02/
```

失敗したケースだけが表示されます。全ケースを見たいときだけ `-v` を付けてください。
1 つの関数に絞るなら:

```bash
go test -run TestChunk -v ./lessons/lesson02/
```

### 使いそうな標準ライブラリ

- `copy(dst, src)` — 組み込み関数。短い方の長さだけコピーし、コピーした個数を返す
- `slices.Clone(s)` — スライスの複製（nil を渡すと nil が返る）
- `slices.Sort(s)` — 昇順ソート（破壊的）
- `maps.Keys(m)` — キーの iterator（Go 1.23+）。`slices.Sorted(maps.Keys(m))` で一発
- `utf8.RuneCountInString(s)` — rune 単位の文字数。`len([]rune(s))` でも可
- `min(a, b)` / `max(a, b)` — 組み込み関数（Go 1.21+）。`Chunk` の端数処理で便利

### 詰まったら

「lesson02 の N が分からない」と聞いてください。ヒント → 解説 → 模範解答の順で出します。
