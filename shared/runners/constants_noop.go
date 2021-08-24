package runners

func NoValidator(context *Context) error {
	return nil
}

func NoExecutor(context *Context) error {
	return nil
}
