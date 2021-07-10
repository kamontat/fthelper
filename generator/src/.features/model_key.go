package features

import (
	"github.com/kamontat/fthelper/shared/configs"
)

func Env(env configs.Env) Feature {
	return Key(env, configs.KEY_DEFAULT, configs.SUBKEY_DEFAULT)
}

func Key(env configs.Env, key configs.Key, subkey configs.SubKey) Feature {
	return KeyModel(configs.NewKeyModel(env, key, subkey))
}

func KeyModel(key *configs.KeyModel) Feature {
	return Raw(KEY_KEY, noDeps, withStaticExecutor(key))
}
