package grammar

type Expression struct {
	Left  *Value   `@@?`
	Right []*Right `@@*`
}

type Right struct {
	Op    *Op    `@@`
	Right *Value `@@?`
}

type Op struct {
	Contents []string `@("+" | "-" | "*" | "/" | "%" | "^" | "&" | "|" | "<" | ">" | "=" | "!")+`
}
