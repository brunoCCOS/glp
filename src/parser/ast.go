package parser

type Node interface{ node() }

type Comp struct {
	Left, Right Node
}

type Tensor struct {
	Left, Right Node
}

// Id is the identity on N wires (id_N: N -> N).
type Id struct{ N uint }

// Swap is the symmetry σ_{M,N}: M+N -> N+M that swaps a block of M wires
// past a block of N wires.
type Swap struct{ M, N uint }

// Gen references a user-defined generator by name. Args holds the numeric
// arguments supplied at the call site (nil for non-parametric calls).
// Resolution against the generator table happens in a later pass.
type Gen struct {
	Name string
	Args []uint
}

func (Comp) node()   {}
func (Tensor) node() {}
func (Id) node()     {}
func (Swap) node()   {}
func (Gen) node()    {}
