package caches

import (
	"fmt"
	"sync"

	"github.com/kamontat/fthelper/shared/maps"
)

type Service struct {
	caches map[string]*Data
	mutex  sync.RWMutex
}

func (s *Service) Size() int {
	return len(s.caches)
}

func (s *Service) Has(key string) bool {
	s.mutex.Lock()
	d, ok := s.caches[key]
	s.mutex.Unlock()

	return ok && d.IsExist()
}

func (s *Service) Get(key string) *Data {
	s.mutex.RLock()
	d, ok := s.caches[key]
	s.mutex.RUnlock()
	if ok {
		return d
	}

	return SData(key, nil)
}

func (s *Service) SetFn(key string, creator Creator, expireAt string) error {
	if s.Has(key) {
		return fmt.Errorf("cannot set data with existing data, use Update() instead")
	}

	var data = NewData(key, func(o interface{}) (interface{}, error) {
		return creator()
	}, parseDuration(expireAt))

	_, err := data.Update()
	if err != nil {
		return err
	}

	s.mutex.Lock()
	s.caches[key] = data
	s.mutex.Unlock()
	return nil
}

func (s *Service) Set(key string, value interface{}, expireAt string) error {
	return s.SetFn(key, func() (interface{}, error) {
		return value, nil
	}, expireAt)
}

func (s *Service) UpdateFn(key string, updater Updater, expireAt string) error {
	if !s.Has(key) {
		return s.SetFn(key, func() (interface{}, error) {
			return updater(nil)
		}, expireAt)
	}

	var data = s.Get(key)
	data.updater = updater
	_, err := data.Update()
	return err
}

func (s *Service) Update(key string, value interface{}, expireAt string) error {
	return s.UpdateFn(key, func(o interface{}) (interface{}, error) {
		return value, nil
	}, expireAt)
}

// Increase is shorted method for increase number everytime it's called by n
func (s *Service) IncreaseN(key string, increase int) error {
	return s.UpdateFn(key, func(o interface{}) (interface{}, error) {
		if o == nil {
			return increase, nil
		}
		return o.(int) + increase, nil
	}, Persistent)
}

// Increase is shorted method for increase number everytime it's called by 1
func (s *Service) Increase(key string) error {
	return s.IncreaseN(key, 1)
}

func (s *Service) Fetch(key string, updater Updater, expireAt string) error {
	if !s.Has(key) {
		return s.SetFn(key, func() (interface{}, error) {
			return updater(nil)
		}, expireAt)
	}

	_, err := s.Get(key).Fetch()
	return err
}

func (s *Service) String() string {
	var m = maps.New()
	for key, value := range s.caches {
		m[key] = value
	}

	var json, err = maps.ToJson(m)
	if err != nil {
		return err.Error()
	}
	return string(json)
}
