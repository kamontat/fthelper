package configs

import (
	"strings"

	"github.com/kamontat/fthelper/shared/maps"
)

func ParseOverride(str string) (key, value string, ok bool) {
	ok = false
	var arr = strings.Split(str, "=")
	if len(arr) != 2 {
		return
	}

	ok = true
	key = arr[0]
	value = arr[1]
	return
}

func BuildClusterConfig(cluster string, config maps.Mapper) maps.Mapper {
	return maps.Merger(config).Add(config.Mi("_").Mi(cluster)).Merge()
}
