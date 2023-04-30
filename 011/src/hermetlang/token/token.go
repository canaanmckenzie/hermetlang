package token

type TokenType string
type Token struct {
	Type TokenType
	Literal string //ints or bytes is faster than string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"

	//Identifiers + literals
	IDENT	= "IDENT" //add,x,y,...
	INT	= "INT" //123

	//Operators
	ASSIGN	= "="
	PLUS	= "+"

	//Delimiters
	COMMA	  = ","
	SEMICOLON = ";"
	
	LPAREN	= "("
	RPAREN	= ")"
	LBRACE	= "{"
	RBRACE	= "}"

	//Keywords
	FUNCTION	= "FUNCTION"
	LET			= "LET"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}