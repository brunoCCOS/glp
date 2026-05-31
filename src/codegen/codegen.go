package codegen

import (
	"fmt"
	"math"
	"strings"

	"propc/src/parser"
	"propc/src/typecheck"
)

const gap = 0.25

// box is the result of laying out a subdiagram. It is positioned in its own
// local frame with origin at (0,0) (lower-left). Left and Right hold the TikZ
// coordinate names of the dangling input and output wires, top-first.
type box struct {
	w, h        float64
	left, right []string
	body        string
}

type renderer struct {
	env     typecheck.Env
	counter int
}

func (r *renderer) id() string {
	r.counter++
	return fmt.Sprintf("n%d", r.counter)
}

func Generate(n parser.Node, env typecheck.Env) (string, error) {
	r := &renderer{env: env}
	b, err := r.render(n)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	sb.WriteString("\\begin{tikzpicture}\n")
	sb.WriteString(b.body)
	sb.WriteString("\\end{tikzpicture}\n")
	return sb.String(), nil
}

func (r *renderer) render(n parser.Node) (box, error) {
	switch v := n.(type) {
	case parser.Id:
		return r.renderId(v), nil
	case parser.Swap:
		return r.renderSwap(v), nil
	case parser.Gen:
		return r.renderGen(v)
	case parser.Tensor:
		return r.renderTensor(v)
	case parser.Comp:
		return r.renderComp(v)
	}
	return box{}, fmt.Errorf("unknown node %T", n)
}

func (r *renderer) renderId(v parser.Id) box {
	// Identity is drawn as a single horizontal wire regardless of N; N is a
	// purely type-level multiplicity. The visible box matches the standard
	// generator frame: width 1, height 1, wire at y = 0.5.
	const w, h, y = 1.0, 1.0, 0.5
	id := r.id()
	left := make([]string, v.N)
	right := make([]string, v.N)
	var sb strings.Builder
	l := fmt.Sprintf("%s-in-0", id)
	rt := fmt.Sprintf("%s-out-0", id)
	fmt.Fprintf(&sb, "  \\coordinate (%s) at (0,%.3f);\n", l, y)
	fmt.Fprintf(&sb, "  \\coordinate (%s) at (%.3f,%.3f);\n", rt, w, y)
	fmt.Fprintf(&sb, "  \\draw (%s) -- (%s);\n", l, rt)
	// Expose every type-level wire as the same physical anchor; composition
	// against a generator that needs multiple tips will fan out from this one.
	for i := uint(0); i < v.N; i++ {
		left[i] = l
		right[i] = rt
	}
	return box{w: w, h: h, left: left, right: right, body: sb.String()}
}

func (r *renderer) renderSwap(v parser.Swap) box {
	// Swap is always drawn as two crossing wires regardless of M and N; the
	// type-level counts only determine how many incoming/outgoing wires alias
	// onto each visual tip. Top bundle (M wires) crosses down to the lower
	// output; bottom bundle (N wires) crosses up to the upper output.
	// Tips match the two-tip generator convention (scope y = h and y = 0)
	// so swap composes flush with generators above/below without bending.
	const w, h, yTop, yBot = 1.0, 1.0, 1.0, 0.0
	id := r.id()
	inTop := fmt.Sprintf("%s-in-0", id)
	inBot := fmt.Sprintf("%s-in-1", id)
	outTop := fmt.Sprintf("%s-out-0", id)
	outBot := fmt.Sprintf("%s-out-1", id)
	var sb strings.Builder
	fmt.Fprintf(&sb, "  \\coordinate (%s) at (0,%.3f);\n", inTop, yTop)
	fmt.Fprintf(&sb, "  \\coordinate (%s) at (0,%.3f);\n", inBot, yBot)
	fmt.Fprintf(&sb, "  \\coordinate (%s) at (%.3f,%.3f);\n", outTop, w, yTop)
	fmt.Fprintf(&sb, "  \\coordinate (%s) at (%.3f,%.3f);\n", outBot, w, yBot)
	// Each wire is a smooth S-curve: it leaves horizontally, bends towards the
	// opposite end, and arrives horizontally. The two bezier control points are
	// at the midpoint x with the start and end y values, which gives a
	// stretched S that crosses its partner cleanly at (w/2, h/2).
	mid := w / 2
	fmt.Fprintf(&sb, "  \\draw (%s) .. controls (%.3f,%.3f) and (%.3f,%.3f) .. (%s);\n",
		inBot, mid, yBot, mid, yTop, outTop)
	fmt.Fprintf(&sb, "  \\draw[preaction={draw=white,line width=4pt}] (%s) .. controls (%.3f,%.3f) and (%.3f,%.3f) .. (%s);\n",
		inTop, mid, yTop, mid, yBot, outBot)
	left := make([]string, v.M+v.N)
	right := make([]string, v.M+v.N)
	for i := uint(0); i < v.M; i++ {
		left[i] = inTop
		right[v.N+i] = outBot
	}
	for j := uint(0); j < v.N; j++ {
		left[v.M+j] = inBot
		right[j] = outTop
	}
	return box{w: w, h: h, left: left, right: right, body: sb.String()}
}

func (r *renderer) renderGen(v parser.Gen) (box, error) {
	g, ok := r.env[v.Name]
	if !ok {
		return box{}, fmt.Errorf("unknown generator %q", v.Name)
	}
	a, c, va, vc, err := typecheck.ResolveGen(g, v.Args)
	if err != nil {
		return box{}, err
	}
	pic := g.Pic
	if pic == "" {
		pic = g.Value
	}
	// Pics live in the frame x in [-0.5, 0.5], y in [-0.5, 0.5]. The box
	// convention is "tips on the left edge at scope-x=0, tips on the right edge
	// at scope-x=w", so we place the pic at (w/2, h/2) to map its local origin
	// to the centre of the box.
	w := g.Width
	if w == 0 {
		w = 1
	}
	h := g.Height
	if h == 0 {
		h = 1
	}
	id := r.id()
	left := make([]string, a)
	right := make([]string, c)
	// Each visual tip is the head of a contiguous bundle of type-level wires.
	// With va dividing a, the bundle size is a/va; type wire i aliases visual
	// tip i / (a/va), top-first. Same on the output side.
	if va > 0 {
		bundle := a / va
		for i := uint(0); i < a; i++ {
			left[i] = fmt.Sprintf("%s-in-%d", id, i/bundle)
		}
	}
	if vc > 0 {
		bundle := c / vc
		for j := uint(0); j < c; j++ {
			right[j] = fmt.Sprintf("%s-out-%d", id, j/bundle)
		}
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "  \\pic (%s) at (%.3f,%.3f) {%s};\n", id, w/2, h/2, pic)
	return box{w: w, h: h, left: left, right: right, body: sb.String()}, nil
}

func (r *renderer) renderTensor(v parser.Tensor) (box, error) {
	a, err := r.render(v.Left)
	if err != nil {
		return box{}, err
	}
	b, err := r.render(v.Right)
	if err != nil {
		return box{}, err
	}
	w := math.Max(a.w, b.w)
	h := a.h + b.h
	// A on top (higher y), B on bottom. Center horizontally.
	axOff, ayOff := (w-a.w)/2, b.h
	bxOff, byOff := (w-b.w)/2, 0.0
	var sb strings.Builder
	emitScoped(&sb, axOff, ayOff, a.body)
	emitScoped(&sb, bxOff, byOff, b.body)
	// Extend wires when child is narrower than the tensor width.
	r.extendWires(&sb, a, axOff, 0, true)
	r.extendWires(&sb, a, axOff+a.w, w, false)
	r.extendWires(&sb, b, bxOff, 0, true)
	r.extendWires(&sb, b, bxOff+b.w, w, false)
	left := append(append([]string{}, a.left...), b.left...)
	right := append(append([]string{}, a.right...), b.right...)
	return box{w: w, h: h, left: left, right: right, body: sb.String()}, nil
}

func (r *renderer) renderComp(v parser.Comp) (box, error) {
	a, err := r.render(v.Left)
	if err != nil {
		return box{}, err
	}
	b, err := r.render(v.Right)
	if err != nil {
		return box{}, err
	}
	if len(a.right) != len(b.left) {
		return box{}, fmt.Errorf("composition mismatch: left coarity %d, right arity %d", len(a.right), len(b.left))
	}
	h := math.Max(a.h, b.h)
	ayOff := (h - a.h) / 2
	byOff := (h - b.h) / 2
	w := a.w + gap + b.w
	var sb strings.Builder
	emitScoped(&sb, 0, ayOff, a.body)
	emitScoped(&sb, a.w+gap, byOff, b.body)
	for i := range a.right {
		fmt.Fprintf(&sb, "  \\draw (%s) to[out=0,in=180] (%s);\n", a.right[i], b.left[i])
	}
	return box{w: w, h: h, left: a.left, right: b.right, body: sb.String()}, nil
}

func emitScoped(sb *strings.Builder, x, y float64, body string) {
	if body == "" {
		return
	}
	if x == 0 && y == 0 {
		sb.WriteString(body)
		return
	}
	fmt.Fprintf(sb, "  \\begin{scope}[shift={(%.3f,%.3f)}]\n%s  \\end{scope}\n", x, y, body)
}

// extendWires draws a horizontal stub from the named anchors of b out to the
// target x coordinate (in the parent frame). If left is true, extends the
// child's left anchors to the parent's left edge; otherwise extends the right
// anchors to the parent's right edge. Skips when the child already reaches
// that edge.
func (r *renderer) extendWires(sb *strings.Builder, b box, currentX, targetX float64, left bool) {
	if math.Abs(targetX-currentX) < 1e-9 {
		return
	}
	anchors := b.left
	if !left {
		anchors = b.right
	}
	for _, a := range anchors {
		stub := r.id()
		fmt.Fprintf(sb, "  \\coordinate (%s) at (%.3f,0 |- %s);\n", stub, targetX, a)
		fmt.Fprintf(sb, "  \\draw (%s) -- (%s);\n", a, stub)
	}
}
