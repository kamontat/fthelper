package runners

// Runner is base interface
type Runner interface {
	Input(i interface{}) Runner
	Information() Information

	Validate() error
	Run() error
}
