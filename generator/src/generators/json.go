package generators

import (
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Json(data maps.Mapper, fsConfig maps.Mapper) runners.Runner {
	var name = data.So("display", "default")
	var log = loggers.Get("generator", "json", name)

	return runners.NewRunner(name, func(i *runners.SingleInfo) error {
		var data = i.Input().(maps.Mapper)

		templates, err := fs.Build(data.Si("templates"), fsConfig)
		if err != nil {
			log.Error("cannot get template information")
			return err
		}

		output, err := fs.Build(data.Si("output"), fsConfig)
		if err != nil {
			log.Error("cannot get output information")
			return err
		}

		content, err := configs.LoadConfigFromFileSystem(templates.Multiple(), fsConfig.Mi("variable"), data.Mi("merger"))
		if err != nil {
			log.Error("cannot load template as json")
			return err
		}

		var outfile = output.Single()
		err = outfile.Build()
		if err != nil {
			log.Error("cannot build output file")
			return err
		}

		json, err := maps.ToFormatJson(content)
		if err != nil {
			log.Error("cannot create formatted json")
			return err
		}

		return outfile.Write(json)
	}, func(i *runners.SingleInfo) error {
		return nil
	}).Input(data)
}
