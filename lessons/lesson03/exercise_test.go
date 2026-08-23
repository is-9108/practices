// このファイルは「採点表」です。編集しないでください。
package lesson03

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"testing"
	"time"
)

// ---------- 課題1: Counter ----------

func TestCounter(t *testing.T) {
	t.Run("ゼロ値は 0", func(t *testing.T) {
		var c Counter
		if got := c.Value(); got != 0 {
			t.Errorf("var c Counter のあと c.Value() = %d, want 0", got)
		}
	})

	t.Run("Inc で 1 ずつ増える", func(t *testing.T) {
		var c Counter
		c.Inc()
		c.Inc()
		c.Inc()
		if got := c.Value(); got != 3 {
			t.Errorf("Inc を3回呼んだあとの Value() = %d, want 3\n"+
				"        レシーバが値 (c Counter) のままだと、メソッドはコピーを書き換えるだけで\n"+
				"        呼び出し側の Counter は変わりません。(c *Counter) に直してください。", got)
		}
	})

	t.Run("Add は負の値も足せる", func(t *testing.T) {
		var c Counter
		c.Add(10)
		c.Add(-3)
		if got := c.Value(); got != 7 {
			t.Errorf("Add(10), Add(-3) のあとの Value() = %d, want 7", got)
		}
	})

	t.Run("Reset で 0 に戻る", func(t *testing.T) {
		var c Counter
		c.Add(42)
		c.Reset()
		if got := c.Value(); got != 0 {
			t.Errorf("Reset() のあとの Value() = %d, want 0", got)
		}
	})

	t.Run("Counter どうしは独立している", func(t *testing.T) {
		var a, b Counter
		a.Add(5)
		if got := b.Value(); got != 0 {
			t.Errorf("別の Counter まで変わりました: b.Value() = %d, want 0", got)
		}
	})
}

// ---------- 課題2: User ----------

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		userName string
		email    string
		wantName string // 成功時に期待する Name（TrimSpace 後）
		wantErr  error
	}{
		{name: "正常", id: 1, userName: "太郎", email: "taro@example.com", wantName: "太郎"},
		{name: "名前の前後の空白は削る", id: 1, userName: "  太郎 ", email: "a@b", wantName: "太郎"},
		{name: "ID が 0", id: 0, userName: "太郎", email: "a@b", wantErr: ErrInvalidID},
		{name: "ID が負", id: -1, userName: "太郎", email: "a@b", wantErr: ErrInvalidID},
		{name: "名前が空", id: 1, userName: "", email: "a@b", wantErr: ErrEmptyName},
		{name: "名前が空白のみ", id: 1, userName: "   ", email: "a@b", wantErr: ErrEmptyName},
		{name: "@ がない", id: 1, userName: "太郎", email: "example.com", wantErr: ErrInvalidEmail},
		{name: "@ で始まる", id: 1, userName: "太郎", email: "@example.com", wantErr: ErrInvalidEmail},
		{name: "@ で終わる", id: 1, userName: "太郎", email: "taro@", wantErr: ErrInvalidEmail},
		{name: "検証順序: ID が優先", id: 0, userName: "", email: "", wantErr: ErrInvalidID},
		{name: "検証順序: 名前がメールより優先", id: 1, userName: "", email: "bad", wantErr: ErrEmptyName},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.id, tt.userName, tt.email)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewUser(%d, %q, %q) error = %v, want %v", tt.id, tt.userName, tt.email, err, tt.wantErr)
			}

			if tt.wantErr != nil {
				if got != nil {
					t.Errorf("エラー時は nil を返してください: got = %+v", got)
				}
				return
			}

			if got == nil {
				t.Fatalf("NewUser(%d, %q, %q) = nil, want *User", tt.id, tt.userName, tt.email)
			}
			if got.ID != tt.id {
				t.Errorf("ID = %d, want %d", got.ID, tt.id)
			}
			if got.Name != tt.wantName {
				t.Errorf("Name = %q, want %q", got.Name, tt.wantName)
			}
			if got.Email != tt.email {
				t.Errorf("Email = %q, want %q (email はそのまま保存)", got.Email, tt.email)
			}
		})
	}
}

func TestUserRename(t *testing.T) {
	t.Run("名前が変わる", func(t *testing.T) {
		u := &User{ID: 1, Name: "太郎", Email: "a@b"}
		if err := u.Rename("  次郎 "); err != nil {
			t.Fatalf("Rename() error = %v, want nil", err)
		}
		if u.Name != "次郎" {
			t.Errorf("Rename 後の Name = %q, want %q (TrimSpace して保存)\n"+
				"        呼び出し側の User が変わらない場合、レシーバがポインタになっているか確認してください。", u.Name, "次郎")
		}
	})

	t.Run("空白のみは失敗し、値は変わらない", func(t *testing.T) {
		u := &User{ID: 1, Name: "太郎", Email: "a@b"}
		err := u.Rename("   ")
		if !errors.Is(err, ErrEmptyName) {
			t.Fatalf("Rename(\"   \") error = %v, want %v", err, ErrEmptyName)
		}
		if u.Name != "太郎" {
			t.Errorf("失敗したのに Name が変わりました: %q, want %q", u.Name, "太郎")
		}
	})
}

// ---------- 課題3: Stack ----------

func TestStack(t *testing.T) {
	t.Run("ゼロ値のまま使える", func(t *testing.T) {
		var s Stack // NewStack を呼ばない
		s.Push(1)
		s.Push(2)
		if got := s.Len(); got != 2 {
			t.Fatalf("Push を2回したあとの Len() = %d, want 2", got)
		}
	})

	t.Run("LIFO の順で取り出せる", func(t *testing.T) {
		var s Stack
		for _, v := range []int{1, 2, 3} {
			s.Push(v)
		}
		want := []int{3, 2, 1}
		for i, w := range want {
			got, ok := s.Pop()
			if !ok {
				t.Fatalf("%d 回目の Pop() の ok = false, want true", i+1)
			}
			if got != w {
				t.Errorf("%d 回目の Pop() = %d, want %d (後入れ先出し)", i+1, got, w)
			}
		}
		if got := s.Len(); got != 0 {
			t.Errorf("全部 Pop したあとの Len() = %d, want 0", got)
		}
	})

	t.Run("Peek は取り除かない", func(t *testing.T) {
		var s Stack
		s.Push(10)
		s.Push(20)

		got, ok := s.Peek()
		if !ok || got != 20 {
			t.Fatalf("Peek() = (%d, %v), want (20, true)", got, ok)
		}
		if l := s.Len(); l != 2 {
			t.Errorf("Peek のあとの Len() = %d, want 2 (Peek は取り除かない)", l)
		}
	})

	t.Run("空のときは ok が false", func(t *testing.T) {
		var s Stack
		if v, ok := s.Pop(); ok || v != 0 {
			t.Errorf("空スタックの Pop() = (%d, %v), want (0, false)", v, ok)
		}
		if v, ok := s.Peek(); ok || v != 0 {
			t.Errorf("空スタックの Peek() = (%d, %v), want (0, false)", v, ok)
		}
		if l := s.Len(); l != 0 {
			t.Errorf("空スタックの Len() = %d, want 0", l)
		}
	})
}

// ---------- 課題4: Money ----------

func mustMoney(t *testing.T, amount int64, currency string) Money {
	t.Helper() // 失敗時の行番号を「呼び出し元」で表示させる。実務でよく使うテクニック
	m, err := NewMoney(amount, currency)
	if err != nil {
		t.Fatalf("NewMoney(%d, %q) が予期せず失敗しました: %v", amount, currency, err)
	}
	return m
}

func TestNewMoney(t *testing.T) {
	t.Run("正常", func(t *testing.T) {
		m := mustMoney(t, 1050, "JPY")
		if m.Amount() != 1050 {
			t.Errorf("Amount() = %d, want 1050", m.Amount())
		}
		if m.Currency() != "JPY" {
			t.Errorf("Currency() = %q, want %q", m.Currency(), "JPY")
		}
	})

	t.Run("マイナス金額は許可する", func(t *testing.T) {
		m := mustMoney(t, -500, "JPY")
		if m.Amount() != -500 {
			t.Errorf("Amount() = %d, want -500", m.Amount())
		}
	})

	t.Run("通貨が空ならエラー", func(t *testing.T) {
		got, err := NewMoney(100, "")
		if !errors.Is(err, ErrInvalidCurrency) {
			t.Fatalf("NewMoney(100, \"\") error = %v, want %v", err, ErrInvalidCurrency)
		}
		if got != (Money{}) {
			t.Errorf("エラー時の戻り値 = %v, want ゼロ値の Money", got)
		}
	})
}

func TestMoneyAdd(t *testing.T) {
	t.Run("同じ通貨なら足せる", func(t *testing.T) {
		a := mustMoney(t, 1000, "JPY")
		b := mustMoney(t, 250, "JPY")

		got, err := a.Add(b)
		if err != nil {
			t.Fatalf("Add() error = %v, want nil", err)
		}
		if got.Amount() != 1250 || got.Currency() != "JPY" {
			t.Errorf("Add() = %d %s, want 1250 JPY", got.Amount(), got.Currency())
		}
	})

	t.Run("元の値は変わらない（不変）", func(t *testing.T) {
		a := mustMoney(t, 1000, "JPY")
		b := mustMoney(t, 250, "JPY")

		_, _ = a.Add(b)

		if a.Amount() != 1000 {
			t.Errorf("Add 後に a が変わりました: %d, want 1000", a.Amount())
		}
		if b.Amount() != 250 {
			t.Errorf("Add 後に b が変わりました: %d, want 250", b.Amount())
		}
	})

	t.Run("通貨が違うとエラー", func(t *testing.T) {
		a := mustMoney(t, 1000, "JPY")
		b := mustMoney(t, 10, "USD")

		got, err := a.Add(b)
		if !errors.Is(err, ErrCurrencyMismatch) {
			t.Fatalf("Add() error = %v, want %v", err, ErrCurrencyMismatch)
		}
		if got != (Money{}) {
			t.Errorf("エラー時の戻り値 = %v, want ゼロ値の Money", got)
		}
	})
}

func TestMoneyValueSemantics(t *testing.T) {
	t.Run("== で比較できる", func(t *testing.T) {
		a := mustMoney(t, 1000, "JPY")
		b := mustMoney(t, 1000, "JPY")
		c := mustMoney(t, 1000, "USD")

		if a != b {
			t.Errorf("同じ金額・同じ通貨なのに a != b になりました")
		}
		if a == c {
			t.Errorf("通貨が違うのに a == c になりました")
		}
	})

	t.Run("String は \"1050 JPY\" の形式", func(t *testing.T) {
		m := mustMoney(t, 1050, "JPY")
		if got := m.String(); got != "1050 JPY" {
			t.Errorf("String() = %q, want %q", got, "1050 JPY")
		}
	})

	t.Run("fmt が String を自動で使う", func(t *testing.T) {
		m := mustMoney(t, -30, "USD")
		if got := fmt.Sprintf("%v", m); got != "-30 USD" {
			t.Errorf("fmt.Sprintf(\"%%v\", m) = %q, want %q\n"+
				"        String() string を持つ型は fmt.Stringer を満たし、fmt が自動的に使います。", got, "-30 USD")
		}
	})
}

// ---------- 課題5: Base / Article ----------

func TestArticle(t *testing.T) {
	created := time.Date(2026, 8, 1, 12, 0, 0, 0, time.UTC)
	now := created.Add(48 * time.Hour)

	t.Run("NewArticle が全フィールドを埋める", func(t *testing.T) {
		a := NewArticle(7, "Go入門", "太郎", created)

		if a.Title != "Go入門" {
			t.Errorf("Title = %q, want %q", a.Title, "Go入門")
		}
		if a.Author != "太郎" {
			t.Errorf("Author = %q, want %q", a.Author, "太郎")
		}
		// 埋め込みフィールドは a.Base.ID とも a.ID とも書ける（昇格）
		if a.Base.ID != 7 {
			t.Errorf("Base.ID = %d, want 7 (埋め込みは Base: Base{...} で初期化します)", a.Base.ID)
		}
		if a.ID != 7 {
			t.Errorf("a.ID = %d, want 7 (埋め込みフィールドは昇格して直接読めます)", a.ID)
		}
		if !a.CreatedAt.Equal(created) {
			t.Errorf("CreatedAt = %v, want %v", a.CreatedAt, created)
		}
	})

	t.Run("Base のメソッドが昇格している", func(t *testing.T) {
		a := NewArticle(7, "Go入門", "太郎", created)

		// Article には Age を定義していないのに呼べる
		if got := a.Age(now); got != 48*time.Hour {
			t.Errorf("a.Age(now) = %v, want %v", got, 48*time.Hour)
		}
	})

	t.Run("Base.Label", func(t *testing.T) {
		b := Base{ID: 7, CreatedAt: created}
		if got := b.Label(); got != "base#7" {
			t.Errorf("Base.Label() = %q, want %q", got, "base#7")
		}
	})

	t.Run("Article.Label が Base.Label を上書きする", func(t *testing.T) {
		a := NewArticle(7, "Go入門", "太郎", created)

		if got := a.Label(); got != "base#7 / Go入門" {
			t.Errorf("a.Label() = %q, want %q\n"+
				"        同名のメソッドは外側 (Article) が優先されます。", got, "base#7 / Go入門")
		}
		if got := a.Base.Label(); got != "base#7" {
			t.Errorf("a.Base.Label() = %q, want %q\n"+
				"        埋め込み側の実装は a.Base.Label() で明示的に呼べます。", got, "base#7")
		}
	})

	t.Run("Age は経過していないとき 0", func(t *testing.T) {
		b := Base{ID: 1, CreatedAt: created}
		if got := b.Age(created); got != 0 {
			t.Errorf("Age(同時刻) = %v, want 0", got)
		}
	})
}

// ---------- 課題6: Profile ----------

func TestProfileClone(t *testing.T) {
	t.Run("Tags が独立している", func(t *testing.T) {
		p := Profile{Name: "太郎", Tags: []string{"go", "docker"}}
		c := p.Clone()

		c.Tags[0] = "rust"
		c.Tags = append(c.Tags, "k8s")

		if !slices.Equal(p.Tags, []string{"go", "docker"}) {
			t.Errorf("Clone の Tags を書き換えたら元まで変わりました: p.Tags = %q, want [go docker]\n"+
				"        struct の代入は浅いコピーです。slice は slices.Clone などで複製してください。", p.Tags)
		}
	})

	t.Run("Meta が独立している", func(t *testing.T) {
		p := Profile{Name: "太郎", Meta: map[string]string{"role": "dev"}}
		c := p.Clone()

		c.Meta["role"] = "ops"
		c.Meta["team"] = "platform"

		want := map[string]string{"role": "dev"}
		if !maps.Equal(p.Meta, want) {
			t.Errorf("Clone の Meta を書き換えたら元まで変わりました: p.Meta = %v, want %v\n"+
				"        map も参照が共有されます。maps.Clone で複製してください。", p.Meta, want)
		}
	})

	t.Run("中身は同じ", func(t *testing.T) {
		p := Profile{
			Name: "太郎",
			Tags: []string{"go"},
			Meta: map[string]string{"role": "dev"},
		}
		c := p.Clone()

		if c.Name != p.Name {
			t.Errorf("Name = %q, want %q", c.Name, p.Name)
		}
		if !slices.Equal(c.Tags, p.Tags) {
			t.Errorf("Tags = %q, want %q", c.Tags, p.Tags)
		}
		if !maps.Equal(c.Meta, p.Meta) {
			t.Errorf("Meta = %v, want %v", c.Meta, p.Meta)
		}
	})

	t.Run("nil は nil のまま", func(t *testing.T) {
		p := Profile{Name: "太郎"} // Tags も Meta も nil
		c := p.Clone()

		if c.Tags != nil {
			t.Errorf("Tags = %v, want nil (nil は nil のまま返してください)", c.Tags)
		}
		if c.Meta != nil {
			t.Errorf("Meta = %v, want nil (nil は nil のまま返してください)", c.Meta)
		}
	})

	t.Run("元を書き換えても複製は影響を受けない", func(t *testing.T) {
		p := Profile{Name: "太郎", Tags: []string{"go"}, Meta: map[string]string{"role": "dev"}}
		c := p.Clone()

		p.Tags[0] = "rust"
		p.Meta["role"] = "ops"

		if !slices.Equal(c.Tags, []string{"go"}) {
			t.Errorf("元を書き換えたら複製まで変わりました: c.Tags = %q, want [go]", c.Tags)
		}
		if !maps.Equal(c.Meta, map[string]string{"role": "dev"}) {
			t.Errorf("元を書き換えたら複製まで変わりました: c.Meta = %v, want map[role:dev]", c.Meta)
		}
	})
}
