package golang

import (
	"errors"

	"github.com/snowmerak/box-script/grammar"
)

func Statement(state *grammar.Statement) (string, error) {
	if state.Using != nil {
		return "import " + *state.Using, nil
	}
	if state.Enum != nil {
		return Enum(state.Enum)
	}
	return "", errors.New("unknown statement")
}
