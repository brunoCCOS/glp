package typecheck

import (
	"fmt"

	"propc/src/expr"
	"propc/src/lexer"
	"propc/src/parser"
)

// Sig is the (arity, coarity) of a PROP morphism: Arity wires in, Coarity out.
type Sig struct{ Arity, Coarity uint }

// Env maps generator names to their declared signatures.
type Env map[string]lexer.Generator

func NewEnv(gens []lexer.Generator) Env {
	e := make(Env, len(gens))
	for _, g := range gens {
		e[g.Value] = g
	}
	return e
}

// ResolveGen evaluates the parametric signature of a generator invocation.
// It is exported because codegen also needs the type and visual arities.
func ResolveGen(g lexer.Generator, args []uint) (typeArity, typeCoarity, visualArity, visualCoarity uint, err error) {
	if len(args) != len(g.Params) {
		return 0, 0, 0, 0, fmt.Errorf("%s: expected %d argument(s), got %d", g.Value, len(g.Params), len(args))
	}
	bind := make(map[string]uint, len(g.Params))
	for i, p := range g.Params {
		bind[p] = args[i]
	}
	a, err := expr.Eval(g.Arity, bind)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("%s.arity: %w", g.Value, err)
	}
	c, err := expr.Eval(g.Coarity, bind)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("%s.coarity: %w", g.Value, err)
	}
	va := g.VisualArity
	if va == 0 {
		va = a
	}
	vc := g.VisualCoarity
	if vc == 0 {
		vc = c
	}
	if va > 0 && a%va != 0 {
		return 0, 0, 0, 0, fmt.Errorf("%s: visualArity %d does not divide type arity %d", g.Value, va, a)
	}
	if vc > 0 && c%vc != 0 {
		return 0, 0, 0, 0, fmt.Errorf("%s: visualCoarity %d does not divide type coarity %d", g.Value, vc, c)
	}
	return a, c, va, vc, nil
}

func Check(n parser.Node, env Env) (Sig, error) {
	switch x := n.(type) {
	case parser.Id:
		return Sig{x.N, x.N}, nil
	case parser.Swap:
		return Sig{x.M + x.N, x.M + x.N}, nil
	case parser.Gen:
		g, ok := env[x.Name]
		if !ok {
			return Sig{}, fmt.Errorf("unknown generator %q", x.Name)
		}
		a, c, _, _, err := ResolveGen(g, x.Args)
		if err != nil {
			return Sig{}, err
		}
		return Sig{a, c}, nil
	case parser.Tensor:
		l, err := Check(x.Left, env)
		if err != nil {
			return Sig{}, err
		}
		r, err := Check(x.Right, env)
		if err != nil {
			return Sig{}, err
		}
		return Sig{l.Arity + r.Arity, l.Coarity + r.Coarity}, nil
	case parser.Comp:
		l, err := Check(x.Left, env)
		if err != nil {
			return Sig{}, err
		}
		r, err := Check(x.Right, env)
		if err != nil {
			return Sig{}, err
		}
		if l.Coarity != r.Arity {
			return Sig{}, fmt.Errorf("composition mismatch: left has coarity %d, right has arity %d", l.Coarity, r.Arity)
		}
		return Sig{l.Arity, r.Coarity}, nil
	}
	return Sig{}, fmt.Errorf("unknown node %T", n)
}
