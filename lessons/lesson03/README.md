# Lesson 03: struct とメソッド

Go にはクラスも継承もありません。代わりに **struct + メソッド + 埋め込み** で組み立てます。
このレッスンで「Go でのオブジェクトの作り方」を身につけます。Phase 3 以降で書く
`User`、`Repository`、`Handler` はすべてこの回の応用です。

## 1. struct とフィールドの可視性

```go
type User struct {
	ID    int      // 大文字始まり → 他パッケージから見える (public)
	Name  string
	email string   // 小文字始まり → 同じパッケージ内だけ (private)
}
```

Go の可視性は **識別子の先頭が大文字かどうかだけ**で決まります。
`public` / `private` というキーワードはありません。単位は「クラス」ではなく **パッケージ**です。
同じパッケージ内なら、小文字フィールドも別の型から自由に触れます。

## 2. 生成の書き方

```go
u := User{ID: 1, Name: "太郎"}   // フィールド名つき（必ずこちらを使う）
u := User{1, "太郎", ""}         // 位置指定。フィールドが増えると壊れるので実務では使わない

p := &User{ID: 1}               // ポインタが欲しいときはこれ
p := new(User)                  // &User{} と同じ。あまり使わない

var u User                      // ゼロ値: ID=0, Name="", email=""
```

**Go にコンストラクタ構文はありません。** 慣習として `NewXxx` 関数を書きます。

```go
func NewUser(id int, name string) (*User, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	return &User{ID: id, Name: name}, nil
}
```

`NewXxx` を用意する理由は **「不正な状態のインスタンスを作らせない」** ためです。
検証をコンストラクタに集約すると、以降のコードは「User があるなら必ず妥当」と信じられます。

## 3. メソッド

```go
func (u User) DisplayName() string { return u.Name }
//    ^^^^^^ レシーバ
```

- クラスの中ではなく、**型の外側**に書きます。同じパッケージ内であればファイルはどこでも構いません
- レシーバ名は `this` や `self` ではなく、**型名の頭文字1〜2文字**が慣習です（`u`, `c`, `srv`）
- メソッドを定義できるのは **自分のパッケージで定義した型だけ**です（`int` や `time.Time` には生やせない）

## 4. 値レシーバ vs ポインタレシーバ（この回の核心）

```go
func (c Counter) IncValue()  { c.n++ }   // c は「コピー」。呼び出し側は変わらない
func (c *Counter) IncPtr()   { c.n++ }   // 呼び出し側の実体が変わる
```

```go
var c Counter
c.IncValue()   // 何も起きない（コンパイルは通る。これが厄介）
c.Value()      // 0
c.IncPtr()
c.Value()      // 1
```

**状態を変えるメソッドは必ずポインタレシーバ。** 値レシーバで書いても
コンパイルエラーにならず、静かにバグります。

### どちらを選ぶか

| 状況 | レシーバ |
|---|---|
| フィールドを書き換える | **ポインタ** |
| struct が大きい（コピーコストが高い） | **ポインタ** |
| `sync.Mutex` などコピー禁止のものを含む | **ポインタ** |
| 小さくて不変な値型（`time.Time` のような） | 値 |

**迷ったらポインタ**で問題ありません。

そして重要なルールが 1 つ。

> **同じ型のメソッドは、レシーバをポインタか値のどちらかに統一する。**

1つでもポインタが必要なら、全部ポインタに揃えます。混在させると、後で interface を
満たすかどうか（メソッドセットの話）で事故ります。この演習でもそう揃えています。

### なぜ `c.IncPtr()` が書けるのか

`c` は `Counter` であって `*Counter` ではないのに、`c.IncPtr()` と書けます。
これは Go が自動的に `(&c).IncPtr()` に読み替えてくれるからです。

ただし **変数のようにアドレスを取れるもの（addressable）に限ります**。

```go
m := map[string]Counter{"a": {}}
m["a"].IncPtr()      // コンパイルエラー: map の要素はアドレスを取れない
newCounter().IncPtr() // コンパイルエラー: 関数の戻り値はアドレスを取れない

c := m["a"]; c.IncPtr(); m["a"] = c   // 一度取り出して書き戻す（lesson02 と同じ話）
```

## 5. ゼロ値を使えるように設計する

Go では **`var x T` がそのまま動く型**が良い設計とされています。

```go
var s Stack     // 初期化不要で Push できる
var buf bytes.Buffer   // 標準ライブラリの例。いきなり buf.WriteString(...) が書ける
var mu sync.Mutex      // いきなり mu.Lock() が書ける
```

`Stack` の中身が `items []int` なら、nil スライスに `append` できるので初期化が要りません
（lesson02 でやった話です）。逆に **map フィールドを持つ型はゼロ値で書き込めない**ので、
その型は `NewXxx` が必須になります。この判断ができると Go らしい設計になります。

## 6. 埋め込み（embedding）— 継承ではない

フィールド名を書かずに型だけ書くと「埋め込み」になります。

```go
type Base struct {
	ID int
}
func (b Base) Label() string { return fmt.Sprintf("base#%d", b.ID) }

type Article struct {
	Base            // ← フィールド名がない = 埋め込み
	Title string
}

a := Article{Base: Base{ID: 7}, Title: "Go入門"}
a.ID        // 7          ← フィールドが「昇格」して直接触れる
a.Label()   // "base#7"   ← メソッドも昇格する
a.Base.ID   // 7          ← 明示的にも書ける
```

見た目は継承ですが、実体は **「Base という名前のフィールドを持っている」だけ**です（has-a）。
`Article` は `Base` の一種ではありません。

同名のメソッドを外側に定義すると、**外側が優先**されます（オーバーライドのように見える）。
埋め込み側は `a.Base.Label()` で明示的に呼べます。

```go
func (a Article) Label() string {
	return a.Base.Label() + " / " + a.Title   // "base#7 / Go入門"
}
```

> 注意: 埋め込みは「共通フィールドをまとめる」用途では便利ですが、
> 継承のつもりで深い階層を作ると破綻します。Go では **interface で振る舞いを抽象化**するのが基本です（lesson04）。

## 7. struct の比較とコピー

```go
a := Point{1, 2}
b := a          // 全フィールドがコピーされる
b.X = 99        // a には影響しない

a == Point{1, 2}   // true。== で比較できる
```

struct は **全フィールドが比較可能なら `==` で比較できます**（配列も可）。
ただし **slice / map / 関数を含む struct は比較不可**でコンパイルエラーになります。

そして最大の落とし穴。

```go
type Profile struct {
	Name string
	Tags []string
}

a := Profile{Name: "太郎", Tags: []string{"go"}}
b := a            // struct はコピーされるが……
b.Tags[0] = "rust"
a.Tags[0]         // "rust"  ← 元まで変わる
```

**struct のコピーは浅いコピー（shallow copy）です。**
slice / map / ポインタのフィールドは「参照先」が共有されます。
完全な複製が欲しいなら `Clone` メソッドを自分で書きます（課題6）。

## 8. `String()` を実装すると出力が変わる

```go
func (m Money) String() string { return fmt.Sprintf("%d %s", m.amount, m.currency) }

fmt.Println(m)              // "1050 JPY"  ← String() が自動的に使われる
fmt.Sprintf("%v", m)        // "1050 JPY"
```

`String() string` を持つ型は `fmt.Stringer` を満たし、`fmt` パッケージが自動で使います。
これが Go の **暗黙的なインターフェース実装**で、lesson04 の主題です。

> 罠: `String()` の中で `fmt.Sprintf("%v", m)` と書くと、自分自身を呼び続けて無限再帰します。
> 必ずフィールドを個別に展開してください（`go vet` が検出してくれます）。

---

## 課題

`exercise.go` の 6 つの型を完成させてください。

| # | 型 | 学ぶこと |
|---|------|---------|
| 1 | `Counter` | ポインタレシーバでないと状態が変わらない（**レシーバの修正が必要な箇所があります**） |
| 2 | `User` | コンストラクタでの検証、`*T` を返す、ポインタレシーバでの更新 |
| 3 | `Stack` | ゼロ値のまま使える型の設計 |
| 4 | `Money` | 値レシーバの不変型、`==` での比較、`fmt.Stringer` |
| 5 | `Base` / `Article` | 埋め込み、フィールドとメソッドの昇格、明示的な呼び分け |
| 6 | `Profile` | 浅いコピーの罠、`Clone` の実装 |

課題1は **stub のレシーバがわざと間違っています**。テストが落ちたら、まずレシーバを疑ってください。

### 答え合わせ

```bash
go test ./lessons/lesson03/
```

型ごとに絞るなら:

```bash
go test -run TestMoney -v ./lessons/lesson03/
```

### 使いそうな標準ライブラリ

- `strings.TrimSpace(s)` — 前後の空白を除去
- `strings.Contains(s, substr)` / `strings.HasPrefix` / `strings.HasSuffix`
- `fmt.Sprintf(format, ...)` — 文字列組み立て
- `slices.Clone(s)` / `maps.Clone(m)` — nil を渡すと nil が返る（課題6でそこが効きます）
- `time.Time.Sub(t)` — 時刻の差を `time.Duration` で得る

### 詰まったら

「lesson03 の N が分からない」と聞いてください。ヒント → 解説 → 模範解答の順で出します。
