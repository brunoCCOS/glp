# propc

Compile PROP-style algebraic expressions into TikZ string diagrams.

Expressions are built from a small set of primitives — identity `id(n)`, swap
`swap(m,n)`, user-defined generators referenced by name — combined with
sequential composition `;` and tensor product `*`. The compiler typechecks the
expression against a generator table (arity / coarity) and emits a `tikzpicture`
that draws the diagram.

```
propc "copy ; mult"
```

## Install

```
go install ./cmd/propc
```

The binary is self-contained: the default generator table, the TikZ pic
definitions, and the styles are embedded at build time.

## Usage

```
propc [flags] [expression]

-i path        read expression from a file (default: stdin if no argument)
-o path        write output to a file (default: stdout)
-g path        use a custom generators.json (default: built-in)
--check        typecheck only; print "ok: A -> C" and exit
--standalone   wrap output in a compilable LaTeX document
```

Quick check:

```
propc --check "(id(1) * mult) ; mult"
# ok: 3 -> 1
```

Standalone document you can pipe straight into `pdflatex`:

```
propc --standalone "copy ; swap(1,1) ; mult" > diag.tex
pdflatex diag.tex
```

## Expression syntax

| form          | meaning                                                     |
|---------------|-------------------------------------------------------------|
| `id(n)`       | identity on `n` wires                                       |
| `swap(m,n)`   | symmetry σ: m+n → n+m                                       |
| `name`        | user-defined generator from the table                       |
| `f ; g`       | sequential composition (coarity of `f` = arity of `g`)      |
| `f * g`       | tensor product (parallel composition)                       |
| `(…)`         | grouping                                                    |

`*` binds tighter than `;`. Composition is left-associative.

## Generator table

A JSON file listing the user-defined generators:

```json
{
  "generators": [
    {"name": "mult", "arity": 2, "coarity": 1, "pic": "multiplication"},
    {"name": "copy", "arity": 1, "coarity": 2, "pic": "copy"}
  ]
}
```

`pic` is the TikZ pic name defined in `generator.tikz`. Each pic lives in the
frame `x ∈ [-1, 1]`, `y ∈ [-0.5, 0.5]` and exposes its input/output tips as
named coordinates `(-in-k)` / `(-out-k)`, top-first.

A default table and pic library are embedded; override either with `-g
generators.json` and by editing `src/assets/generator.tikz` (kept in sync with
the canonical copy under `material/`).

## Development

```
make test        # go test ./...
make build       # build the binary into ./bin/propc
make sync-assets # copy material/generator.tikz* into src/assets/
```

## Layout

```
cmd/propc/        CLI entry point
src/lexer/        tokenizer + generator metadata type
src/parser/       expression parser + AST
src/typecheck/    arity / coarity checker
src/codegen/      TikZ emitter
src/loader/       generators.json reader
src/assets/       embedded default tikz + json (do not edit by hand;
                  the canonical files live in material/)
```
