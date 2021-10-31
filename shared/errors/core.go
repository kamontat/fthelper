package errors

// Create error handler
func New() *Handler {
	return &Handler{
		errors: make([]error, 0),
	}
}
