package errors

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Handler use for checking error
type Handler struct {
	// number of total error that check on this handlers. including error that is nil
	total  int
	errors []error
}

// Total will get number of errors
func (h *Handler) Total() int {
	return h.total
}

// Length will get length of errors
func (h *Handler) Length() int {
	return len(h.errors)
}

// And will add error to handler
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

// Merge will merge error to handler
func (h *Handler) Merge(nh *Handler) *Handler {
	h.errors = append(h.errors, nh.errors...)
	h.total += nh.total
	return h
}

// HasError will check whether handler has error or not
func (h *Handler) HasError() bool {
	return len(h.errors) > 0
}

// First will get the first error in handler
func (h *Handler) First() error {
	if h.HasError() {
		return h.errors[0]
	}
	return nil
}

// String use for converting object data to string
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

// Error will convert handler to error if handler contains any error
func (h *Handler) Error() error {
	if h.HasError() {
		return errors.New(h.String())
	}
	return nil
}

// Exit will exit the program if handler contain any error
func (h *Handler) Exit(n int) {
	if h.HasError() {
		os.Exit(n)
	}
}

// Same as Exit but always return exit code 1
func (h *Handler) Exit1() {
	h.Exit(1)
}

// Panic program if error
func (h *Handler) Panic() {
	if h.HasError() {
		panic(h.Error())
	}
}
