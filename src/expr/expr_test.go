package expr

import "testing"

func TestEval(t *testing.T) {
	vars := map[string]uint{"n": 3, "m": 4}
	cases := []struct {
		in   string
		want uint
	}{
		{"0", 0},
		{"42", 42},
		{"n", 3},
		{"2*n", 6},
		{"2*n + 1", 7},
		{"(n+1)*2", 8},
		{"n*m", 12},
		{"m - n", 1},
		{"  m / n ", 1},
		{"2*(n+m)/n", 4},
	}
	for _, tc := range cases {
		got, err := Eval(tc.in, vars)
		if err != nil {
			t.Fatalf("%q: %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("%q: got %d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestEvalErrors(t *testing.T) {
	cases := []string{"", "1+", "(1+2", "x", "1/0", "1-2", "@"}
	for _, in := range cases {
		if _, err := Eval(in, map[string]uint{}); err == nil {
			t.Errorf("%q: expected error", in)
		}
	}
}
