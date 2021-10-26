package errors

import (
	"errors"
	"fmt"
	"os"
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
		h.total += 1
		if err != nil {
			h.errors = append(h.errors, err)
		}
	}

	return h
}

// AndD will help user to not need variable for checking error
// For example if method return (string, error)
// instead create variable and pass only error to And method you can pass entirely to this function
// and use error handler to handle error instead
func (h *Handler) AndD(data interface{}, err error) (interface{}, *Handler) {
	return data, h.And(err)
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
		str.WriteString(fmt.Sprintf("found '%d' errors (total=%d)\n", h.Length(), h.Total()))
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

func (h *Handler) Exit(n int) {
	if h.HasError() {
		os.Exit(n)
	}
}

func (h *Handler) Exit1() {
	h.Exit(1)
}

func (h *Handler) Panic() {
	if h.HasError() {
		panic(h.Error())
	}
}
