# Master's qualification — speaking guide

**Deck:** [qualification.pdf](qualification.pdf)  
**Source:** [qualification.tex](qualification.tex)  
**Language / slot:** English / 25 minutes  
**Structure:** exactly 27 slides, including the opening and closing; no overlays or backup slides.  
**Timing:** 24 minutes of planned speech + 1 minute of breathing room. Committee questions are not included.

This is a delivery plan, not a transcript to memorize. The “Say” passages give the reasoning to preserve. Speak in your own words. The screen carries the mathematical anchors; your voice supplies the motivation, interpretation, and transitions.

## Selection before sequencing

The central message is:

> LP value functions admit a compositional diagrammatic language; the extension from constraints to objectives is small, and epigraphs provide the route to completeness.

The committee should leave able to identify **the research problem, the new contribution, the proof mechanism, and what remains to validate**.

### Sequencing principle

The deck follows the audience's dependency order, **not** the dissertation's chapter order. The running example comes first, then the notation is introduced before it is used: how to read the diagrams (5), composition as elimination (6), reformulation by regrouping (7), only then the semantic step from feasibility to cost (8) and the quantale that abstracts both (9). Inherited GPA (10–13) sits before the extension (14–19), and epigraphs appear immediately before the proof that uses them (21–23).

### Topics to explain carefully

| Topic | Why it earns speaking time | Slides |
|---|---|---|
| RHS parameters and value-function semantics | Prevents confusion between an LP instance, an optimizer, and the object represented by the language. | 2–3, 15 |
| Reading the diagrams and composition as elimination | The conceptual foundation connecting the graphical syntax to optimization. | 5–6, 8, 16 |
| GPA versus GLP | Makes the novelty boundary unambiguous: GPA is inherited; GLP adds genuine costs. | 10–14 |
| The objective generator and its use | Makes “one generator” a concrete claim rather than a slogan. | 14, 17–18 |
| Normal form versus uniqueness | A proof-supporting shape does not mean a unique matrix representation. | 20 |
| Epigraph form and completeness | This is the technical centerpiece and should receive the slowest delivery. | 21–23 |
| Semantic scope and qualification priorities | Shows that the result is bounded and that the final-defense work is understood. | 26–27 |

### Topics that deserve only a mention

| Topic | Treatment |
|---|---|
| Quantales | One slide: the shared combine/aggregate pattern. No lattice-theory lecture. |
| GLA / GAA | Locate them in the lineage. Do not enumerate their generators or prove their completeness. |
| Fourier–Motzkin elimination | One lower-bound/upper-bound example. Name its role in projection. |
| PROP terminology | If needed: “Objects count real-variable wires; tensor concatenates interfaces.” Do not give the full definition orally. |
| Prior applications of string diagrams | References establish context; quantum computing and electrical circuits need not become a detour. |
| Quadratic extensions | One future-work sentence; no claimed theorem. |

### Topics deliberately omitted from the 27 slides

- Full inherited axiom tables and Frobenius/Hopf algebra derivations.
- Complete matrix-level FME proofs and the polar/Farkas proof machinery.
- Column-stochastic objective identities: interesting, but not needed to understand the result.
- Metric-space examples, transitivity developments, and the full categorical hierarchy.
- Legendre–Fenchel duality, probability, and detailed GQA developments.
- Literature-survey slides and a full-page bibliography; references are given locally instead.
- Claims about speedups, implemented solvers, termination of a chosen rewrite strategy, or benchmark results: the draft does not supply those.
- The draft's broad equivalence between all extended-valued polyhedral functions and finite-valued maxima of affine functions. The defense-safe version requires domain and boundary-value qualifications; see [qualification-notes.md](qualification-notes.md).

## Pace and presentation dynamic

Use a mixed-category-theory/optimization committee as the working audience assumption. The entry point is elementary LP reasoning; do not assume everyone already reads string diagrams.

- Slides 2, 3, 17, and 21 return to **one running example**. This removes the cost of learning a new example during the hardest part of the talk.
- The attention resets are conceptual changes: diagrams before algebra (5); feasibility to cost on 8; inherited GPA to new GLP on 14; algebra to epigraph geometry on 21.
- Do not ask the committee to solve a problem during the 25-minute slot. A brief rhetorical question followed by a pause is enough.
- Use the pointer on the internal wire, the cost glyph, and the epigraph boundary. Avoid tracing every wire on every slide.
- Treat slide 23 as the payoff. Give its arrows more time than the background slides.
- End with the next milestone, not just “thank you.” Leave slide 27 visible for questions.
- Small source credits are provenance, not text to read aloud.

### Rehearsal actions

1. Rehearse without reading the “Say” passages verbatim.
2. Check elapsed time at slides 8, 13, 19, and 23.
3. Protect the epigraph/completeness time if the opening runs long.
4. Use the recovery cuts below instead of speaking faster.

## Timing map

The slide targets sum to **24:00**. The spare minute absorbs short pauses and transitions.

| Slide | Title | Target | Cumulative |
|---:|---|---:|---:|
| 1 | Opening | 0:20 | 0:20 |
| 2 | One demand, two sources | 1:00 | 1:20 |
| 3 | The value function is the observable | 0:50 | 2:10 |
| 4 | The research question | 0:45 | 2:55 |
| 5 | How to read the diagrams | 0:50 | 3:45 |
| 6 | Composition already eliminates variables | 0:50 | 4:35 |
| 7 | Reformulation can be a change of grouping | 0:45 | 5:20 |
| 8 | Replace feasibility by cost | 1:00 | 6:20 |
| 9 | Quantales explain the shared pattern | 0:40 | 7:00 |
| 10 | Build on the existing graphical languages | 0:40 | 7:40 |
| 11 | GPA expresses feasible regions | 0:50 | 8:30 |
| 12 | Projection is the computational primitive | 0:45 | 9:15 |
| 13 | GPA provides the proof engine | 0:50 | 10:05 |
| 14 | The new generator assigns linear cost | 1:00 | 11:05 |
| 15 | The semantics: LPs with external interfaces | 0:50 | 11:55 |
| 16 | LP blocks compose into LP blocks | 1:00 | 12:55 |
| 17 | Our running LP becomes a diagram | 1:00 | 13:55 |
| 18 | Three axioms control the new cost | 0:50 | 14:45 |
| 19 | Expressivity and soundness are direct | 0:45 | 15:30 |
| 20 | Normal form collects the objective | 0:45 | 16:15 |
| 21 | The epigraph turns cost into geometry | 1:20 | 17:35 |
| 22 | Epigraph form is a syntactic bridge | 1:10 | 18:45 |
| 23 | The completeness argument | 1:40 | 20:25 |
| 24 | Simplify the example: collect cost, then eliminate | 1:00 | 21:25 |
| 25 | The simplified diagram has no source decisions | 1:00 | 22:25 |
| 26 | Result, scope, and remaining work | 1:00 | 23:25 |
| 27 | Optimization becomes part of the syntax | 0:35 | 24:00 |
| — | Breathing room | 1:00 | 25:00 |

## Slide-by-slide delivery

### 01 — Opening · 20 seconds

**Goal:** state the problem before definitions.

**Say:** “This work asks whether linear programs can be treated as composable objects, with a graphical language for proving that different formulations compute the same optimal value. I will start with a small LP, then show the extension of the language and the central completeness argument.”

**Delivery:** look at the committee, not the title.

**Transition:** “First, the family of problems I want the language to express.”

### 02 — One demand, two sources · 60 seconds

**Goal:** anchor the notation in a familiar optimization problem.

**Say:** “We must meet demand $$d$$ using two sources. The first costs one per unit, but has capacity one. The second costs two per unit and has no upper capacity bound. The decisions are $$x_1$$ and $$x_2$$. Everything except demand is fixed. As demand changes, this is a family of linear programs whose value is a function of $$d$$; only a right-hand side parameter varies.”

**Delivery:** point first to $$d$$, then to the cheap source's capacity. Do not read all constraints separately. Mention that allowing negative $$d$$ makes the example a function on all of R; the optimum is then zero.

**Transition:** “The object I will identify across different formulations is this value function.”

### 03 — The value function is the observable · 50 seconds

**Goal:** distinguish value functions from optimizers.

**Say:** “For demand between zero and one, the cheap source meets the demand, so the cost is $$d$$. Above one, it is fully used and the second source supplies the remainder, giving $$2d-1$$. The function is piecewise affine and convex. Adding auxiliary variables or decomposing the constraints may change the formulation without changing this curve. That is the equality the calculus will capture.”

**Delivery:** trace the change of slope at one. Do not describe argmin sets.

**Transition:** “So the challenge is not just drawing an LP; it is reasoning about equality of the functions represented.”

### 04 — The research question · 45 seconds

**Goal:** make the contribution evaluable.

**Say:** “The target is a language that expresses the whole semantic class, whose rules preserve meaning, and whose rules are sufficient to prove equalities of meaning. Those are expressivity, soundness, and completeness. We start from an existing language for polyhedral constraints, GPA. The extension for linear objectives adds only one generator and three axioms.”

**Delivery:** pause after “completeness.” The prior/new distinction matters.

**Transition:** “But first, the notation the composition will act on.”

### 05 — How to read the diagrams · 50 seconds

**Goal:** make all later diagrams intelligible.

**Say:** “Read these diagrams left to right. Wires are real-variable interfaces; boxes denote relations or cost bifunctions. Sequential composition identifies the shared wire and eliminates its variable. Parallel composition places independent blocks alongside each other and adds their costs. These are typed diagrams, so the number of wires determines which connections are legal.”

**Delivery:** gesture horizontally for sequential, then to the independent parallel wires. Avoid calling the boxes deterministic functions.

**Transition:** “The first diagram already encodes the central move of optimization.”

### 06 — Composition already eliminates variables · 50 seconds

**Goal:** establish elimination from ordinary relations.

**Say:** “Relational composition connects the output variable of one relation to the input variable of the next. The composite relates $$x$$ to $$z$$ whenever some intermediate $$y$$ satisfies both relations. After composition, $$y$$ is no longer part of the external interface. Connecting these blocks already means eliminating a shared variable.”

**Delivery:** point to the middle wire and then to the two ends.

**Transition:** “One more reading of the same pictures before I change what the boxes mean.”

### 07 — Reformulation can be a change of grouping · 45 seconds

**Goal:** show a benefit of syntax without claiming arbitrary equivalence is trivial.

**Say:** “This is the monoidal interchange law. We may compose each row and then put the results in parallel, or first put the columns in parallel and then compose. We did not change the wiring. A coordinate expression makes this a law to invoke; the string diagram keeps that structural equality visible.”

**Delivery:** trace one row, then one dashed column. Do not read the full equation aloud.

**Transition:** “Now the semantic step: replace feasibility by cost.”

### 08 — Replace feasibility by cost · 60 seconds

**Goal:** establish the single main semantic formula.

**Say:** “Represent a feasible relation by its indicator: zero when feasible and infinity otherwise. Add two such indicators and minimize over the shared variable. We get zero precisely when a feasible connection exists. With genuine costs instead of just zero and infinity, the same operation computes the best combined cost. This is the sense in which optimization is internal to composition.”

**Delivery:** speak the meaning of the infimum over $$y$$ before reading the surrounding expression. Let the formula sit briefly.

**Transition:** “Quantales describe the algebraic pattern common to these two readings.”

### 09 — Quantales explain the shared pattern · 40 seconds

**Goal:** supply context without derailing the talk.

**Say:** “There are two operations here: combine the values along a candidate connection, then aggregate over intermediate variables. For truth, they are conjunction and existential aggregation. For costs, they are addition and infimum. A quantale packages the laws that make this composition associative. I use this as the background framework; I will not develop the general lattice theory today.”

**Delivery:** point between the two table rows. Do not introduce more quantale examples.

**Transition:** “This is the structure the dissertation extends, and the extension builds on prior work.”

### 10 — Build on the existing graphical languages · 40 seconds

**Goal:** locate novelty precisely.

**Say:** “Graphical linear algebra represents linear relations. An affine point extends it to affine relations. An order generator extends it to polyhedral relations. Those theories are inherited from prior work. This dissertation recalls and develops GPA as the foundation, and then adds a linear cost generator to obtain GLP.”

**Delivery:** move through the sequence once. Name the references collectively rather than reading author lists.

**Transition:** “The crucial inherited component is the language for constraints.”

### 11 — GPA expresses feasible regions · 50 seconds

**Goal:** explain the constraint block used throughout the proof.

**Say:** “The order generator denotes the relation $u\ge v$. Combined with matrices, constants, copying and the existing relational structure, it represents finite affine systems of inequalities. An equality can be encoded as two inequalities. In this view, the block still only assigns zero or infinity; it tells us what is feasible, but does not yet assign a linear objective.”

**Delivery:** point to the order glyph and then to the polyhedron formula.

**Transition:** “Composing those constraint blocks must eliminate their hidden variables.”

### 12 — Projection is the computational primitive · 45 seconds

**Goal:** explain FME's role with the smallest useful example.

**Say:** “Suppose $$y$$ must lie above $$x$$ and below $$z$$. Such a $$y$$ exists exactly when $$x$$ lies below $$z$$. Fourier–Motzkin generalizes this by pairing lower bounds with upper bounds, removing the internal variable and leaving inequalities on the remaining ones. In GPA, the rules realize this polyhedral projection.”

**Delivery:** point to the two bounds and then the resulting inequality. No matrix row partitioning.

**Transition:** “This supports the existing presentation result that GLP will reuse.”

### 13 — GPA provides the proof engine · 50 seconds

**Goal:** prepare the audience to accept GPA completeness as a dependency.

**Say:** “GPA presents polyhedral relations. Expressivity means every polyhedral relation can be drawn. Soundness means the equations preserve the relation. Completeness means that if two drawings denote the same polyhedron, the rules can prove them equal. The last direction is the one we will borrow: once LP equality is translated into polyhedral equality, GPA can certify that equality.”

**Delivery:** slow down for the direction of completeness. Do not claim unique coefficient matrices.

**Transition:** “Now we can identify exactly what GLP adds.”

### 14 — The new generator assigns linear cost · 60 seconds

**Goal:** explain the new primitive's type and denotation.

**Say:** “The new generator has no input and one real-variable output. Its denotation is the scalar cost $$x$$ assigned to that output variable. It does not produce a specific optimizer; the variable remains available to connect to constraints. The inherited linear structure turns this primitive into arbitrary linear objectives. Optimizing happens when those variables become internal.”

**Delivery:** distinguish the empty input from the output wire. The glyph is named $\ell$ in the talk.

**Transition:** “This requires semantics with costs on both sides of an interface.”

### 15 — The semantics: LPs with external interfaces · 50 seconds

**Goal:** give the exact semantic class, without reading the whole matrix formula.

**Say:** “A morphism has external variables $$x$$ and $$y$$. The internal decisions are $$z$$. The matrices and objective coefficients are fixed; parameters enter affinely on the right-hand side through the interface. The denotation is the optimum as a function of the external variables. We use extended costs so infeasibility and unboundedness can be represented, with an explicit convention for adding them.”

**Delivery:** point to the variables, not individual matrix entries.

**Transition:** “Connecting two such interfaces stays within linear programming.”

### 16 — LP blocks compose into LP blocks · 60 seconds

**Goal:** justify semantic closure, rather than simply asserting composability.

**Say:** “To compose two LP blocks, collect both sets of internal decisions and their shared interface variable. Their objectives add, and their linear constraints are imposed together. The result is again a linear program. In parallel, there is no shared variable: we simply retain both interfaces and add the independent costs. This is a semantic composition rule, not a claim that evaluating it is computationally free.”

**Delivery:** say why the objective and constraints remain linear. Mention the convention only once.

**Transition:** “Return to the example and see how the primitive actually expresses a weighted objective.”

### 17 — Our running LP becomes a diagram · 60 seconds

**Goal:** demonstrate expressivity concretely.

**Say:** “The copy branches check source one's bounds and source two's nonnegativity. One cost source assigns x1; the reflected scalar on the second branch enforces u equals twice x2. The white addition node combines supplies, and the order generator checks demand. Connecting these generators leaves only demand external.”

**Delivery:** trace the two objective branches, then point to $$d$$. Identify the bound branches as existing GPA generators, not a constraint-box abbreviation.

**Transition:** “Only three new axioms are needed to reason about these costs.”

### 18 — Three axioms control the new cost · 50 seconds

**Goal:** explain what each new rule does, not prove the full axiom system.

**Say:** “The first says that the scalar cost respects addition. The second gives zero when the variable is constrained to zero. The third says that minimizing $$t$$ subject to $$t\ge x$$ recovers $$x$$. That last identity matters most for completeness: it lets us replace an objective by the relation describing all values above it.”

**Delivery:** point to each row briefly; spend the last half on the infimum rule.

**Transition:** “The easy presentation directions follow directly from these ingredients.”

### 19 — Expressivity and soundness are direct · 45 seconds

**Goal:** clear the simpler proof obligations before the centerpiece.

**Say:** “For expressivity, encode the constraints in GPA, encode the objective using the new primitive and linear structure, and compose. For soundness, the old GPA equations are already valid; the three new equations have the scalar identities we just discussed. These directions are direct. The real difficulty is the converse: equality of value functions need not mean identical decision variables or identical constraints.”

**Delivery:** no line-by-line derivation. Emphasize “converse.”

**Transition:** “The proof first organizes arbitrary diagrams into a common shape.”

### 20 — Normal form collects the objective · 45 seconds

**Goal:** explain the role and limits of normal form.

**Say:** “The normal-form theorem collects the objective into one cost source and leaves the constraints in a polyhedral block. It is proved by structural induction, checking the generators and closure under sequential and parallel composition. ‘Normal form’ here means a common structural shape. It does not mean that the matrices are uniquely determined or that the original LP formulation is uniquely recoverable.”

**Delivery:** point to the cost source and the constraint portion only.

**Transition:** “The geometry needed to compare these normal forms is the epigraph.”

### 21 — The epigraph turns cost into geometry · 80 seconds

**Goal:** make the completeness bridge visually inevitable.

**Say:** “The epigraph is everything on or above the graph of a function. For our example, the shaded region is the intersection of the three halfspaces shown: above zero, above $$d$$, and above $$2d-1$$. It is a polyhedron, so GPA can express it. If two LPs have the same value function, they have the same epigraph. We have turned an equality of cost functions into an equality of geometric relations in a language whose completeness is already available.”

**Delivery:** trace the boundary, then point into the shaded region. Pause at “GPA can express it.”

**Transition:** “But a semantic observation alone is not yet a syntactic proof.”

### 22 — Epigraph form is a syntactic bridge · 70 seconds

**Goal:** distinguish existence of an equivalent representation from derivability by the rules.

**Say:** “The figure on screen is the paper's epigraph form: the original diagram rewritable, by the GLP axioms, into one linear cost source followed by a pure GPA epigraph block in the dotted region. Reading it semantically: introduce a threshold variable $$t$$ that bounds the objective, eliminate the original decisions, and what remains is a polyhedron $$E$$ in $$GPA$$. Taking the infimum of $$t$$ over that relation reconstructs the value function. Without the syntactic rewrite — the direction in which the axioms force this form — we would only have a semantic analogy, not completeness.”

**Delivery:** point to the cost source, then to the dotted epigraph block. Do not trace the antipode node unless asked.

**Transition:** “Now the completeness argument has all its dependencies.”

### 23 — The completeness argument · 100 seconds

**Goal:** deliver the main proof slowly enough to be followed.

**Say:** “Start with diagrams $$c$$ and $$d$$ whose value functions are equal. Their epigraph polyhedra are therefore equal. Rewrite each original diagram into its epigraph form. The two epigraph blocks are GPA diagrams with equal polyhedral semantics. GPA completeness equates these blocks by its rules. Finally, attach the same scalar cost generator to both. Equality is preserved by composition, so the original GLP diagrams are equal. The new theory inherits a powerful proof mechanism without requiring an unrelated global equivalence rule.”

**Delivery:** follow the top arrows once. Then point to the epigraph-form assumption at the bottom and explain why it is needed. Finally connect the two inputs to the conclusion. Leave a short pause after the final equality.

**Do not say:** equal semantics forces equal matrices, identical LP formulations, or a unique algorithmic normal form.

**Transition:** “Before stating the result, apply this mechanism to our running example.”

### 24 — Simplify the example: collect cost, then eliminate · 60 seconds

**Goal:** make the axiomatic simplification concrete.

**Say:** “Apply linearity to collect the two costs into c, then use the infimum rule to introduce a threshold t above c. Set s equal to total supply. Since c equals 2s minus x1, the bounds on x1 are summarized by the displayed lower and upper bounds. GPA's Fourier–Motzkin rewrites pair them and eliminate x1. Eliminating s leaves just three halfspaces on t and demand. These are inherited constraint rewrites, not an appeal to GLP completeness.”

**Delivery:** point to the white splitting node, then the elimination lines.

**Transition:** “Now draw those three inequalities with the generators we already have.”

### 25 — The simplified diagram has no source decisions · 60 seconds

**Goal:** show the actual reduced generator diagram.

**Say:** “One cost source supplies t. Copy it to check nonnegativity, t at least d, and t plus one divided by two at least d. The white node adds one; the scalar halves that sum. The black merge requires both demand wires to agree. Only t is internal now: minimizing it recovers the maximum of the three affine lower bounds. The original source decisions have disappeared.”

**Delivery:** point to each order generator and the demand output. Do not claim a globally minimal diagram or a terminating rewrite algorithm.

**Transition:** “This example illustrates the result and its scope.”

### 26 — Result, scope, and remaining work · 60 seconds

**Goal:** make this a qualification, not an unqualified final-defense claim.

**Say:** “The draft presents GLP as a presentation of the LP value-function semantics: expressivity, soundness, and completeness together give the correspondence. The new extension is one generator and three axioms. Before the final defense, the priorities are to make the treatment of infeasible and unbounded cases fully explicit and to audit the syntactic epigraph derivations. The result is an equational calculus, not a benchmarked fast LP solver.”

**Delivery:** distinguish “presented in the draft” from “validated implementation.” The proposed priorities need advisor confirmation.

**Transition:** “The conceptual outcome is compact.”

### 27 — Optimization becomes part of the syntax · 35 seconds

**Goal:** leave a clear contribution and next action.

**Say:** “The key point is that composing these diagrams eliminates internal decisions, while the equations certify equality of the resulting value functions. Starting from a calculus for feasibility, a small extension brings linear objectives into the same compositional language. The next milestone is to finalize the semantic conventions and validate the epigraph proof. Quadratic extensions and computational tooling remain future directions. Thank you; I welcome your questions.”

**Delivery:** look away from the screen for the final sentences. Do not reopen the entire hierarchy.

## Recovery cuts if behind schedule

Protect slides 21–23. Do not recover by rushing the central proof.

| Slide | Compressed treatment | Approximate saving |
|---:|---|---:|
| 7 | State that the same wiring satisfies interchange; do not trace both groupings. | 20 s |
| 9 | “Quantales package the combine-and-aggregate laws; here those operations are addition and infimum.” | 20 s |
| 10 | “GLA, GAA, and GPA are the inherited lineage; GLP adds cost.” | 20 s |
| 12 | Explain the displayed bounds in one sentence and name FME. | 20 s |
| 19 | State one sentence each for expressivity and soundness. | 20 s |

If interrupted by a question, answer the narrow issue and return to the next transition. If extensive discussion is expected within the same 25-minute slot, shorten the opening/context in rehearsal; do not silently include Q&A in the proof's time budget.
