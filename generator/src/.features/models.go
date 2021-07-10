package features

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
)

type Feature interface {
	Name() string
	Depends() Dependencies
	Run(input models.Mapper) runners.Runner

	String() string
}

type FeatureImpl struct {
	name         string
	Dependencies Dependencies

	executor Executor

	Logger *loggers.Logger
}

func (r *FeatureImpl) Name() string {
	return r.name
}

func (r *FeatureImpl) Depends() Dependencies {
	return r.Dependencies
}

func (f *FeatureImpl) Run(input models.Mapper) runners.Runner {
	return runners.NewRunner(f.name, runners.NoValidate, func(i *runners.SingleInfo) error {
		d, e := f.executor(f, i.Input().(models.Mapper))
		if e != nil {
			return e
		}

		i.Out(d)
		return nil
	}).Input(input)
}

func (f *FeatureImpl) String() string {
	return fmt.Sprintf("Function <%s>();", f.Name())
}
