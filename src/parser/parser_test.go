package parser

import (
	"reflect"
	"testing"
)

func TestParseAtoms(t *testing.T) {
	cases := []struct {
		in   string
		want Node
	}{
		{"id(3)", Id{N: 3}},
		{"swap(1,2)", Swap{M: 1, N: 2}},
		{"foo", Gen{Name: "foo"}},
	}
	for _, tc := range cases {
		got, err := Parse(tc.in)
		if err != nil {
			t.Fatalf("%q: %v", tc.in, err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%q: got %#v, want %#v", tc.in, got, tc.want)
		}
	}
}

func TestParsePrecedence(t *testing.T) {
	// * binds tighter than ;  =>  a*b ; c*d  parses as (a*b) ; (c*d)
	got, err := Parse("a * b ; c * d")
	if err != nil {
		t.Fatal(err)
	}
	want := Comp{
		Left:  Tensor{Left: Gen{Name: "a"}, Right: Gen{Name: "b"}},
		Right: Tensor{Left: Gen{Name: "c"}, Right: Gen{Name: "d"}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestParseLeftAssoc(t *testing.T) {
	got, err := Parse("a ; b ; c")
	if err != nil {
		t.Fatal(err)
	}
	want := Comp{Left: Comp{Left: Gen{Name: "a"}, Right: Gen{Name: "b"}}, Right: Gen{Name: "c"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestParseParens(t *testing.T) {
	got, err := Parse("a ; (b ; c)")
	if err != nil {
		t.Fatal(err)
	}
	want := Comp{Left: Gen{Name: "a"}, Right: Comp{Left: Gen{Name: "b"}, Right: Gen{Name: "c"}}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestParseErrors(t *testing.T) {
	bad := []string{"id", "id(", "id()", "swap(1)", "(a;b", "a;", ";a", "a b"}
	for _, in := range bad {
		if _, err := Parse(in); err == nil {
			t.Errorf("%q: expected error", in)
		}
	}
}
