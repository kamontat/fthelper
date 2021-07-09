package runners

type Executor func(i *SingleInfo) error
type Validator func(i *SingleInfo) error
