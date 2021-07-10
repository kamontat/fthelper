package hooks

type Type string

const (
	// This is first hook that will run before every thing (#1).
	// The config on this stage will always be empty
	// You cannot use this in plugin of course before it execute before plugin
	// You might use this hook to add some information before cli flow will be start
	BEFORE_PLUGIN = "before-plugin"

	// run after all plugins are loaded

	// This will execute after all plugins are loaded (#2).
	// The config on this stage will contains plugins setup data
	// You might use this to setup config that relate to multiple plugins
	AFTER_PLUGIN = "after-plugin"

	// This will execute after all flags are loaded (#3).
	// The config on this stage will contains mostly every information setup by client
	// You might use this to update data to config data that relate to flags
	AFTER_FLAG = "after-flag"

	// This will execute before the command will be execute (#4).
	// The config on this stage will contains extra data added by cli itself.
	//   `internal.command` - current execution command
	//   `internal.args`    - cli arguments that leftover to command itself
	// You might use this to final setup before command will be run
	BEFORE_COMMAND = "before-command"

	// This will execute after the command is executed (#5).
	// The config on this stage will contains extra more information added by cli itself.
	//   `internal.error` - command execution result (error)
	// You might use this to override the error or print summary result
	AFTER_COMMAND = "after-command"
)
