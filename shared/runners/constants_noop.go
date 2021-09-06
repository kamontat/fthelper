package runners

var NoValidator Validator = func(context *Context) error {
	return nil
}

var NoExecutor Executor = func(context *Context) error {
	return nil
}
