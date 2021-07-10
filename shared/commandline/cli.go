package commandline

import (
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/commandline/hooks"
	"github.com/kamontat/fthelper/shared/commandline/models"
	"github.com/kamontat/fthelper/shared/commandline/plugins"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type cli struct {
	Metadata *models.Metadata
	flags    *flags.Manager // global flags
	commands *commands.Manager
	hooks    *hooks.Manager
	plugins  *plugins.Manager

	cache  *caches.Service
	logger *loggers.Logger
}

func (c *cli) Flag(flag flags.Flag) *cli {
	c.flags.Add(flag)
	return c
}

func (c *cli) Command(cmd *commands.Command) *cli {
	c.commands.Add(cmd)
	return c
}

func (c *cli) Hook(t hooks.Type, hook hooks.Hook) *cli {
	c.hooks.Add(t, hook)
	return c
}

func (c *cli) Plugin(plugin plugins.Plugin) *cli {
	c.plugins.Add(plugin)
	return c
}

func (c *cli) Start(args []string) error {
	var config = maps.New()
	if err := c.hooks.Start(hooks.BEFORE_PLUGIN, config); err != nil {
		return err
	}

	var err = c.plugins.Build(&plugins.PluginParameter{
		Metadata:   *c.Metadata,
		NewCommand: c.commands.Add,
		NewFlags:   c.flags.Add,
		NewHook:    c.hooks.Add,
		Config:     config,
		Logger:     c.logger,
	})
	if err == nil {
		err = c.hooks.Start(hooks.AFTER_PLUGIN, config)
	}
	if err != nil {
		return err
	}

	option, parsed, err := c.flags.Build(args[0], args[1:])
	if err != nil {
		return err
	}
	for _, value := range option {
		config = maps.Merger(config).Add(value).Merge()
	}
	if err = c.hooks.Start(hooks.AFTER_FLAG, config); err != nil {
		return err
	}

	var cmd, final = c.commands.Get(parsed, config)
	if len(final) > 0 {
		config.Set("internal.command", cmd.Name).Set("internal.args", final)
	}

	if err = c.hooks.Start(hooks.BEFORE_COMMAND, config); err != nil {
		return err
	}

	var cmderr = cmd.Start(&commands.ExecutorParameter{
		Name:   cmd.Name,
		Meta:   c.Metadata,
		Args:   final,
		Config: config,

		Cache:  c.cache,
		Logger: loggers.Get("command", cmd.Name),
	})

	config.Set("internal.error", cmderr)
	if err = c.hooks.Start(hooks.AFTER_COMMAND, config); err != nil {
		return err
	}

	return cmderr
}
