package runners

import (
	"github.com/kamontat/fthelper/shared/loggers"
)

type Context struct {
	Disabled bool
	Logger   *loggers.Logger

	input  interface{}
	output interface{}
}

func (c *Context) In(i interface{}) {
	c.input = i
}

func (c *Context) Input() interface{} {
	return c.input
}

func (c *Context) Out(o interface{}) {
	c.output = o
}

func (c *Context) Output() interface{} {
	return c.output
}

func NewContext(name string) *Context {
	return &Context{
		Disabled: false,
		Logger:   loggers.Get("runner", name),
	}
}
