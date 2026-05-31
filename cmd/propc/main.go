// propc compiles PROP-style algebraic expressions into TikZ diagrams.
//
// Usage:
//
//	propc [flags] [expression]
//
// If no expression is given, stdin is read. Output goes to stdout unless -o is
// set. With --standalone, the emitted file is a self-contained .tex document
// that can be compiled directly with pdflatex.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"propc/src/assets"
	"propc/src/codegen"
	"propc/src/lexer"
	"propc/src/loader"
	"propc/src/parser"
	"propc/src/typecheck"
)

func main() {
	var (
		inPath     = flag.String("i", "", "input file containing the expression (default: stdin)")
		outPath    = flag.String("o", "", "output file (default: stdout)")
		gensPath   = flag.String("g", "", "generators.json (default: built-in)")
		checkOnly  = flag.Bool("check", false, "typecheck only, do not emit tikz")
		standalone = flag.Bool("standalone", false, "wrap output in a compilable LaTeX document")
	)
	flag.Parse()

	if err := run(*inPath, *outPath, *gensPath, *checkOnly, *standalone, flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, "propc:", err)
		os.Exit(1)
	}
}

func run(inPath, outPath, gensPath string, checkOnly, standalone bool, args []string) error {
	expr, err := readExpr(inPath, args)
	if err != nil {
		return err
	}
	gens, err := loadGens(gensPath)
	if err != nil {
		return err
	}
	env := typecheck.NewEnv(gens)

	node, err := parser.Parse(expr)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	sig, err := typecheck.Check(node, env)
	if err != nil {
		return fmt.Errorf("typecheck: %w", err)
	}
	if checkOnly {
		fmt.Fprintf(os.Stderr, "ok: %d -> %d\n", sig.Arity, sig.Coarity)
		return nil
	}

	tikz, err := codegen.Generate(node, env)
	if err != nil {
		return fmt.Errorf("codegen: %w", err)
	}
	out := tikz
	if standalone {
		out = wrapStandalone(tikz)
	}
	return writeOut(outPath, out)
}

func readExpr(path string, args []string) (string, error) {
	if path != "" {
		b, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	if len(strings.TrimSpace(string(b))) == 0 {
		return "", fmt.Errorf("no expression provided (use -i, an argument, or stdin)")
	}
	return string(b), nil
}

func loadGens(path string) ([]lexer.Generator, error) {
	if path == "" {
		return loader.LoadDefault()
	}
	return loader.LoadFile(path)
}

func writeOut(path, content string) error {
	if path == "" {
		_, err := io.WriteString(os.Stdout, content)
		return err
	}
	return os.WriteFile(path, []byte(content), 0o644)
}

func wrapStandalone(tikz string) string {
	var sb strings.Builder
	sb.WriteString("\\documentclass[tikz,border=2pt]{standalone}\n")
	sb.WriteString("\\usepackage{tikz}\n")
	sb.WriteString("\\usetikzlibrary{shapes,calc}\n")
	sb.WriteString("\\makeatletter\n")
	sb.WriteString(assets.GeneratorTikzstyles)
	sb.WriteString("\n")
	sb.WriteString(assets.GeneratorTikz)
	sb.WriteString("\n\\makeatother\n")
	sb.WriteString("\\begin{document}\n")
	sb.WriteString(tikz)
	sb.WriteString("\\end{document}\n")
	return sb.String()
}
