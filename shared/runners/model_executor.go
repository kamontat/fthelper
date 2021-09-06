package runners

type Validator func(context *Context) error

type Executor func(context *Context) error
