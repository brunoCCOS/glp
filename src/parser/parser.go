package parser

import (
	"fmt"
	"strconv"

	"propc/src/lexer"
)

type Parser struct {
	toks []lexer.Token
	pos  int
}

func New(toks []lexer.Token) *Parser {
	return &Parser{toks: toks}
}

func Parse(input string) (Node, error) {
	toks, err := lexer.New(input).Tokenize()
	if err != nil {
		return nil, err
	}
	p := New(toks)
	n, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	if p.peek().Kind != lexer.EOF {
		return nil, fmt.Errorf("unexpected token %s at position %d", p.peek().Kind, p.peek().Pos)
	}
	return n, nil
}

func (p *Parser) peek() lexer.Token { return p.toks[p.pos] }

func (p *Parser) advance() lexer.Token {
	t := p.toks[p.pos]
	if p.pos < len(p.toks)-1 {
		p.pos++
	}
	return t
}

func (p *Parser) parseExpr() (Node, error) {
	left, err := p.parseTerm()
	if err != nil {
		return nil, err
	}
	for p.peek().Kind == lexer.COMP {
		p.advance()
		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}
		left = Comp{Left: left, Right: right}
	}
	return left, nil
}

func (p *Parser) parseTerm() (Node, error) {
	left, err := p.parseAtom()
	if err != nil {
		return nil, err
	}
	for p.peek().Kind == lexer.TENSOR {
		p.advance()
		right, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		left = Tensor{Left: left, Right: right}
	}
	return left, nil
}

func (p *Parser) parseAtom() (Node, error) {
	t := p.peek()
	switch t.Kind {
	case lexer.ID:
		p.advance()
		args, err := p.parseArgs(1)
		if err != nil {
			return nil, err
		}
		return Id{N: args[0]}, nil
	case lexer.SWAP:
		p.advance()
		args, err := p.parseArgs(2)
		if err != nil {
			return nil, err
		}
		return Swap{M: args[0], N: args[1]}, nil
	case lexer.IDENT:
		p.advance()
		var args []uint
		if p.peek().Kind == lexer.LPAREN {
			a, err := p.parseVariadicArgs()
			if err != nil {
				return nil, err
			}
			args = a
		}
		return Gen{Name: t.Value, Args: args}, nil
	case lexer.LPAREN:
		p.advance()
		n, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		if p.peek().Kind != lexer.RPAREN {
			return nil, fmt.Errorf("expected ')' at position %d, got %s", p.peek().Pos, p.peek().Kind)
		}
		p.advance()
		return n, nil
	}
	return nil, fmt.Errorf("unexpected token %s at position %d", t.Kind, t.Pos)
}

// parseVariadicArgs parses '(' NUMBER (',' NUMBER)* ')'. Used for generator
// arguments where the count is determined by the generator's declared
// parameter list rather than fixed in the grammar.
func (p *Parser) parseVariadicArgs() ([]uint, error) {
	if p.peek().Kind != lexer.LPAREN {
		return nil, fmt.Errorf("expected '(' at position %d, got %s", p.peek().Pos, p.peek().Kind)
	}
	p.advance()
	var out []uint
	for i := 0; ; i++ {
		if i > 0 {
			if p.peek().Kind != lexer.COMMA {
				break
			}
			p.advance()
		}
		t := p.peek()
		if t.Kind != lexer.NUMBER {
			return nil, fmt.Errorf("expected number at position %d, got %s", t.Pos, t.Kind)
		}
		v, err := strconv.ParseUint(t.Value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q at position %d", t.Value, t.Pos)
		}
		out = append(out, uint(v))
		p.advance()
	}
	if p.peek().Kind != lexer.RPAREN {
		return nil, fmt.Errorf("expected ')' at position %d, got %s", p.peek().Pos, p.peek().Kind)
	}
	p.advance()
	return out, nil
}

func (p *Parser) parseArgs(n int) ([]uint, error) {
	if p.peek().Kind != lexer.LPAREN {
		return nil, fmt.Errorf("expected '(' at position %d, got %s", p.peek().Pos, p.peek().Kind)
	}
	p.advance()
	out := make([]uint, 0, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			if p.peek().Kind != lexer.COMMA {
				return nil, fmt.Errorf("expected ',' at position %d, got %s", p.peek().Pos, p.peek().Kind)
			}
			p.advance()
		}
		t := p.peek()
		if t.Kind != lexer.NUMBER {
			return nil, fmt.Errorf("expected number at position %d, got %s", t.Pos, t.Kind)
		}
		v, err := strconv.ParseUint(t.Value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q at position %d", t.Value, t.Pos)
		}
		out = append(out, uint(v))
		p.advance()
	}
	if p.peek().Kind != lexer.RPAREN {
		return nil, fmt.Errorf("expected ')' at position %d, got %s", p.peek().Pos, p.peek().Kind)
	}
	p.advance()
	return out, nil
}
