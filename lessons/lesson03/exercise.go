// Package lesson03 は struct とメソッドを学ぶための演習です。
// 各メソッドの panic("TODO") を自分の実装に置き換えてください。
package lesson03

import (
	"errors"
	"time"
)

// ============================================================
// 課題1: Counter — ポインタレシーバ
// ============================================================

// Counter はカウンタです。
// n は小文字始まりなので、外部パッケージからは直接触れません（メソッド経由のみ）。
type Counter struct {
	n int
}

// Inc はカウンタを 1 増やします。
//
// ⚠ このメソッドは「レシーバがわざと間違っています」。
// 値レシーバはコピーを受け取るだけなので、呼び出し側の Counter は変わりません。
// レシーバごと修正してください。
func (c Counter) Inc() {
	panic("TODO: Inc を実装してください（レシーバも見直すこと）")
}

// Add はカウンタに delta を足します。delta は負でも構いません。
func (c *Counter) Add(delta int) {
	panic("TODO: Add を実装してください")
}

// Value は現在の値を返します。
func (c *Counter) Value() int {
	panic("TODO: Value を実装してください")
}

// Reset はカウンタを 0 に戻します。
func (c *Counter) Reset() {
	panic("TODO: Reset を実装してください")
}

// ============================================================
// 課題2: User — コンストラクタでの検証
// ============================================================

var (
	// ErrInvalidID は ID が 0 以下のときに返します。
	ErrInvalidID = errors.New("invalid id")
	// ErrEmptyName は名前が空（空白のみを含む）のときに返します。
	ErrEmptyName = errors.New("empty name")
	// ErrInvalidEmail はメールアドレスの形式が不正なときに返します。
	ErrInvalidEmail = errors.New("invalid email")
)

// User は利用者を表します。
type User struct {
	ID    int
	Name  string
	Email string
}

// NewUser は User を生成します。ポインタで返してください。
//
// 検証は以下の順で行い、最初に見つかった問題のエラーを返します（このとき戻り値は nil, err）。
//  1. id が 0 以下         -> ErrInvalidID
//  2. name が空、または空白のみ -> ErrEmptyName
//  3. email が "@" を含まない、または "@" で始まる／終わる -> ErrInvalidEmail
//
// name は前後の空白を取り除いて保存してください（"  太郎 " -> "太郎"）。
// email はそのまま保存します。
//
// ポイント: 検証をコンストラクタに集約すると、
// 「User が存在するなら必ず妥当」と以降のコードが信じられるようになります。
func NewUser(id int, name, email string) (*User, error) {
	panic("TODO: NewUser を実装してください")
}

// Rename は名前を変更します。
// 検証ルールは NewUser と同じ（空・空白のみは ErrEmptyName、保存時に TrimSpace）。
// エラーのときは u を一切変更しないでください。
func (u *User) Rename(name string) error {
	panic("TODO: Rename を実装してください")
}

// ============================================================
// 課題3: Stack — ゼロ値のまま使える型
// ============================================================

// Stack は int の LIFO スタックです。
// var s Stack と宣言しただけで（NewStack なしで）使えるように実装してください。
type Stack struct {
	items []int
}

// Push は値を積みます。
func (s *Stack) Push(v int) {
	panic("TODO: Push を実装してください")
}

// Pop は一番上の値を取り出して取り除きます。
// 空のときは (0, false) を返してください。
func (s *Stack) Pop() (int, bool) {
	panic("TODO: Pop を実装してください")
}

// Peek は一番上の値を取り除かずに返します。
// 空のときは (0, false) を返してください。
func (s *Stack) Peek() (int, bool) {
	panic("TODO: Peek を実装してください")
}

// Len は積まれている個数を返します。
func (s *Stack) Len() int {
	panic("TODO: Len を実装してください")
}

// ============================================================
// 課題4: Money — 値レシーバの不変型
// ============================================================

var (
	// ErrInvalidCurrency は通貨コードが空のときに返します。
	ErrInvalidCurrency = errors.New("invalid currency")
	// ErrCurrencyMismatch は通貨が異なる Money を演算しようとしたときに返します。
	ErrCurrencyMismatch = errors.New("currency mismatch")
)

// Money は金額です。amount は最小単位（円なら 1 = 1円）で保持します。
// フィールドが両方とも小文字なので、生成後に書き換える手段がありません（不変な値型）。
type Money struct {
	amount   int64
	currency string
}

// NewMoney は Money を生成します。
// currency が空文字列のときは Money{}, ErrInvalidCurrency を返してください。
//
// ポイント: 小さな不変の値型なので、ポインタではなく「値」で返します。
func NewMoney(amount int64, currency string) (Money, error) {
	panic("TODO: NewMoney を実装してください")
}

// Amount は金額を返します。
func (m Money) Amount() int64 {
	panic("TODO: Amount を実装してください")
}

// Currency は通貨コードを返します。
func (m Money) Currency() string {
	panic("TODO: Currency を実装してください")
}

// Add は m と other を足した「新しい Money」を返します。
// m 自身は変更しません（値レシーバなので、そもそも変更できません）。
// 通貨が異なるときは Money{}, ErrCurrencyMismatch を返してください。
func (m Money) Add(other Money) (Money, error) {
	panic("TODO: Add を実装してください")
}

// String は "1050 JPY" の形式で返します（金額、半角スペース、通貨コード）。
//
// ポイント: この String() を持つだけで fmt.Stringer を満たし、
// fmt.Println(m) や %v での出力が自動的にこの形になります。
// 中で fmt.Sprintf("%v", m) と書くと無限再帰するので注意してください。
func (m Money) String() string {
	panic("TODO: String を実装してください")
}

// ============================================================
// 課題5: Base / Article — 埋め込み
// ============================================================

// Base は ID と作成日時を持つ共通部分です。
type Base struct {
	ID        int
	CreatedAt time.Time
}

// Age は now 時点での経過時間を返します。
func (b Base) Age(now time.Time) time.Duration {
	panic("TODO: Age を実装してください")
}

// Label は "base#<ID>" の形式で返します（例: ID が 7 なら "base#7"）。
func (b Base) Label() string {
	panic("TODO: Base.Label を実装してください")
}

// Article は Base を埋め込んだ記事です。
// Base にフィールド名が付いていないことに注目してください（これが埋め込み）。
type Article struct {
	Base
	Title  string
	Author string
}

// NewArticle は Article を生成します。
//
// ポイント: 埋め込みフィールドは Base: Base{...} という名前で初期化します。
func NewArticle(id int, title, author string, createdAt time.Time) Article {
	panic("TODO: NewArticle を実装してください")
}

// Label は "<Base の Label> / <Title>" を返します（例: "base#7 / Go入門"）。
// Base 側の Label と同名なので、外側のこちらが優先されます。
//
// ポイント: 埋め込み側の実装は a.Base.Label() で明示的に呼べます。
// ここでは必ずそれを使って組み立ててください（自前で "base#..." を書かないこと）。
func (a Article) Label() string {
	panic("TODO: Article.Label を実装してください")
}

// ============================================================
// 課題6: Profile — 浅いコピーの罠
// ============================================================

// Profile はプロフィールです。slice と map のフィールドを持ちます。
type Profile struct {
	Name string
	Tags []string
	Meta map[string]string
}

// Clone は p の完全な複製を返します。
// 戻り値の Tags や Meta を書き換えても、p 側が影響を受けてはいけません。
//
// ただし nil は nil のまま返してください（Tags が nil なら戻り値の Tags も nil）。
//
// ポイント: b := a という代入は「浅いコピー」で、slice と map は参照先が共有されます。
// slices.Clone / maps.Clone は nil を渡すと nil を返すので、そのまま要求を満たせます。
func (p Profile) Clone() Profile {
	panic("TODO: Clone を実装してください")
}
