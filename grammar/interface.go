package grammar

type Interface struct {
	Name    *string   `@Ident "{"`
	Methods []*Method `@@* "}"`
}

type Method struct {
	Name    *string     `@Ident "("`
	Args    []*Argument `(@@ ("," @@)*)* ")"`
	Returns []*Type     `(@@ ("," @@)*)*`
}
