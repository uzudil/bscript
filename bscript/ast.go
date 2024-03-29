package bscript

import (
	"fmt"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"

	"strings"
)

type Program struct {
	Pos lexer.Position

	TopLevel []*TopLevel `( @@ )*`
}

type TopLevel struct {
	Pos lexer.Position

	Const   *Const   `( @@ ";"`
	Command *Command `| @@ )`
}

type Const struct {
	Pos lexer.Position

	Name  string      `"const" @Ident`
	Value *Expression `"=" @@`
}

type Fun struct {
	Pos lexer.Position

	Name     string      `"def" @Ident "("`
	Params   []*FunParam `( @@ ( "," @@ )* )*`
	Commands []*Command  `")" "{" ( @@ )* "}"`
}

type AnonFun struct {
	Pos lexer.Position

	Params        []*FunParam `( "(" ( @@ ( "," @@ )* )* ")" "=" ">"`
	SingleParam   *FunParam   `| @@ "=" ">" )`
	Commands      []*Command  `( "{" ( @@ )* "}"`
	SingleCommand *Expression `| @@ )`
}

type FunParam struct {
	Pos lexer.Position

	Name         string      `@Ident`
	DefaultValue *Expression `( "=" @@ )?`
}

type Command struct {
	Pos lexer.Position

	Remark   *Remark   `(   @@ `
	Del      *Del      `  | @@ ";" `
	Return   *Return   `  | @@ ";" `
	If       *If       `  | @@ `
	While    *While    `  | @@ `
	Fun      *Fun      `  | @@`
	Variable *Variable `  | @@ ";"`
	Let      *Let      `  | @@ ";" )`
}

type Del struct {
	Pos lexer.Position

	Variable *Variable `"del" @@`
}

type While struct {
	Pos lexer.Position

	Condition *Expression `"while" "(" @@ ")" "{"`
	Commands  []*Command  `( @@ )* "}"`
}

type If struct {
	Pos lexer.Position

	Condition    *Expression `"if" "(" @@ ")" "{"`
	Commands     []*Command  `( @@ )* "}"`
	ElseIf       []*ElseIf   `( @@ )*`
	ElseCommands []*Command  `( "else" "{" ( @@ )* "}" )?`
}

type ElseIf struct {
	Condition *Expression `"else" "if" "(" @@ ")" "{"`
	Commands  []*Command  `( @@ )* "}"`
}

type Remark struct {
	Pos lexer.Position

	Comment string `@Comment`
}

type Let struct {
	Pos lexer.Position

	Variable *Variable   `@@`
	LetOp    *LetOp      `":" @@`
	Value    *Expression `@@`
}

type LetOp struct {
	Pos lexer.Position

	Assign *string `@"="`
	Add    *string `| @"+"`
	Sub    *string `| @"-"`
	Mul    *string `| @"*"`
	Div    *string `| @"/"`
}

type Return struct {
	Pos lexer.Position

	Value *Expression `"return" ( @@ )?`
}

type Operator string

func (o *Operator) Capture(s []string) error {
	*o = Operator(strings.Join(s, ""))
	return nil
}

type Value struct {
	Pos lexer.Position

	Array         *Array        ` @@`
	Map           *Map          `| @@`
	AnonFun       *AnonFun      `| @@`
	Null          *string       `| @"null"`
	Number        *SignedNumber `| @@`
	Boolean       *string       `| @("true" | "false")`
	Variable      *Variable     `| @@`
	String        *string       `| @String`
	Subexpression *Expression   `| "(" @@ ")"`
}

type SignedNumber struct {
	Pos lexer.Position

	Sign   *string `@("+" | "-")?`
	Number float64 `@Number`
}

type Variable struct {
	Pos lexer.Position

	Variable string            `@Ident`
	Suffixes []*VariableSuffix `( @@ )*`
}

type VariableSuffix struct {
	Pos lexer.Position

	Index      *ArrayIndex `@@`
	MapKey     *MapKey     `| @@`
	CallParams *CallParams `| @@`
}

type MapKey struct {
	Pos lexer.Position
	Key string `"." @Ident`
}

type CallParams struct {
	Pos lexer.Position

	Args []*Expression `"(" [ @@ { "," @@ } ] ")"`
}

type ArrayIndex struct {
	Pos   lexer.Position
	Index *Expression `"[" @@ "]"`
}

type Array struct {
	Pos lexer.Position

	LeftValue   *Expression   `"[" @@*`
	RightValues []*Expression `( "," @@ )* ","? "]"`
}

type Map struct {
	Pos lexer.Position

	LeftNameValuePair   *NameValuePair   `"{" @@*`
	RightNameValuePairs []*NameValuePair `( "," @@ )* ","? "}"`
}

type NameValuePair struct {
	Pos lexer.Position

	Name  string      `(@String | @Ident) ":"`
	Value *Expression `@@`
}

type Factor struct {
	Pos lexer.Position

	Base     *Value `@@`
	Exponent *Value `[ "^" @@ ]`
}

type OpFactor struct {
	Pos lexer.Position

	Operator Operator `@("*" | "/" | "%")`
	Factor   *Factor  `@@`
}

type Term struct {
	Pos lexer.Position

	Left  *Factor     `@@`
	Right []*OpFactor `{ @@ }`
}

type OpTerm struct {
	Pos lexer.Position

	Operator Operator `@("+" | "-")`
	Term     *Term    `@@`
}

type Cmp struct {
	Pos lexer.Position

	Left  *Term     `@@`
	Right []*OpTerm `{ @@ }`
}

type BoolCmp struct {
	Pos lexer.Position

	Positive *Cmp `@@`
	Negative *Cmp `|"!" @@`
}

type OpCmp struct {
	Pos lexer.Position

	Operator Operator `@("=" | "<" "=" | ">" "=" | "<" | ">" | "!" "=")`
	Cmp      *BoolCmp `@@`
}

type BoolTerm struct {
	Pos lexer.Position

	Left  *BoolCmp `@@`
	Right []*OpCmp `{ @@ }`
}

type OpBoolTerm struct {
	Pos lexer.Position

	Operator Operator  `@("&" "&" | "|" "|")`
	Right    *BoolTerm `@@`
}

type Expression struct {
	Pos lexer.Position

	BoolTerm   *BoolTerm     `@@`
	OpBoolTerm []*OpBoolTerm `{ @@ }`
}

var (
	benjiLexer = lexer.Must(ebnf.New(`
		Comment = "#" { "\u0000"…"\uffff"-"\n"-"\r" } .
		Ident = (alpha | "_") { "_" | alpha | digit } .
		String = "\"" { "\u0000"…"\uffff"-"\""-"\\" | "\\" any } "\"" .
		Number = ("." | digit) { "." | digit } .
		Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" .
		Whitespace = ( " " | "\t" | "\n" | "\r" ) { " " | "\t" | "\n" | "\r" } .

		alpha = "a"…"z" | "A"…"Z" .
		digit = "0"…"9" .
		any = "\u0000"…"\uffff" .
	`))

	Parser = participle.MustBuild(&Program{},
		participle.Lexer(benjiLexer),
		participle.CaseInsensitive("Ident"),
		participle.Unquote("String"),
		participle.UseLookahead(100),
		participle.Elide("Whitespace"),
	)

	CommandParser = participle.MustBuild(&Command{},
		participle.Lexer(benjiLexer),
		participle.CaseInsensitive("Ident"),
		participle.Unquote("String"),
		participle.UseLookahead(100),
		participle.Elide("Whitespace"),
	)
)

func ParseString(command string, ast Evaluatable, ctx *Context) (interface{}, error) {
	err := CommandParser.ParseString(command, ast)
	if err != nil {
		// try a few things to make it compile
		for _, c := range []string{
			fmt.Sprintf("return %s;", command),
			fmt.Sprintf("%s;", command),
		} {
			err2 := CommandParser.ParseString(c, ast)
			if err2 == nil {
				err = nil
				break
			}
		}
	}
	if err != nil {
		return nil, err
	}
	// repr.Println(ast)
	return ast.Evaluate(ctx)
}
