package hooks

type Type string

const (
	// run before execute plugins, This will not works on plugin
	BEFORE_PLUGIN = "before-plugin"
	// run after all plugins are loaded
	AFTER_PLUGIN = "after-plugin"
	// run after all flags are loaded
	AFTER_FLAG = "after-flag"
	// run before final execute command
	BEFORE_COMMAND = "before-command"
)
