// Package expr evaluates tiny arithmetic expressions over named uint
// variables. Supported operators are + - * / and parentheses, with the
// conventional precedence and left-associativity. Subtraction that would
// underflow uint and division by zero are reported as errors.
package expr

import (
	"fmt"
	"strings"
	"unicode"
)

// Eval evaluates s against the given variable bindings.
func Eval(s string, vars map[string]uint) (uint, error) {
	p := &state{src: []rune(strings.TrimSpace(s)), vars: vars}
	v, err := p.parseExpr()
	if err != nil {
		return 0, err
	}
	p.skip()
	if p.pos < len(p.src) {
		return 0, fmt.Errorf("unexpected character %q at position %d", p.src[p.pos], p.pos)
	}
	return v, nil
}

type state struct {
	src  []rune
	pos  int
	vars map[string]uint
}

func (p *state) skip() {
	for p.pos < len(p.src) && unicode.IsSpace(p.src[p.pos]) {
		p.pos++
	}
}

func (p *state) parseExpr() (uint, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}
	for {
		p.skip()
		if p.pos >= len(p.src) {
			return left, nil
		}
		c := p.src[p.pos]
		if c != '+' && c != '-' {
			return left, nil
		}
		p.pos++
		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}
		if c == '+' {
			left += right
		} else {
			if right > left {
				return 0, fmt.Errorf("subtraction underflow at position %d", p.pos)
			}
			left -= right
		}
	}
}

func (p *state) parseTerm() (uint, error) {
	left, err := p.parseFactor()
	if err != nil {
		return 0, err
	}
	for {
		p.skip()
		if p.pos >= len(p.src) {
			return left, nil
		}
		c := p.src[p.pos]
		if c != '*' && c != '/' {
			return left, nil
		}
		p.pos++
		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}
		if c == '*' {
			left *= right
		} else {
			if right == 0 {
				return 0, fmt.Errorf("division by zero at position %d", p.pos)
			}
			left /= right
		}
	}
}

func (p *state) parseFactor() (uint, error) {
	p.skip()
	if p.pos >= len(p.src) {
		return 0, fmt.Errorf("unexpected end of expression")
	}
	c := p.src[p.pos]
	if c == '(' {
		p.pos++
		v, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		p.skip()
		if p.pos >= len(p.src) || p.src[p.pos] != ')' {
			return 0, fmt.Errorf("expected ')' at position %d", p.pos)
		}
		p.pos++
		return v, nil
	}
	if unicode.IsDigit(c) {
		start := p.pos
		for p.pos < len(p.src) && unicode.IsDigit(p.src[p.pos]) {
			p.pos++
		}
		var n uint
		for _, r := range p.src[start:p.pos] {
			n = n*10 + uint(r-'0')
		}
		return n, nil
	}
	if unicode.IsLetter(c) || c == '_' {
		start := p.pos
		for p.pos < len(p.src) {
			r := p.src[p.pos]
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
				break
			}
			p.pos++
		}
		name := string(p.src[start:p.pos])
		v, ok := p.vars[name]
		if !ok {
			return 0, fmt.Errorf("unknown variable %q", name)
		}
		return v, nil
	}
	return 0, fmt.Errorf("unexpected character %q at position %d", c, p.pos)
}
