package golang

import (
	"errors"
	"strings"

	"github.com/snowmerak/box-script/grammar"
)

func Enum(enum *grammar.Enum) (string, error) {
	sb := strings.Builder{}

	sb.WriteString("type " + *enum.Name + " interface{}\n\n")

	sb.WriteString("const (\n")
	for _, item := range enum.Items {
		sb.WriteString("\t" + *item.Name + " " + strings.Join(item.Type.Name, ".") + " = ")
		if item.Value == nil {
			return "", errors.New("enum item value is nil")
		}
		value := Value(item.Type, item.Value)
		sb.WriteString(value + "\n")
	}
	sb.WriteString(")\n")
	return sb.String(), nil
}
