package caches

import "sync"

func New() *Service {
	return &Service{
		caches: make(map[string]*Data),
		mutex:  sync.RWMutex{},
	}
}

var Global = New()
