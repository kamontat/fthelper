package errors

import "github.com/kamontat/fthelper/shared/utils"

func New(name ...string) *Handler {
	return &Handler{
		Name:   utils.JoinString(":", name...),
		errors: make([]error, 0),
	}
}
