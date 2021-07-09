package caches

type Updater func(o interface{}) (interface{}, error)
type Creator func() (interface{}, error)
