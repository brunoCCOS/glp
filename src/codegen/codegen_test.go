package codegen

import (
	"strings"
	"testing"

	"propc/src/lexer"
	"propc/src/parser"
	"propc/src/typecheck"
)

// Smoke tests only. The layout (heights, alignment, wire routing) is being
// reworked, so these assert that codegen runs end-to-end and emits a valid
// tikzpicture wrapper rather than freezing a particular geometry.

func env() typecheck.Env {
	return typecheck.NewEnv([]lexer.Generator{
		{Value: "mult", Arity: "2", Coarity: "1", Pic: "multiplication"},
		{Value: "copy", Arity: "1", Coarity: "2", Pic: "copy"},
		{Value: "nmult", Params: []string{"n"}, Arity: "2*n", Coarity: "n",
			VisualArity: 2, VisualCoarity: 1, Pic: "multiplication"},
	})
}

func gen(t *testing.T, expr string) string {
	t.Helper()
	n, err := parser.Parse(expr)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	out, err := Generate(n, env())
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	return out
}

func TestGenerateWrapsTikzpicture(t *testing.T) {
	out := gen(t, "id(1)")
	if !strings.Contains(out, "\\begin{tikzpicture}") || !strings.Contains(out, "\\end{tikzpicture}") {
		t.Fatalf("missing tikzpicture wrapper:\n%s", out)
	}
}

func TestGenerateGeneratorEmitsPic(t *testing.T) {
	out := gen(t, "mult")
	if !strings.Contains(out, "\\pic") || !strings.Contains(out, "multiplication") {
		t.Fatalf("expected \\pic with multiplication, got:\n%s", out)
	}
}

func TestGenerateComposition(t *testing.T) {
	// copy ; mult is well-typed (1 → 2 → 1) so codegen should succeed.
	out := gen(t, "copy ; mult")
	if !strings.Contains(out, "multiplication") || !strings.Contains(out, "copy") {
		t.Fatalf("expected both pics, got:\n%s", out)
	}
}

func TestGenerateParametric(t *testing.T) {
	// nmult(3) typechecks 6 -> 3. Codegen must emit one multiplication pic
	// regardless, with 6 left anchors and 3 right anchors aliased onto its
	// two visual in tips and one visual out tip.
	out := gen(t, "nmult(3)")
	if !strings.Contains(out, "multiplication") {
		t.Fatalf("expected pic, got:\n%s", out)
	}
}

func TestGenerateUnknownGenerator(t *testing.T) {
	n, err := parser.Parse("nope")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Generate(n, env()); err == nil {
		t.Fatal("expected error on unknown generator")
	}
}
