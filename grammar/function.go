package grammar

type Function struct {
	Name    *string      `@Ident "("`
	Args    []*Argument  `(@@ ("," @@)*)* ")"`
	Returns []*Type      `(@@ ("," @@)*)*`
	Body    []*Statement `{ @@ }`
}

type Argument struct {
	Name *string `@Ident`
	Type *Type   `@@`
}

type Return struct {
	Value *string `"return" @Ident`
}
