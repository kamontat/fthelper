package runners

type Status string

const (
	INITIAL  Status = "initial"
	DISABLED Status = "disabled"
	SUCCESS  Status = "success"
	ERROR    Status = "error"
	INVALID  Status = "invalid"
)
