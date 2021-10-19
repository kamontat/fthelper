package configs

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/utils"
)

type Builder struct {
	name     string
	config   maps.Mapper
	override maps.Mapper
	strategy maps.Mapper
	logger   *loggers.Logger
}

func (b *Builder) Strategy(strategy maps.Mapper) *Builder {
	b.strategy = strategy
	return b
}

func (b *Builder) OverrideStrings(strings []string) *Builder {
	var m = maps.New()
	for _, str := range strings {
		if k, v, ok := ParseOverride(str); ok {
			m.Set(k, v)
		}
	}
	return b.OverrideMap(m)
}

func (b *Builder) OverrideMap(m maps.Mapper) *Builder {
	m.ForEach(func(key string, value interface{}) {
		b.override.Set(key, value)
	})
	return b
}

func (b *Builder) log(t string, key, value interface{}, def interface{}) {
	var oldStr = fmt.Sprint(def)
	var newStr = fmt.Sprint(value)
	if oldStr != newStr {
		var oldMask = utils.MaskString(oldStr, utils.MEDIUM)
		var newMask = utils.MaskString(newStr, utils.MEDIUM)

		b.logger.Debug(fmt.Sprintf("override '%s' from '%v [%T] -> '%v [%T] (%s)", key, oldMask, def, newMask, value, t))
	}
}

func (b *Builder) updateResult(t string, base, input maps.Mapper) {
	if !input.IsEmpty() {
		b.logger.Debug("found some config in %s", t)
	}

	input.ForEach(func(key string, value interface{}) {
		var old, err = base.Get(key) // try to get old data
		var result = value

		if str, ok := datatype.ToString(value); ok {
			if err == nil {
				result = datatype.ConvertStringTo(str, old)
			} else {
				result = datatype.ConvertString(str)
			}
		}

		b.log(t, key, result, old) // log resule
		base.Set(key, result)      // update base mapper
	})
}

func (b *Builder) Build(environments []string) (maps.Mapper, error) {
	var result = maps.Merger(maps.New()).Add(b.config).SetConfig(b.strategy).Merge()
	b.logger.Debug("base configuration is %v", result)

	var args = make([]string, 0)
	for _, v := range result.Mi("internal").Ai("args") {
		args = append(args, v.(string))
	}
	b.OverrideStrings(args)

	configs, err := fs.Build(result.Mi("fs").Mi(b.name), result.Mi("variables"))
	if err != nil {
		return result, err
	}

	// 1. load config from directories and files
	b.logger.Debug("loading config from %v", configs.String())
	fromFile, err := LoadConfigFromFileSystem(configs.Multiple(), maps.New(), b.strategy) // create empty data to not pass template yet
	if err != nil {
		return result, err
	}
	fromFile.ForEach(func(key string, value interface{}) {
		result.Set(key, value)
	})

	// 2. override it with environment
	fromEnv, err := ParseConfigFromEnv(environments)
	if err != nil {
		return result, err
	}
	b.updateResult("env", result, fromEnv)

	// 3. override it will override map
	b.updateResult("argument", result, b.override)

	return result, nil
}
