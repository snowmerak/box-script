package grammar

type Enum struct {
	Name  *string     `@Ident "{"`
	Items []*EnumItem `(@@ (@@)*)* "}"`
}

type EnumItem struct {
	Name  *string `@Ident`
	Type  *Type   `@@ "="`
	Value *Value  `@@`
}
