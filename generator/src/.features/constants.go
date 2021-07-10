package features

const (
	KEY_KEY      = "key"
	KEY_NAME     = "name"
	KEY_TEXT     = "text"
	KEY_CHMOD    = "chmod"
	KEY_USER     = "user"
	KEY_MODE     = "mode"
	KEY_STRATEGY = "strategy"
	KEY_CONFIG   = "config"

	KEY_DISABLE = "disable" // can use with `KEY_MODE`
	KEY_DEBUG   = "debug"   // can use with `KEY_MODE`

	KEY_ADDON = "addon" // can use with `KEY_CONFIG`

	KEY_OUTPUT   = "output"
	KEY_INPUT    = "input"
	KEY_TEMPLATE = "template"
	KEY_DATA     = "data"

	KEY_FORMAT = "format" // can use with `KEY_BASEDIR` | `KEY_DIRNAME` | `KEY_FILENAME` | `KEY_FILENAMES`
)

var noDeps Dependencies = make(Dependencies)
