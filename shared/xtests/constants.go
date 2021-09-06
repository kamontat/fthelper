package xtests

type MustChecker string

const (
	MUST_BE_NIL         = "nil"
	MUST_NOT_BE_NIL     = "!nil"
	MUST_ERROR          = "error"
	MUST_NOT_ERROR      = "!error"
	MUST_EQUAL          = "equal"
	MUST_NOT_EQUAL      = "!equal"
	MUST_DEEP_EQUAL     = "deep-equal"
	MUST_EQUAL_STRING   = "equal-string"
	MUST_EQUAL_ERROR    = "equal-error"
	MUST_EQUAL_FLOAT    = "equal-float"
	MUST_EQUAL_REGEX    = "equal-regex"
	MUST_GREATER_THAN   = "greater-than"
	MUST_LESS_THAN      = "less-than"
	MUST_CONTAINS       = "contain"
	MUST_CONTAINS_ERROR = "contain-error"
)
