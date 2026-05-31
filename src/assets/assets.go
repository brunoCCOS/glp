// Package assets embeds the default tikz pic definitions, tikz styles, and
// generator index so the propc binary is self-contained. Keep these files in
// sync with material/ at the repo root (which the LaTeX document also reads).
package assets

import _ "embed"

//go:embed generator.tikz
var GeneratorTikz string

//go:embed generator.tikzstyles
var GeneratorTikzstyles string

//go:embed generators.json
var GeneratorsJSON []byte
