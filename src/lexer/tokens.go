package lexer

type TokenKind int

const (
	EOF TokenKind = iota
	LPAREN
	RPAREN
	COMMA

	// operators
	COMP   // ;
	TENSOR // *

	// literals
	NUMBER

	// keywords
	ID
	SWAP

	// user-defined generator name; resolved against the generator table later
	IDENT
)

type Token struct {
	Kind  TokenKind
	Value string
	Pos   int
}

var keywords = map[string]TokenKind{
	"id":   ID,
	"swap": SWAP,
}

func (k TokenKind) String() string {
	switch k {
	case EOF:
		return "EOF"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case COMMA:
		return "COMMA"
	case NUMBER:
		return "NUMBER"
	case COMP:
		return "COMP"
	case TENSOR:
		return "TENSOR"
	case ID:
		return "ID"
	case SWAP:
		return "SWAP"
	case IDENT:
		return "IDENT"
	}
	return "UNKNOWN"
}
