// Package lesson01 は Go の基本文法を学ぶための演習です。
// 各関数の panic("TODO") を自分の実装に置き換えてください。
package lesson01

import "errors"

// ErrDivideByZero は Divide でゼロ除算が起きたときに返すエラーです。
// 実装済みなので、そのまま使ってください。
var ErrDivideByZero = errors.New("division by zero")

// ErrInvalidScore は Grade に 0〜100 の範囲外が渡されたときに返すエラーです。
var ErrInvalidScore = errors.New("invalid score")

// 課題1: SumAndAvg は nums の合計と平均を返します。
// nums が空の場合は sum=0, avg=0 を返してください。
//
// ポイント: Go の関数は複数の値を返せます。ここでは「名前付き戻り値」を使っています。
func SumAndAvg(nums []int) (sum int, avg float64) {
	panic("TODO: SumAndAvg を実装してください")
}

// 課題2: Divide は a を b で割った結果を返します。
// b が 0 の場合は、第2戻り値に ErrDivideByZero を返してください（第1戻り値は 0）。
//
// ポイント: Go には例外がありません。失敗しうる処理は error を「値として」返します。
func Divide(a, b float64) (float64, error) {
	panic("TODO: Divide を実装してください")
}

// 課題3: Grade は点数を評価に変換します。
//
//	90以上 -> "A"
//	80以上 -> "B"
//	70以上 -> "C"
//	60以上 -> "D"
//	それ未満 -> "F"
//
// score が 0 未満または 100 より大きい場合は、"" と ErrInvalidScore を返してください。
//
// ポイント: Go の switch は条件式を省略して `switch { case 条件: }` と書けます。
// また break は不要です（自動で抜けます）。
func Grade(score int) (string, error) {
	panic("TODO: Grade を実装してください")
}

// 課題4: WordCount は空白区切りの単語の出現回数を数えます。
// 例: "go go gopher" -> map[string]int{"go": 2, "gopher": 1}
// 大文字小文字は区別します。空文字列の場合は空の map を返してください。
//
// ポイント: strings.Fields が使えます。
// map から存在しないキーを読むと「ゼロ値」(int なら 0) が返るので、
// 事前の存在チェックなしに m[w]++ と書けます。
func WordCount(s string) map[string]int {
	panic("TODO: WordCount を実装してください")
}

// 課題5: FizzBuzz は 1 から n までの FizzBuzz 結果をスライスで返します。
//
//	3の倍数 -> "Fizz"
//	5の倍数 -> "Buzz"
//	15の倍数 -> "FizzBuzz"
//	それ以外 -> 数値の文字列 ("1", "2", ...)
//
// n が 0 以下の場合は長さ 0 のスライスを返してください。
//
// ポイント: strconv.Itoa で int -> string に変換します。
// (string(i) は文字コード変換になってしまうので間違いです)
// スライスは make([]string, 0, n) で容量を先に確保しておくと効率的です。
func FizzBuzz(n int) []string {
	panic("TODO: FizzBuzz を実装してください")
}

// 課題6: Reverse は文字列を逆順にします。
// 例: "abc" -> "cba", "こんにちは" -> "はちにんこ"
//
// ポイント: Go の string は「UTF-8 のバイト列」です。
// s[i] で取り出すと byte (1バイト) なので、日本語が壊れます。
// []rune(s) に変換すると「文字」単位で扱えます。
func Reverse(s string) string {
	panic("TODO: Reverse を実装してください")
}
