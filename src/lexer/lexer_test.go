package lexer

import "testing"

func kinds(toks []Token) []TokenKind {
	out := make([]TokenKind, len(toks))
	for i, t := range toks {
		out[i] = t.Kind
	}
	return out
}

func TestTokenizeBasic(t *testing.T) {
	cases := []struct {
		in   string
		want []TokenKind
	}{
		{"id(2)", []TokenKind{ID, LPAREN, NUMBER, RPAREN, EOF}},
		{"swap(1,3)", []TokenKind{SWAP, LPAREN, NUMBER, COMMA, NUMBER, RPAREN, EOF}},
		{"f ; g", []TokenKind{IDENT, COMP, IDENT, EOF}},
		{"f * g", []TokenKind{IDENT, TENSOR, IDENT, EOF}},
		{"(f;g)*h", []TokenKind{LPAREN, IDENT, COMP, IDENT, RPAREN, TENSOR, IDENT, EOF}},
		{"  \t\n", []TokenKind{EOF}},
	}
	for _, tc := range cases {
		got, err := New(tc.in).Tokenize()
		if err != nil {
			t.Fatalf("%q: %v", tc.in, err)
		}
		g := kinds(got)
		if len(g) != len(tc.want) {
			t.Fatalf("%q: got %v, want %v", tc.in, g, tc.want)
		}
		for i := range g {
			if g[i] != tc.want[i] {
				t.Errorf("%q: tok %d got %s, want %s", tc.in, i, g[i], tc.want[i])
			}
		}
	}
}

func TestTokenizeIdentValues(t *testing.T) {
	toks, err := New("multiplication foo_bar x1").Tokenize()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"multiplication", "foo_bar", "x1"}
	for i, w := range want {
		if toks[i].Value != w {
			t.Errorf("tok %d: got %q, want %q", i, toks[i].Value, w)
		}
	}
}

func TestTokenizeError(t *testing.T) {
	if _, err := New("f @ g").Tokenize(); err == nil {
		t.Fatal("expected error on '@'")
	}
}
