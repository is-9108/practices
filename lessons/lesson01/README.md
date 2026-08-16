# Lesson 01: Go の基本文法

他言語の経験がある前提で、**「Go だとここが違う」**という点に絞って解説します。

## 1. 変数宣言とゼロ値

```go
var a int        // 0     ← 明示的な初期化なしでも「ゼロ値」が入る
var s string     // ""
var m map[string]int // nil ← map と slice と pointer のゼロ値は nil
b := 10          // 型推論つきの短縮宣言。関数の中でのみ使える
```

**未使用の変数・未使用の import はコンパイルエラー**になります。これは警告ではなくエラーです。
最初は驚きますが、デッドコードが残らない仕組みだと思ってください。

## 2. 関数と複数戻り値

```go
func divmod(a, b int) (int, int) {  // 同じ型の引数はまとめて書ける
	return a / b, a % b
}

q, r := divmod(7, 2)
q, _ := divmod(7, 2)   // 使わない戻り値は _ (ブランク識別子) で捨てる
```

戻り値に名前をつけることもできます（今回の `SumAndAvg` がこの形）。

```go
func SumAndAvg(nums []int) (sum int, avg float64) {
	// sum と avg は最初からゼロ値で宣言済みの変数として使える
	return sum, avg
}
```

## 3. エラーハンドリング — Go に例外はない

try/catch はありません。失敗しうる関数は **最後の戻り値に `error`** を返すのが絶対的な慣習です。

```go
f, err := os.Open("data.txt")
if err != nil {
	return err   // 呼び出し側に返す。これを毎回書くのが Go のスタイル
}
```

エラーの作り方:

```go
errors.New("something went wrong")          // 固定メッセージ
fmt.Errorf("user %d not found", id)         // フォーマットあり
```

`errors.New` で作ったエラーは、パッケージレベルの変数（**センチネルエラー**）にしておくと
呼び出し側が `errors.Is(err, ErrDivideByZero)` で判定できます。今回の課題2で使います。

> エラーメッセージは小文字始まり・句読点なしが慣習です（`"division by zero"`）。
> 呼び出し側で `fmt.Errorf("calc: %w", err)` のように連結されるためです。

## 4. if と switch

```go
if v, err := strconv.Atoi(s); err == nil {  // if の中で変数を宣言できる（スコープは if の中だけ）
	fmt.Println(v)
}

switch {                  // 条件式を省略すると if-else if の代わりになる
case score >= 90:
	return "A"
case score >= 80:
	return "B"
}
```

**`break` は不要**です。Go の switch は自動で抜けます（逆に落としたいときだけ `fallthrough`）。

## 5. slice

```go
s := []int{1, 2, 3}
s = append(s, 4)              // append は「新しいスライスを返す」ので代入が必須
s2 := make([]string, 0, 10)   // 長さ0・容量10。要素数が読めるときは容量を渡すと再確保が減る

for i, v := range s {         // i はインデックス、v は値
	fmt.Println(i, v)
}
for _, v := range s { }       // インデックスが不要なら _ で捨てる
```

`nil` スライスに `append` してもOKで、`len(nil スライス)` は 0 です。
「空かどうか」は `len(s) == 0` で判定するのが安全です。

## 6. map

```go
m := map[string]int{}       // または make(map[string]int)
m["go"]++                   // 存在しないキーでもゼロ値 0 から始まるので、これだけで動く

v, ok := m["go"]            // ok は存在したかどうかの bool
delete(m, "go")
```

**`var m map[string]int` (nil map) への書き込みは panic します。** 必ず `make` か `{}` で初期化してください。
読み取りだけなら nil map でも安全（ゼロ値が返る）です。

## 7. string と rune — 最重要の落とし穴

Go の `string` は **UTF-8 のバイト列**です。

```go
s := "あい"
len(s)        // 6 (バイト数!!  文字数ではない)
s[0]          // 227 (byte)  ← 日本語が壊れる

r := []rune(s)
len(r)        // 2 (文字数)
string(r[0])  // "あ"

for i, c := range s {   // range で回すと c は rune (文字) になる
	fmt.Println(i, string(c))   // i はバイト位置なので 0, 3 と飛ぶ
}
```

数値 → 文字列の変換は `strconv.Itoa(i)` を使います。
`string(65)` は `"65"` ではなく `"A"`（文字コード変換）になるので注意。

---

## 課題

`exercise.go` の 6 つの関数を実装してください。

| # | 関数 | 学ぶこと |
|---|------|---------|
| 1 | `SumAndAvg` | 複数戻り値、名前付き戻り値、ゼロ除算の回避 |
| 2 | `Divide` | error を値として返す、センチネルエラー |
| 3 | `Grade` | 条件式なし switch、範囲チェック |
| 4 | `WordCount` | map の初期化とゼロ値、`strings.Fields` |
| 5 | `FizzBuzz` | slice の append、`make` の容量、`strconv.Itoa` |
| 6 | `Reverse` | string と rune の違い |

### 答え合わせ

```bash
go test -v ./lessons/lesson01/
```

全部 PASS すれば Lesson 01 クリアです。

### 使いそうな標準ライブラリ

- `strings.Fields(s)` — 空白（連続・タブ・改行含む）で分割して `[]string` を返す
- `strconv.Itoa(i)` — int を string に変換
- `errors.New(msg)` — エラーを作る（今回は定義済みのものを使うだけ）

import は自分で追加してください。VS Code なら保存時に `goimports` が自動で入れてくれます。
手動なら以下でも整います。

```bash
go mod tidy
```
