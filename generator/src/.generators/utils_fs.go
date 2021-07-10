package generators

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

func buildTemplateData(input models.Mapper, p *fs.PathModel) models.Mapper {
	var empty = make(models.Mapper)

	var fileext = path.Ext(p.Filename)

	if raw, ok := input.A(features.KEY_KEY); ok {
		if key, ok := raw.(*configs.KeyModel); ok {
			empty.Set("env", string(key.Env))
			empty.Set("key", string(key.Key))
			empty.Set("subkey", string(key.Subkey))
		}
	}

	var config = features.GetConfig(input)
	if common, err := config.ToCommon(); err == nil {
		empty.Set("cluster", common.Cluster)
	}

	empty.Set("path.basedir", p.Basedir)
	empty.Set("path.dirname", p.Dirname)
	empty.Set("path.filename", p.Filename)
	if p.Filename != "" {
		if fileext != "" {
			empty.Set("filename.ext", fileext[1:])
		}
		empty.Set("filename.name", strings.Replace(p.Filename, fileext, "", 1))
	}

	return empty
}

func buildFormatedPathModel(format *fs.FormatModel, input models.Mapper, p *fs.PathModel) *fs.PathModel {
	var err error
	var empty = p.Clone()
	var data = buildTemplateData(input, p)

	if format.Basedir != "" {
		var basedir bytes.Buffer
		err = xtemplates.BuildText("basedir-path", format.Basedir, data, &basedir)
		if err == nil {
			empty.Basedir = basedir.String()
		}
	}

	if format.Dirname != "" {
		var dirname bytes.Buffer
		err = xtemplates.BuildText("dirname-path", format.Dirname, data, &dirname)
		if err == nil {
			empty.Dirname = dirname.String()
		}
	}

	if format.Filename != "" {
		var filename bytes.Buffer
		err = xtemplates.BuildText("filename-path", format.Filename, data, &filename)
		if err == nil {
			empty.Filename = filename.String()
		}
	}

	return empty
}

func buildDirectory(featureName string, input models.Mapper) fs.FileSystem {
	var config = input.Ai(featureName).(*fs.PathModel)
	if config.Type == fs.FILE {
		panic(fmt.Errorf("we will try to create directory path from file type (%v)", config))
	}

	if input.Has(features.MergeKeys(featureName, features.KEY_FORMAT)) {
		var f = input.Ao(features.MergeKeys(featureName, features.KEY_FORMAT), fs.EmptyFormatModel).(*fs.FormatModel)
		config = buildFormatedPathModel(f, input, config)
	}

	f, e := fs.NewFileSystem(config, nil)
	if e != nil {
		panic(e)
	}

	return f
}

func buildFile(featureName string, input models.Mapper) fs.FileSystem {
	var config = input.Ai(featureName).(*fs.PathModel)
	if config.Type == fs.DIRECTORY {
		panic(fmt.Errorf("we will try to create file path from directory type (%v)", config))
	}

	if input.Has(features.MergeKeys(featureName, features.KEY_FORMAT)) {
		var f = input.Ao(features.MergeKeys(featureName, features.KEY_FORMAT), fs.EmptyFormatModel).(*fs.FormatModel)
		config = buildFormatedPathModel(f, input, config)
	}

	f, e := fs.NewFileSystem(config, nil)
	if e != nil {
		panic(e)
	}

	return f
}

func buildFiles(featureName string, input models.Mapper) []fs.FileSystem {
	var config = input.Ai(featureName).(*fs.PathModel)
	if input.Has(features.MergeKeys(featureName, features.KEY_FORMAT)) {
		var f = input.Ao(features.MergeKeys(featureName, features.KEY_FORMAT), fs.EmptyFormatModel).(*fs.FormatModel)
		config = buildFormatedPathModel(f, input, config)
	}

	f, e := fs.NewFileSystems(make(map[string]bool), config, nil)
	if e != nil {
		panic(e)
	}

	return f
}
