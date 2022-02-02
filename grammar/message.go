package grammar

type Message struct {
	Name  *string        `@Ident "{"`
	Items []*MessageItem `@@* "}"`
}

type MessageItem struct {
	Name *string `@Ident`
	Type *Type   `@@`
}
