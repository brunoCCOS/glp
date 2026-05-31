package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"propc/src/assets"
	"propc/src/expr"
	"propc/src/lexer"
)

type file struct {
	Generators []lexer.Generator `json:"generators"`
}

// LoadDefault returns the generators embedded in the binary.
func LoadDefault() ([]lexer.Generator, error) {
	return parse(assets.GeneratorsJSON)
}

// LoadFile reads a generators.json file from disk.
func LoadFile(path string) ([]lexer.Generator, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parse(b)
}

func parse(b []byte) ([]lexer.Generator, error) {
	var f file
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("generators.json: %w", err)
	}
	seen := make(map[string]struct{}, len(f.Generators))
	for _, g := range f.Generators {
		if g.Value == "" {
			return nil, fmt.Errorf("generators.json: entry with empty name")
		}
		if _, dup := seen[g.Value]; dup {
			return nil, fmt.Errorf("generators.json: duplicate generator %q", g.Value)
		}
		seen[g.Value] = struct{}{}
		// For non-parametric generators, verify the arity/coarity expressions
		// evaluate as constants. Parametric ones are checked at call sites.
		if len(g.Params) == 0 {
			if _, err := expr.Eval(g.Arity, nil); err != nil {
				return nil, fmt.Errorf("generators.json: %s.arity: %w", g.Value, err)
			}
			if _, err := expr.Eval(g.Coarity, nil); err != nil {
				return nil, fmt.Errorf("generators.json: %s.coarity: %w", g.Value, err)
			}
		}
	}
	return f.Generators, nil
}
