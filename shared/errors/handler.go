package errors

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	Name   string
	errors []error
}

func (h *Handler) And(errs ...error) *Handler {
	for _, err := range errs {
		if err != nil {
			h.errors = append(h.errors, err)
		}
	}

	return h
}

func (h *Handler) Merge(nh *Handler) *Handler {
	h.errors = append(h.errors, nh.errors...)
	return h
}

func (h *Handler) HasError() bool {
	return len(h.errors) > 0
}

func (h *Handler) First() error {
	if h.HasError() {
		return h.errors[0]
	}
	return nil
}

func (h *Handler) String() string {
	var str strings.Builder

	if h.HasError() {
		str.WriteString("\n")
		for _, err := range h.errors {
			str.WriteString(fmt.Sprintf("%s\n", err.Error()))
		}
	} else {
		str.WriteString("no errors")
	}

	return str.String()
}

func (h *Handler) Error() error {
	if h.HasError() {
		return errors.New(h.String())
	}
	return nil
}

func (h *Handler) Panic() {
	if h.HasError() {
		panic(h.Error())
	}
}
