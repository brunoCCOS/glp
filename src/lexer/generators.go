package lexer

// Generator describes a user-defined PROP generator.
//
// Arity and Coarity are arithmetic expressions over Params, evaluated at
// call sites: a generator declared with Params ["n"] and Arity "2*n" reports
// type arity 6 when invoked as foo(3). Non-parametric generators leave Params
// empty and use bare integer literals for Arity/Coarity.
//
// VisualArity and VisualCoarity are the number of *drawn* tips the associated
// pic exposes on the left and right. They are independent of the parametric
// counts: a parametric multiplication is depicted as one pic with two input
// tips and one output tip regardless of n. When omitted on a non-parametric
// generator, they default to the constant-evaluated Arity / Coarity.
//
// VisualArity must divide the evaluated type arity (same for coarity); the
// type-level wires are partitioned into equal-sized bundles, each bundle
// aliasing one visual tip top-first.
type Generator struct {
	Value         string   `json:"name"`
	Params        []string `json:"params,omitempty"`
	Arity         string   `json:"arity"`
	Coarity       string   `json:"coarity"`
	VisualArity   uint     `json:"visualArity,omitempty"`
	VisualCoarity uint     `json:"visualCoarity,omitempty"`
	Symbol        string   `json:"symbol,omitempty"`
	Pic           string   `json:"pic,omitempty"`
	Width         float64  `json:"width,omitempty"`
	Height        float64  `json:"height,omitempty"`
}
