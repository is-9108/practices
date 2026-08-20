// このファイルは「採点表」です。編集しないでください。
// ただし読むのは推奨です。
package lesson02

import (
	"maps"
	"slices"
	"testing"
)

func TestAppendSafe(t *testing.T) {
	tests := []struct {
		name string
		base []int
		v    int
		want []int
	}{
		{name: "通常", base: []int{1, 2, 3}, v: 4, want: []int{1, 2, 3, 4}},
		{name: "1要素", base: []int{1}, v: 2, want: []int{1, 2}},
		{name: "空スライス", base: []int{}, v: 1, want: []int{1}},
		{name: "nil スライス", base: nil, v: 1, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendSafe(tt.base, tt.v)
			if !slices.Equal(got, tt.want) {
				t.Errorf("AppendSafe(%v, %d) = %v, want %v", tt.base, tt.v, got, tt.want)
			}
		})
	}

	// ここからが本題。容量に余りがある base を渡して、元の配列が無事かを見ます。
	t.Run("元の配列を壊さない", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		base := arr[:2] // len=2, cap=5 ← 容量が余っているのが罠

		got := AppendSafe(base, 99)

		if !slices.Equal(got, []int{1, 2, 99}) {
			t.Errorf("AppendSafe(%v, 99) = %v, want [1 2 99]", base, got)
		}
		if !slices.Equal(arr, []int{1, 2, 3, 4, 5}) {
			t.Errorf("AppendSafe が元の配列を書き換えました: arr = %v, want [1 2 3 4 5]\n"+
				"        base は len=2 cap=5 です。素の append は余った容量にそのまま書き込むので arr[2] が 99 になります。\n"+
				"        base を壊さないためには、先に新しい配列へコピーする必要があります。", arr)
		}
		if len(base) != 2 {
			t.Errorf("AppendSafe が base の長さを変えました: len(base) = %d, want 2", len(base))
		}
	})

	t.Run("戻り値と base が配列を共有していない", func(t *testing.T) {
		base := make([]int, 2, 10) // len=2, cap=10
		base[0], base[1] = 1, 2

		got := AppendSafe(base, 3)
		got[0] = 777 // 戻り値をいじる

		if base[0] != 1 {
			t.Errorf("戻り値を書き換えたら base まで変わりました: base[0] = %d, want 1\n"+
				"        戻り値は base とは別の配列を指している必要があります。", base[0])
		}
	})
}

func TestRemoveAt(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		i    int
		want []int
	}{
		{name: "真ん中", s: []int{1, 2, 3}, i: 1, want: []int{1, 3}},
		{name: "先頭", s: []int{1, 2, 3}, i: 0, want: []int{2, 3}},
		{name: "末尾", s: []int{1, 2, 3}, i: 2, want: []int{1, 2}},
		{name: "1要素", s: []int{1}, i: 0, want: []int{}},
		{name: "範囲外(上)", s: []int{1, 2, 3}, i: 3, want: []int{1, 2, 3}},
		{name: "範囲外(負)", s: []int{1, 2, 3}, i: -1, want: []int{1, 2, 3}},
		{name: "空スライス", s: []int{}, i: 0, want: []int{}},
		{name: "nil スライス", s: nil, i: 0, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.s) // 失敗メッセージ用に元の値を控えておく
			got := RemoveAt(tt.s, tt.i)
			if !slices.Equal(got, tt.want) {
				t.Errorf("RemoveAt(%v, %d) = %v, want %v", input, tt.i, got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		size int
		want [][]int
	}{
		{name: "割り切れない", s: []int{1, 2, 3, 4, 5}, size: 2, want: [][]int{{1, 2}, {3, 4}, {5}}},
		{name: "割り切れる", s: []int{1, 2, 3, 4}, size: 2, want: [][]int{{1, 2}, {3, 4}}},
		{name: "size が len より大きい", s: []int{1, 2, 3}, size: 5, want: [][]int{{1, 2, 3}}},
		{name: "size が 1", s: []int{1, 2}, size: 1, want: [][]int{{1}, {2}}},
		{name: "空スライス", s: []int{}, size: 2, want: [][]int{}},
		{name: "nil スライス", s: nil, size: 2, want: nil},
		{name: "size が 0", s: []int{1, 2, 3}, size: 0, want: nil},
		{name: "size が負", s: []int{1, 2, 3}, size: -1, want: nil},
	}

	equal := func(a, b [][]int) bool {
		return slices.EqualFunc(a, b, func(x, y []int) bool { return slices.Equal(x, y) })
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.s, tt.size)
			if !equal(got, tt.want) {
				t.Errorf("Chunk(%v, %d) = %v, want %v", tt.s, tt.size, got, tt.want)
			}
		})
	}

	t.Run("チャンクの容量が切り詰められている", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6}
		got := Chunk(s, 2)
		if len(got) != 3 {
			t.Fatalf("Chunk(%v, 2) の要素数 = %d, want 3", s, len(got))
		}

		for i, c := range got {
			if cap(c) != len(c) {
				t.Errorf("Chunk(%v, 2) の %d 番目のチャンク: len = %d, cap = %d (len と同じであるべき)\n"+
					"        s[low:high] だと cap が元スライスの末尾まで伸びてしまいます。\n"+
					"        3 インデックスのスライス式 s[low:high:high] で cap を切り詰めてください。",
					s, i, len(c), cap(c))
			}
		}

		// 容量が伸びていると、先頭チャンクへの append が次のチャンクを踏み潰します。
		_ = append(got[0], 99)
		if got[1][0] != 3 {
			t.Errorf("先頭チャンクに append したら隣のチャンクが壊れました: got[1] = %v, want [3 4]", got[1])
		}
	})
}

func TestDedupe(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		want []string
	}{
		{name: "通常", s: []string{"a", "b", "a", "c", "b"}, want: []string{"a", "b", "c"}},
		{name: "全部同じ", s: []string{"a", "a", "a"}, want: []string{"a"}},
		{name: "重複なし", s: []string{"a", "b", "c"}, want: []string{"a", "b", "c"}},
		{name: "大文字小文字を区別", s: []string{"Go", "go", "Go"}, want: []string{"Go", "go"}},
		{name: "空文字列も要素", s: []string{"", "a", ""}, want: []string{"", "a"}},
		{name: "空スライス", s: []string{}, want: []string{}},
		{name: "nil スライス", s: nil, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dedupe(tt.s)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Dedupe(%q) = %q, want %q (並び順は最初に現れた順)", tt.s, got, tt.want)
			}
		})
	}
}

func TestGroupByLength(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[int][]string
	}{
		{
			name:  "通常",
			words: []string{"go", "rust", "c", "java", "x"},
			want:  map[int][]string{2: {"go"}, 4: {"rust", "java"}, 1: {"c", "x"}},
		},
		{
			name:  "日本語は rune 単位",
			words: []string{"あい", "go", "こんにちは"},
			want:  map[int][]string{2: {"あい", "go"}, 5: {"こんにちは"}},
		},
		{name: "空文字列", words: []string{""}, want: map[int][]string{0: {""}}},
		{name: "空スライス", words: []string{}, want: map[int][]string{}},
		{name: "nil スライス", words: nil, want: map[int][]string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupByLength(tt.words)
			if got == nil {
				t.Fatalf("GroupByLength(%q) = nil, want %v (nil ではなく空の map を返してください)", tt.words, tt.want)
			}
			ok := maps.EqualFunc(got, tt.want, func(a, b []string) bool { return slices.Equal(a, b) })
			if !ok {
				t.Errorf("GroupByLength(%q) = %v, want %v (グループ内は入力順)", tt.words, got, tt.want)
			}
		})
	}
}

func TestSortedKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want []string
	}{
		{name: "通常", m: map[string]int{"banana": 2, "apple": 1, "cherry": 3}, want: []string{"apple", "banana", "cherry"}},
		{name: "1要素", m: map[string]int{"go": 1}, want: []string{"go"}},
		{name: "空 map", m: map[string]int{}, want: []string{}},
		{name: "nil map", m: nil, want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedKeys(tt.m)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SortedKeys(%v) = %q, want %q", tt.m, got, tt.want)
			}
		})
	}

	// map の反復順序はランダムなので、ソートを忘れるとここで落ちます。
	t.Run("何度呼んでも同じ順序", func(t *testing.T) {
		m := map[string]int{"e": 1, "d": 2, "c": 3, "b": 4, "a": 5, "f": 6, "g": 7, "h": 8}
		want := SortedKeys(m)
		for i := 0; i < 100; i++ {
			if got := SortedKeys(m); !slices.Equal(got, want) {
				t.Fatalf("SortedKeys を繰り返し呼んだら結果が変わりました (%d 回目): got = %q, 初回 = %q\n"+
					"        Go の map は反復順序が毎回ランダムです。キーを集めたあとソートしてください。", i+2, got, want)
			}
		}
	})
}

func TestMergeCounts(t *testing.T) {
	t.Run("dst に足し込まれる", func(t *testing.T) {
		dst := map[string]int{"a": 1, "b": 2}
		src := map[string]int{"b": 3, "c": 4}

		MergeCounts(dst, src)

		want := map[string]int{"a": 1, "b": 5, "c": 4}
		if !maps.Equal(dst, want) {
			t.Errorf("MergeCounts 後の dst = %v, want %v\n"+
				"        map は参照のように振る舞うので、戻り値なしでも呼び出し側の dst が変わります。", dst, want)
		}
	})

	t.Run("src は書き換えない", func(t *testing.T) {
		dst := map[string]int{"a": 1}
		src := map[string]int{"a": 10, "b": 20}
		srcWant := maps.Clone(src)

		MergeCounts(dst, src)

		if !maps.Equal(src, srcWant) {
			t.Errorf("MergeCounts が src を書き換えました: src = %v, want %v", src, srcWant)
		}
	})

	t.Run("dst が nil でも panic しない", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("MergeCounts(nil, src) が panic しました: %v\n"+
					"        nil map への書き込みは panic します。先に dst == nil をガードしてください。", r)
			}
		}()
		MergeCounts(nil, map[string]int{"a": 1})
	})

	t.Run("src が nil でも panic しない", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("MergeCounts(dst, nil) が panic しました: %v", r)
			}
		}()
		dst := map[string]int{"a": 1}
		MergeCounts(dst, nil)
		if !maps.Equal(dst, map[string]int{"a": 1}) {
			t.Errorf("src が nil なのに dst が変わりました: dst = %v, want map[a:1]", dst)
		}
	})

	t.Run("両方 nil でも panic しない", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("MergeCounts(nil, nil) が panic しました: %v", r)
			}
		}()
		MergeCounts(nil, nil)
	})
}
