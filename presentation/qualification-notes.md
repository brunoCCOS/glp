# Qualification notes — scope, provenance, and committee preparation

These are private preparation notes, not extra slides. The deck is based on the **current source included by main.tex**, not every .tex file in the repository.

The active mathematical text is: sections/introduction.tex; sections/background.tex, which includes sections/quantale.tex and then sections/categorical.tex; sections/gpa.tex; sections/glp.tex; sections/conclusion.tex; and the inherited notation in sections/appendix.tex. The separate sections/gqa.tex is **not included** by main.tex and is not treated as an established chapter of this dissertation.

No dissertation source has been edited. The proposed qualification work plan is an editorial recommendation and must be agreed with the advisor.

## Evidence map

| Deck content | Source |
|---|---|
| Title, author, advisor, programme | config/informacoes_gerais.tex |
| Composition/elimination motivation | sections/introduction.tex; sections/quantale.tex; sections/categorical.tex |
| Graphical language hierarchy and novelty boundary | Introduction's language-hierarchy table |
| GPA syntax, FME, normal forms, presentation | sections/gpa.tex |
| LPRel interface and GLP primitive | sections/glp.tex, semantics and syntax sections |
| The three new equations | tikz/AxiomsGLP.tex, included without redrawing |
| Normal-form picture | tikz/NormalFormGLP.tex, included without redrawing |
| Epigraph-form route to completeness | sections/glp.tex, equal-epigraph proposition, Epigraph Form corollary, completeness lemma |
| Quadratic objectives as future work | sections/conclusion.tex |
| Running two-source example and its plots | Added specifically for the oral presentation; not claimed to be a numerical experiment from the dissertation |
| Final-defense priorities | Proposed preparation roadmap, not sourced claims of completed milestones |

## Mathematical guardrails before presenting

### Extended values must be handled consistently

The draft's LPRel definition allows arbitrary real linear objectives and parameterized constraints. Such LPs can be infeasible or unbounded below. The primitive $\ell(x)=x$ itself is signed, so the nonnegative Lawvere quantale is not enough for the full language.

The deck uses the draft's broader convention from sections/quantale.tex: values in $\overline{\mathbb R}$ and **inf-addition**, with $+\infty$ absorbing, including against $-\infty$.

- Infeasible fiber: value $+\infty$.
- Unbounded-below feasible fiber: value $-\infty$.
- Finite feasible optimum: standard LP attainment applies.

Some definitions early in sections/glp.tex restrict the codomain to $\mathbb R\cup\{+\infty\}$, whereas the broader relational semantics admit $-\infty$. Resolve this mismatch before treating the final presentation theorem as fully audited.

Simply restricting all morphisms to proper finite-valued LPs is not automatically a solution: arbitrary diagram composition can introduce infeasibility or unboundedness. The chosen semantic class must be closed under the language's operations.

### Polyhedral does not imply finite everywhere

The draft's early equivalence statement between polyhedral functions and finite maxima of affine functions needs domain qualifications.

For a **proper** extended-valued polyhedral convex function, the standard form is a finite maximum of affine functions **on a polyhedral domain**, with $+\infty$ outside that domain. A bare maximum of finitely many real affine functions is finite everywhere, so it cannot express infeasible parameter values.

Also distinguish the normalized value function
\[
v(b)=\inf_x\{c^\top x\mid Ax\ge b\}
\]
from the affine-RHS family actually used in LPRel:
\[
F(q)=\inf_z\{c^\top z\mid Az\ge Bq+r\}.
\]
The normalized expression has additional structure; for example, positive scaling of $b$ scales its value. It should not be used without qualification as a representation of arbitrary affine piecewise-linear functions.

The talk does **not** rely on those broad equivalences. It uses a directly computed finite example and the polyhedral epigraph representation associated with the affine-RHS interface.

### The epigraph argument needs both semantic and syntactic parts

For fixed matrices and an objective, consider
\[
E=\{(q,t)\mid\exists z:\ Az\ge Bq+r,\ c^\top z\le t\}.
\]
Projection of a polyhedron is a polyhedron.

For a finite feasible LP value, attainment identifies this set with the ordinary epigraph. For an unbounded feasible fiber, all finite thresholds are achievable; for an infeasible fiber, none are. This is consistent with epigraphs of the extended values $-\infty$ and $+\infty$.

However, the main theorem still needs a **derivation using GLP's axioms** that takes each arbitrary diagram to its epigraph form. Merely drawing some other diagram with the same value function proves neither syntactic equality nor completeness.

Audit in particular the normal-form induction, the epigraph-form corollary, variable types, and the rules used when splitting matrices and bending wires.

### A normal-form shape is not a unique matrix encoding

The deck deliberately says “common structural shape,” not “unique canonical coefficient matrices.”

Equality of polyhedra may be represented by different redundant inequality systems. Equality of LP values may involve different numbers of decision variables. GPA completeness supplies derivable equality of diagrams, not literal equality of those coefficient lists.

A normal-form theorem alone does not establish a chosen terminating/confluent rewrite algorithm. A finite derivation guaranteed by completeness is not a runtime complexity claim.

### Do not overstate categorical embeddings

For the oral talk, ordinary relations embed into cost-valued relations via indicators. This inclusion is faithful and structure-preserving, but it is not full into the category of *all* cost-valued relations: there are genuinely weighted morphisms between the same objects that are not indicators.

Similarly, the draft's unrestricted class of proper, lower-semicontinuous convex bifunctions needs care before asserting closure under arbitrary infimal composition. Projection can lose closedness, and composition can fail properness. The deck does not use that broad class as an unqualified theorem; it concentrates on the LP/polyhedral setting and explicit extended-cost conventions.

### Keep prior work distinct from the new contribution

GPA is from Bonchi, Di Giorgio, and Sobociński's Diagrammatic Polyhedral Algebra (2021). The draft supplies its development and proofs as the foundation for GLP.

GLP's presentation is the draft's claimed contribution. Before the final defense, compare its statement and exact parameter/value-function scope with related prior diagrammatic work, including the bibliography's From Farkas' Lemma to Linear Programming (Bonchi, Di Giorgio, and Zanasi, 2021). Do not turn “new GLP extension in this draft” into an unverified claim of being the first diagrammatic treatment of every kind of LP.

## Likely committee questions

### What exactly is being identified?

**Answer:** Pointwise equality of the optimal-value functions on the external interface. Not equality of optimizer sets, iteration histories, or matrix presentations.

### What is new compared with GPA?

**Answer:** GPA represents feasibility relations, equivalently zero/infinity indicators. GLP adds a primitive scalar linear cost, its three interaction axioms, and the normal-form/epigraph route to a presentation of affine-RHS parametric LP value functions.

### Why is the objective generator a state $0\to1$?

**Answer:** It assigns cost to a variable on its output interface. It does not choose that variable. When connected to other relations, composition minimizes over the internal variable. Compact relational structure permits bending interfaces when changing the presentation.

### How does one scalar cost represent $c^\top z$?

**Answer:** Use inherited linear relations to relate a scalar objective variable to $c^\top z$, then attach the cost generator. The running example shows a reflected scalar enforcing $u=2x_2$.

### Why do LP blocks remain LP blocks after composition?

**Answer:** Combine the decisions, add the linear objectives, and conjoin the linear constraints. The shared external variable becomes an internal decision. Check the extended-value conventions when interpreting all boundary cases.

### Where is linear programming duality used?

**Answer:** It is discussed in the draft's semantic background, but this talk's completeness sketch uses epigraphs, polyhedral projection, and GPA completeness. It does not need an unqualified “finite maximum of affine functions” duality statement.

### What makes the epigraph a polyhedron?

**Answer:** Add the linear inequality $c^\top z\le t$ to the original linear constraints, then project out $z$. Polyhedra are closed under linear projection.

### Is equality of epigraphs enough?

**Answer:** Semantically, it identifies the value functions. Syntactically, we also require the epigraph-form rewrite theorem, followed by GPA completeness and composition with the common cost source.

### Does the calculus solve LPs faster?

**Answer:** No speedup is established here. The result is about representation, composition, and equational reasoning. A solver interface, rewrite implementation, or complexity analysis would be separate work.

### What still needs to be completed for the final defense?

**Answer:** Resolve the semantic boundary-value/domain conventions, audit the normal-form and epigraph-form derivations, and sharpen the comparison to prior work. Application tooling and quadratic extensions are future directions rather than prerequisites for the present claim.

## Proposed work packages

| Priority | Work package | Evidence of completion |
|---|---|---|
| P1 | Align the codomains, infeasible/unbounded behavior, and composition conventions across definitions. | One consistent closed semantic class, checked on boundary cases. |
| P2 | Audit the syntactic normal-form and epigraph derivations. | Each rewrite has a named axiom and consistent typed variables. |
| P3 | Clarify the novelty statement against the cited related work. | Explicit comparison of semantic scope and theorem statements. |
| P4 | Validate representative examples independently. | Values and epigraphs agree before and after the chosen diagrammatic rewrites. |
| Beyond current scope | Prototype rewriting or extend toward quadratic objectives. | Separate implementation or theorem work; no implied completion date. |

Dates are deliberately not invented. Agree a calendar with the advisor after qualification.
