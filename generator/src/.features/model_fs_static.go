package features

import (
	"github.com/kamontat/fthelper/shared/fs"
)

func FileSystemStatic(featureName string, def *fs.PathModel) Feature {
	return FileSystemOConfig(featureName, fs.DefaultSearcher(), def)
}
