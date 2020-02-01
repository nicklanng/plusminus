package plusminus

import "errors"

var (
	ErrSyntaxError = errors.New("syntax error")
)

type expr interface {
	toString() string
}
