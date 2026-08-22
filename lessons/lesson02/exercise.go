// Package lesson02 は slice と map の内部構造、および実務で踏みやすい落とし穴を学ぶ演習です。
// 各関数の panic("TODO") を自分の実装に置き換えてください。
package lesson02

import "slices"

// 課題1: AppendSafe は base に v を足した「新しいスライス」を返します。
// base 自身と、base が参照している配列を一切書き換えてはいけません。
//
// ポイント: これが append 最大の罠です。
// base に容量(cap)の余りがあると、append は同じ配列にそのまま書き込むため、
// その配列を共有している別のスライスまで巻き添えで壊れます。
// 先に新しい配列へコピーしてから append してください。
func AppendSafe(base []int, v int) []int {
	// panic("TODO: AppendSafe を実装してください")
	res := slices.Clone(base)
	return append(res, v)
}

// 課題2: RemoveAt は s から i 番目の要素を取り除いたスライスを返します。
// i が範囲外(i < 0 または i >= len(s))のときは s をそのまま返してください。
//
// 課題1とは逆に、こちらは s の配列を書き換えて構いません（実務で使う標準イディオム）。
//
// ポイント: append(s[:i], s[i+1:]...) が定番です。
// slice... は「可変長引数にスライスを展開して渡す」記法です。
func RemoveAt(s []int, i int) []int {
	// panic("TODO: RemoveAt を実装してください")
	if i < 0 || i >= len(s) {
		return s
	}
	result := append(s[:i], s[i+1:]...)
	return result
}

// 課題3: Chunk は s を size 個ずつのまとまりに分割します。
// 例: Chunk([]int{1,2,3,4,5}, 2) -> [][]int{{1,2},{3,4},{5}}
// 最後の端数はそのまま入れてください。
// size が 0 以下のときは nil を返してください。
//
// 追加要求: 各チャンクは「3インデックスのスライス式」 s[low:high:max] を使って、
// cap が len と等しくなるように切り詰めてください。
// これをしないと、あるチャンクに append したときに次のチャンクの中身が壊れます。
func Chunk(s []int, size int) [][]int {
	// panic("TODO: Chunk を実装してください")
	if size <= 0 {
		return nil
	}
	var result [][]int
	for low := 0; low < len(s); low += size {
		high := min(low+size, len(s))
		result = append(result, s[low:high:high])
	}
	return result
}

// 課題4: Dedupe は重複を取り除きます。並び順は「最初に現れた順」を保ってください。
// 例: []string{"a","b","a","c"} -> []string{"a","b","c"}
//
// ポイント: Go に set 型はないので map[string]struct{} で代用します。
// struct{} はサイズ 0 の型で、「値に意味がない集合」を表す慣習です。
// 存在チェックは _, ok := m[k] の ok を使います。
func Dedupe(s []string) []string {
	// panic("TODO: Dedupe を実装してください")
	var res []string
	for _, r := range s {
		if slices.Contains(res, r) {
			continue
		}
		res = append(res, r)
	}
	return res
}

// 課題5: GroupByLength は単語を「文字数」でグループ分けします。
// 例: []string{"go","c","rust"} -> map[int][]string{2:{"go"}, 1:{"c"}, 4:{"rust"}}
// 各グループ内の並び順は入力順を保ってください。
// words が空のときは、nil ではなく空の map を返してください。
//
// 文字数は rune 単位で数えます（lesson01 の Reverse と同じ話です）。
//
// ポイント: map の値がスライスのとき、未登録のキーを読むとゼロ値の nil スライスが返り、
// nil スライスには append できます。つまり存在チェックなしで
// m[k] = append(m[k], v) と書けます。
func GroupByLength(words []string) map[int][]string {
	panic("TODO: GroupByLength を実装してください")
}

// 課題6: SortedKeys は m のキーを昇順に並べて返します。
// m が nil または空のときは、長さ 0 のスライスを返してください。
//
// ポイント: Go の map は反復順序が毎回ランダムです（意図的な仕様）。
// 出力を安定させたいなら必ずソートします。
// slices.Sort、または slices.Sorted(maps.Keys(m)) が使えます。
func SortedKeys(m map[string]int) []string {
	panic("TODO: SortedKeys を実装してください")
}

// 課題7: MergeCounts は src の集計を dst に足し込みます。戻り値はありません。
// 呼び出し側が渡した dst が書き換わります。src は書き換えないでください。
// dst が nil のときは何もせずに戻ってください（panic させないこと）。
// src が nil のときも何もしません。
//
// ポイント: map は参照のように振る舞うので、関数に渡した map への書き込みは
// 呼び出し側にも見えます（append で戻り値を受け直すスライスとは挙動が違います）。
// ただし nil map への書き込みは panic するので、ガードが必要です。
func MergeCounts(dst, src map[string]int) {
	panic("TODO: MergeCounts を実装してください")
}
