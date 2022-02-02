package grammar

type Statement struct {
	Using     *Using     `"using" @@`
	Let       *Let       `| "let" @@`
	Enum      *Enum      `| "enum" @@`
	Message   *Message   `| "message" @@`
	Function  *Function  `| "func" @@`
	Actor     *Actor     `| "actor" @@`
	Interface *Interface `| "interface" @@`
	Return    *Return    `| "return" @@`
}

type Type struct {
	Name []string `@Ident ("." @Ident)*`
}

type Using struct {
	Name *string `"using" @String`
}
