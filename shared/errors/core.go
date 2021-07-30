package errors

func New() *Handler {
	return &Handler{
		errors: make([]error, 0),
	}
}
