package generators

import (
	"io"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
)

func CopyDir(name string, feature ...features.Feature) runners.Runner {
	return New(name, features.Dependencies{
		features.KEY_CONFIG: features.REQUIRE,
		features.KEY_INPUT:  features.REQUIRE,
		features.MergeKeys(features.KEY_INPUT, features.KEY_FORMAT): features.OPTIONAL,
		features.KEY_OUTPUT: features.REQUIRE,
		features.MergeKeys(features.KEY_OUTPUT, features.KEY_FORMAT): features.OPTIONAL,
	}, func(param *ExecutorParameter, input models.Mapper) error {
		var infs = buildDirectory(features.KEY_INPUT, input)
		var outfs = buildDirectory(features.KEY_OUTPUT, input)

		var infiles, err = infs.ReadDir()
		if err != nil {
			return err
		}

		var total int64 = 0
		for _, infile := range infiles {
			inReader, err := infile.Reader()
			if err != nil {
				return err
			}

			relative := infile.Relative(infs)
			outfile := fs.NextFile(outfs, relative)
			err = outfile.Build()
			if err != nil {
				return err
			}

			outWriter, err := outfile.Writer()
			if err != nil {
				return err
			}

			n, err := io.Copy(outWriter, inReader)
			if err != nil {
				return err
			}

			total += n
			param.Logger.Debug("coping '%s' -> '%s' (%d)", infile.Abs(), outfile.Abs(), n)
		}

		param.Logger.Info("copying to %s (%d)", outfs.Abs(), total)
		return nil
	}, feature...)
}
