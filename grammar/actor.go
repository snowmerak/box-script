package grammar

type Actor struct {
	Name          *string         `@Ident`
	InterfaceName []string        `("impl" @Ident (@Ident)*)*`
	Fields        []*Field        `"{" @@*`
	Methods       []*MemberMethod `@@* "}"`
}

type Field struct {
	Name *string `@Ident`
	Type *Type   `@@`
}

type MemberMethod struct {
	Method *Method      `@@`
	Body   []*Statement `{ @@* }`
}
