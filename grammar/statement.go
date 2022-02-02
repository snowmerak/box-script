package grammar

type Statement struct {
	Using     *string    `  "using" @String`
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

const (
	Int8    = "int8"
	Int16   = "int16"
	Int32   = "int32"
	Int64   = "int64"
	Uint8   = "uint8"
	Uint16  = "uint16"
	Uint32  = "uint32"
	Uint64  = "uint64"
	Byte    = Uint8
	Rune    = Int32
	Float32 = "float32"
	Float64 = "float64"
	String  = "string"
	Char    = "char"
	Bool    = "bool"
)
