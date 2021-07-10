package generators

import (
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/models"
)

func fsAddon(config *models.Config, file fs.FileSystem) models.Mapper {
	var result = make(models.Mapper)

	result.Set("addon.output.abs", file.Abs())
	result.Set("addon.output.dirpath", file.Dirpath())
	result.Set("addon.output.dirname", file.Dirname())
	result.Set("addon.output.filename", file.Basename())
	result.Set("addon.output.name", file.Name())

	return result
}
