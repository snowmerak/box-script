package grammar

import "math/big"

type Value struct {
	String      *string      `  @String`
	Int         *Integer     `| @@`
	Float       *Float       `| @@`
	Bool        *bool        `| @("true"|"false")`
	Char        *string      `| @Char`
	Array       *Array       `| @@`
	Expression  *Expression  `| "(" @@ ")"`
	Constructor *Constructor `| @@`
}

type Integer struct {
	Sign *bool    `@("-"|"+")?`
	Int  *big.Int `@Int`
}

type Float struct {
	Sign  *bool      `@("-"|"+")?`
	Float *big.Float `@Float`
}

type Array struct {
	Type   *Type    `"<" @@ ">"`
	Values []*Value `"[" @@ ("," @@)* "]"`
}

type Constructor struct {
	Name *string  `@Ident "("`
	Args []*Value `(@@ ("," @@)*)* ")"`
}
