package xtests

type MustChecker string

const (
	MUST_ERROR          = "error"
	MUST_NOT_ERROR      = "!error"
	MUST_EQUAL          = "equal"
	MUST_NOT_EQUAL      = "!equal"
	MUST_DEEP_EQUAL     = "deep-equal"
	MUST_EQUAL_STRING   = "equal-string"
	MUST_EQUAL_ERROR    = "equal-error"
	MUST_CONTAINS_ERROR = "contain-error"
)
