// このファイルは「採点表」です。編集しないでください。
// ただし読むのは推奨です。Go で標準的な table-driven test の書き方のお手本になっています。
package lesson01

import (
	"errors"
	"maps"
	"math"
	"slices"
	"testing"
)

func TestSumAndAvg(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantSum int
		wantAvg float64
	}{
		{name: "通常", nums: []int{1, 2, 3, 4}, wantSum: 10, wantAvg: 2.5},
		{name: "1要素", nums: []int{7}, wantSum: 7, wantAvg: 7},
		{name: "負の数を含む", nums: []int{-5, 5, 10}, wantSum: 10, wantAvg: 10.0 / 3.0},
		{name: "空スライス", nums: []int{}, wantSum: 0, wantAvg: 0},
		{name: "nil スライス", nums: nil, wantSum: 0, wantAvg: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotAvg := SumAndAvg(tt.nums)
			if gotSum != tt.wantSum {
				t.Errorf("SumAndAvg(%v) sum = %d, want %d", tt.nums, gotSum, tt.wantSum)
			}
			if math.Abs(gotAvg-tt.wantAvg) > 1e-9 {
				t.Errorf("SumAndAvg(%v) avg = %f, want %f", tt.nums, gotAvg, tt.wantAvg)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr error
	}{
		{name: "割り切れる", a: 10, b: 2, want: 5},
		{name: "小数になる", a: 1, b: 4, want: 0.25},
		{name: "負の数", a: -9, b: 3, want: -3},
		{name: "ゼロ除算", a: 1, b: 0, want: 0, wantErr: ErrDivideByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Divide(%v, %v) error = %v, want %v", tt.a, tt.b, err, tt.wantErr)
			}
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Divide(%v, %v) = %f, want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		name    string
		score   int
		want    string
		wantErr error
	}{
		{name: "満点", score: 100, want: "A"},
		{name: "A の下限", score: 90, want: "A"},
		{name: "B の上限", score: 89, want: "B"},
		{name: "B の下限", score: 80, want: "B"},
		{name: "C", score: 75, want: "C"},
		{name: "D の下限", score: 60, want: "D"},
		{name: "F の上限", score: 59, want: "F"},
		{name: "0点", score: 0, want: "F"},
		{name: "範囲外(下)", score: -1, want: "", wantErr: ErrInvalidScore},
		{name: "範囲外(上)", score: 101, want: "", wantErr: ErrInvalidScore},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Grade(tt.score)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Grade(%d) error = %v, want %v", tt.score, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want map[string]int
	}{
		{name: "通常", s: "go go gopher", want: map[string]int{"go": 2, "gopher": 1}},
		{name: "大文字小文字を区別", s: "Go go", want: map[string]int{"Go": 1, "go": 1}},
		{name: "連続する空白", s: "a   b\tc\nd", want: map[string]int{"a": 1, "b": 1, "c": 1, "d": 1}},
		{name: "空文字列", s: "", want: map[string]int{}},
		{name: "空白のみ", s: "   ", want: map[string]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordCount(tt.s)
			if got == nil {
				t.Fatalf("WordCount(%q) = nil, want %v (nil ではなく空の map を返してください)", tt.s, tt.want)
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("WordCount(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{name: "1から15", n: 15, want: []string{
			"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
			"11", "Fizz", "13", "14", "FizzBuzz",
		}},
		{name: "1のみ", n: 1, want: []string{"1"}},
		{name: "0", n: 0, want: []string{}},
		{name: "負の数", n: -3, want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.n)
			if !slices.Equal(got, tt.want) {
				t.Errorf("FizzBuzz(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "英字", s: "abc", want: "cba"},
		{name: "日本語", s: "こんにちは", want: "はちにんこ"},
		{name: "混在", s: "Go言語", want: "語言oG"},
		{name: "1文字", s: "あ", want: "あ"},
		{name: "空文字列", s: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.s)
			if got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
