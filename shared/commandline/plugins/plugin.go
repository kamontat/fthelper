package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/commandline/hooks"
	"github.com/kamontat/fthelper/shared/commandline/models"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type PluginParameter struct {
	Metadata models.Metadata

	NewCommand commands.Creator
	NewFlags   flags.Creator
	NewHook    hooks.Creator

	Config maps.Mapper
	Logger *loggers.Logger
}

type Plugin func(p *PluginParameter) error
