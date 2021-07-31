package errors

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	// number of total error that check on this handlers. including error that is nil
	total  int
	errors []error
}

func (h *Handler) Total() int {
	return h.total
}

func (h *Handler) Length() int {
	return len(h.errors)
}

func (h *Handler) And(errs ...error) *Handler {
	for _, err := range errs {
		if err != nil {
			h.errors = append(h.errors, err)
		} else {
			h.total += 1
		}
	}

	return h
}

func (h *Handler) AndD(_ interface{}, err error) *Handler {
	return h.And(err)
}

func (h *Handler) Merge(nh *Handler) *Handler {
	h.errors = append(h.errors, nh.errors...)
	h.total += nh.total
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
		str.WriteString(fmt.Sprintf("found '%d' errors (%d)\n", h.Length(), h.Total()))
		for _, err := range h.errors {
			str.WriteString(fmt.Sprintf("- %s\n", err.Error()))
		}
	} else {
		str.WriteString("not found any errors")
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
