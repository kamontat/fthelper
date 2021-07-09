package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/hooks"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

// Support is example how we can implement plugins support
func SupportBanner(p *PluginParameter) error {
	p.NewHook(hooks.BEFORE_COMMAND, func(m maps.Mapper) error {
		if loggers.IsDebug() {
			p.Logger.Debug("%-12s: %s", "metadata", p.Metadata.String())
			p.Logger.Debug("%-12s: %v", "config", m.String())

			return nil
		} else {
			table := p.Logger.Table(3)

			p.Logger.Newline()
			p.Logger.Line()
			table.Header("Name", "Version", "Commit")
			table.Row(p.Metadata.Name, p.Metadata.Version, p.Metadata.Commit)
			var err = table.End()
			p.Logger.Line()
			p.Logger.Newline()

			return err
		}
	})
	return nil
}
