package freqtrade

import "github.com/kamontat/fthelper/shared/maps"

type Config struct {
	Cache *CacheConfig
	Query *QueryConfig
}

func newConfig(config maps.Mapper) *Config {
	var cache = &CacheConfig{data: make(maps.Mapper)}
	_ = config.Mi("cache").Struct(&cache.data) // load cache data

	return &Config{
		Cache: cache,
		Query: newQueryConfig(config.Mi("query")),
	}
}
