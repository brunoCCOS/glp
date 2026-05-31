package typecheck

import (
	"testing"

	"propc/src/lexer"
	"propc/src/parser"
)

func env() Env {
	return NewEnv([]lexer.Generator{
		{Value: "mult", Arity: "2", Coarity: "1"},
		{Value: "copy", Arity: "1", Coarity: "2"},
		{Value: "nmult", Params: []string{"n"}, Arity: "2*n", Coarity: "n", VisualArity: 2, VisualCoarity: 1},
	})
}

func mustParse(t *testing.T, s string) parser.Node {
	t.Helper()
	n, err := parser.Parse(s)
	if err != nil {
		t.Fatalf("parse %q: %v", s, err)
	}
	return n
}

func TestCheckAtoms(t *testing.T) {
	cases := []struct {
		in   string
		want Sig
	}{
		{"id(3)", Sig{3, 3}},
		{"swap(2,3)", Sig{5, 5}},
		{"mult", Sig{2, 1}},
		{"copy", Sig{1, 2}},
	}
	for _, tc := range cases {
		got, err := Check(mustParse(t, tc.in), env())
		if err != nil {
			t.Fatalf("%q: %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("%q: got %+v, want %+v", tc.in, got, tc.want)
		}
	}
}

func TestCheckTensor(t *testing.T) {
	got, err := Check(mustParse(t, "mult * copy"), env())
	if err != nil {
		t.Fatal(err)
	}
	if got != (Sig{3, 3}) {
		t.Errorf("got %+v", got)
	}
}

func TestCheckComp(t *testing.T) {
	got, err := Check(mustParse(t, "copy ; mult"), env())
	if err != nil {
		t.Fatal(err)
	}
	if got != (Sig{1, 1}) {
		t.Errorf("got %+v", got)
	}
}

func TestCheckCompMismatch(t *testing.T) {
	if _, err := Check(mustParse(t, "mult ; mult"), env()); err == nil {
		t.Fatal("expected composition mismatch")
	}
}

func TestCheckUnknownGen(t *testing.T) {
	if _, err := Check(mustParse(t, "nope"), env()); err == nil {
		t.Fatal("expected unknown-generator error")
	}
}

func TestCheckParametric(t *testing.T) {
	got, err := Check(mustParse(t, "nmult(3)"), env())
	if err != nil {
		t.Fatal(err)
	}
	if got != (Sig{6, 3}) {
		t.Errorf("nmult(3): got %+v, want {6 3}", got)
	}
}

func TestCheckParametricArgsMismatch(t *testing.T) {
	if _, err := Check(mustParse(t, "nmult"), env()); err == nil {
		t.Fatal("expected missing-args error")
	}
	if _, err := Check(mustParse(t, "nmult(1,2)"), env()); err == nil {
		t.Fatal("expected too-many-args error")
	}
}

func TestCheckComposite(t *testing.T) {
	// (id(1) * mult) ; swap(1,0) ; (id(1) * id(0))  --> id-style
	// Use a simpler one: (id(1) * mult) ; mult  has arities 1+2=3 → 1+1=2 → 1
	got, err := Check(mustParse(t, "(id(1) * mult) ; mult"), env())
	if err != nil {
		t.Fatal(err)
	}
	if got != (Sig{3, 1}) {
		t.Errorf("got %+v", got)
	}
}
