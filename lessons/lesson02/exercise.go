// Package lesson02 は slice と map の内部構造・落とし穴を学ぶための演習です。
// 各関数の panic("TODO") を自分の実装に置き換えてください。
package lesson02

// 課題1: AppendSafe は base に v を足した「新しいスライス」を返します。
// base 自身と、base が参照している配列を一切書き換えてはいけません。
//
// ポイント: これが append の最大の罠です。
// base に容量の余りがあると、append は同じ配列に書き込むので、
// base の「隣」を共有している別のスライスまで巻き添えで壊れます。
func AppendSafe(base []int, v int) []int {
	panic("TODO: AppendSafe を実装してください")
}

// 課題2: RemoveAt は s から i 番目の要素を取り除いたスライスを返します。
// i が範囲外（i < 0 または i >= len(s)）のときは s をそのまま返してください。
//
// 課題1とは逆に、こちらは s の配列を書き換えて構いません（実務で使うイディオム）。
//
// ポイント: append(s[:i], s[i+1:]...) が定番の書き方です。
// 可変長引数に「スライスを展開して渡す」記法 slice... を使います。
func RemoveAt(s []int, i int) []int {
	panic("TODO: RemoveAt を実装してください")
}

// 課題3: Chunk は s を size 個ずつのまとまりに分割します。
// 例: Chunk([]int{1,2,3,4,5}, 2) -> [][]int{{1,2},{3,4},{5}}
// 最後の端数はそのまま入れてください。
// size が 0 以下のときは nil を返してください。
//
// 追加の要求: 各チャンクは「3 インデックスのスライス式」s[low:high:max] を使って
// 容量を len と同じまで切り詰めてください。
// そうしないと、あるチャンクに append したときに次のチャンクの中身が壊れます。
func Chunk(s []int, size int) [][]int {
	panic("TODO: Chunk を実装してください")
}

// 課題4: Dedupe は重複を取り除きます。並び順は「最初に現れた順」を保ってください。
// 例: []string{"a","b","a","c"} -> []string{"a","b","c"}
//
// ポイント: Go に set 型はないので map[string]struct{} で代用します。
// struct{} は「サイズ 0 の型」で、値を持たない集合を表すときの慣習です。
// 存在チェックは v, ok := m[k] の ok の方を使います。
func Dedupe(s []string) []string {
	panic("TODO: Dedupe を実装してください")
}

// 課題5: GroupByLength は単語を「文字数」でグループ分けします。
// 例: []string{"go","c","rust"} -> map[int][]string{2:{"go"}, 1:{"c"}, 4:{"rust"}}
// 各グループの中の並び順は、入力の順番を保ってください。
// words が空のときは、空の map（nil ではない）を返してください。
//
// 文字数は rune 単位で数えてください（lesson01 の Reverse と同じ話です）。
//
// ポイント: map の値がスライスのとき、キーが未登録でもゼロ値の nil スライスが返り、
// nil スライスには append できます。つまり存在チェックなしで
// m[k] = append(m[k], v) と書けます。
func GroupByLength(words []string) map[int][]string {
	panic("TODO: GroupByLength を実装してください")
}

// 課題6: SortedKeys は m のキーを昇順に並べて返します。
// m が nil または空のときは、長さ 0 のスライスを返してください。
//
// ポイント: Go の map は反復順序が「毎回ランダム」です（意図的な仕様）。
// 出力を安定させたいときは必ずソートします。
// slices.Sort や maps.Keys が使えます。
func SortedKeys(m map[string]int) []string {
	panic("TODO: SortedKeys を実装してください")
}

// 課題7: MergeCounts は src の集計を dst に足し込みます。戻り値はありません。
// 呼び出し側が渡した dst が書き換わります。
// dst が nil のときは何もせずに戻ってください（panic させないこと）。
// src が nil のときも何もしません。
//
// ポイント: map は「参照のように振る舞う」型なので、関数に渡した map への書き込みは
// 呼び出し側にも見えます（スライスの append とは挙動が違います）。
// ただし nil map への書き込みは panic するので、ガードが要ります。
func MergeCounts(dst, src map[string]int) {
	panic("TODO: MergeCounts を実装してください")
}
