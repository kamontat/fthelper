package fs

type Result string

const (
	EMPTY_RESULT   Result = "non-valid"
	MISSING_RESULT Result = "missing"
	VALID_RESULT   Result = "valid"
)
