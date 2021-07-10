package generators

import (
	"fmt"
	"os"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/exception"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
)

// CreateDir directory at output
func CreateDir(name string, fs ...features.Feature) runners.Runner {
	return New(name, features.Dependencies{
		features.KEY_OUTPUT: features.REQUIRE,
		features.MergeKeys(features.KEY_OUTPUT, features.KEY_FORMAT): features.OPTIONAL,
		features.KEY_USER:  features.OPTIONAL,
		features.KEY_CHMOD: features.OPTIONAL,
	}, func(param *ExecutorParameter, input models.Mapper) error {
		var outDir = buildDirectory(features.KEY_OUTPUT, input)

		var err = outDir.Build()
		if err != nil {
			return err
		}

		if user, ok := input.M(features.KEY_USER); ok {
			var cuid = os.Getuid()
			var cgid = os.Getgid()
			param.Logger.Debug("current permission is %d:%d", cuid, cgid)

			var uid = user.Io("uid", cuid)
			var gid = user.Io("gid", cgid)
			err = os.Chown(outDir.Abs(), uid, gid)
			if err != nil {
				exception.If(err, "generator", name).Print()
				return fmt.Errorf("sudo chown -R %d:%d '%s'", uid, gid, outDir.Abs())
			}
		}

		param.Logger.Info("create directory at %s", outDir.Abs())
		if a, ok := input.A(features.KEY_CHMOD); ok {
			var mode = a.(os.FileMode)
			err = outDir.Chmod(mode)
			if err != nil {
				exception.If(err, "generator", name).Print()
				return fmt.Errorf("sudo chmod %o '%s'", mode, outDir.Abs())
			}
		}

		return nil
	}, fs...)
}
