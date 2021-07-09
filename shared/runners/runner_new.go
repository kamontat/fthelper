package runners

func NewRunner(name string, vl Validator, ec Executor) Runner {
	return &RunnerImpl{
		validator: vl,
		executor:  ec,
		info:      NewSingleInfo(name),
	}
}
