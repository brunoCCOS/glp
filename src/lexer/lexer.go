package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	src []rune
	pos int
}

func New(input string) *Lexer {
	return &Lexer{src: []rune(input)}
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.src) {
		return 0
	}
	return l.src[l.pos]
}

func (l *Lexer) advance() rune {
	r := l.src[l.pos]
	l.pos++
	return r
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.src) && unicode.IsSpace(l.src[l.pos]) {
		l.pos++
	}
}

func isIdentStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isIdentPart(r rune) bool {
	return isIdentStart(r) || unicode.IsDigit(r)
}

func (l *Lexer) Next() (Token, error) {
	l.skipWhitespace()
	start := l.pos

	if l.pos >= len(l.src) {
		return Token{Kind: EOF, Pos: start}, nil
	}

	r := l.peek()

	switch r {
	case '(':
		l.advance()
		return Token{Kind: LPAREN, Value: "(", Pos: start}, nil
	case ')':
		l.advance()
		return Token{Kind: RPAREN, Value: ")", Pos: start}, nil
	case ';':
		l.advance()
		return Token{Kind: COMP, Value: ";", Pos: start}, nil
	case '*':
		l.advance()
		return Token{Kind: TENSOR, Value: "*", Pos: start}, nil
	case ',':
		l.advance()
		return Token{Kind: COMMA, Value: ",", Pos: start}, nil
	}

	if unicode.IsDigit(r) {
		for l.pos < len(l.src) && unicode.IsDigit(l.src[l.pos]) {
			l.pos++
		}
		return Token{Kind: NUMBER, Value: string(l.src[start:l.pos]), Pos: start}, nil
	}

	if isIdentStart(r) {
		for l.pos < len(l.src) && isIdentPart(l.src[l.pos]) {
			l.pos++
		}
		word := string(l.src[start:l.pos])
		if kw, ok := keywords[word]; ok {
			return Token{Kind: kw, Value: word, Pos: start}, nil
		}
		return Token{Kind: IDENT, Value: word, Pos: start}, nil
	}

	return Token{}, fmt.Errorf("unexpected character %q at position %d", r, start)
}

func (l *Lexer) Tokenize() ([]Token, error) {
	var out []Token
	for {
		t, err := l.Next()
		if err != nil {
			return nil, err
		}
		out = append(out, t)
		if t.Kind == EOF {
			return out, nil
		}
	}
}
