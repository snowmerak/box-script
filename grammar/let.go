package grammar

type Let struct {
	Mutable *bool   `@("mut")?`
	Name    *string `@Ident`
	Type    *Type   `@@`
	Value   *Value  `"=" @@`
}
