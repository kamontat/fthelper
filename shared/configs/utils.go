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
	if cluster == "" {
		return config
	}

	// We check both cluster directly and cluster with lower case
	// for user who setup from environment variable, which cannot setup with upper case
	var c, _ = config.Mi("_").Gets(cluster, strings.ToLower(cluster))
	var clusterConfig, _ = maps.ToMapper(c)

	return maps.Merger(config).Add(clusterConfig).Merge()
}
