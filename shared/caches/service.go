package caches

import (
	"fmt"
	"sync"

	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type Service struct {
	caches map[string]*Data
	mutex  sync.RWMutex

	logger *loggers.Logger
}

func (s *Service) Size() int {
	return len(s.caches)
}

func (s *Service) Has(key string) bool {
	s.mutex.RLock()
	d, ok := s.caches[key]
	s.mutex.RUnlock()

	return ok && d.IsExist()
}

func (s *Service) Get(key string) *Data {
	s.mutex.RLock()
	d, ok := s.caches[key]
	s.mutex.RUnlock()
	if ok {
		return d
	}

	s.logger.Debug("cannot get data key: %s", key)
	return SData(key, nil)
}

func (s *Service) SetData(data *Data) error {
	if s.Has(data.Key) {
		return fmt.Errorf("cannot set data with existing data, use Update() instead")
	} else if data.Key == "" {
		return fmt.Errorf("cache data key cannot be empty string")
	}

	err := data.Update()
	if err != nil {
		return err
	}

	s.logger.Debug("creating '%s' data", data.Key)
	s.mutex.Lock()
	s.caches[data.Key] = data
	s.mutex.Unlock()
	return nil
}

func (s *Service) SetFn(key string, creator Creator, expireAt string) error {
	var data = NewData(key, func(o interface{}) (interface{}, error) {
		return creator()
	}, parseDuration(expireAt))

	return s.SetData(data)
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

	s.logger.Debug("updating '%s' data", key)
	var data = s.Get(key)
	data.updater = updater
	return data.Update()
}

func (s *Service) Update(key string, value interface{}, expireAt string) {
	_ = s.UpdateFn(key, func(o interface{}) (interface{}, error) {
		return value, nil
	}, expireAt)
}

// Increase is shorted method for increase number everytime it's called by n
func (s *Service) IncreaseN(key string, increase int) {
	_ = s.UpdateFn(key, func(o interface{}) (interface{}, error) {
		if o == nil {
			return increase, nil
		}
		return o.(int) + increase, nil
	}, Persistent)
}

// Increase is shorted method for increase number everytime it's called by 1
func (s *Service) Increase(key string) {
	s.IncreaseN(key, 1)
}

// Bucket will keep value as []interface{}
func (s *Service) Bucket(key string, value interface{}, limit int, expireAt string) error {
	return s.UpdateFn(key, func(o interface{}) (interface{}, error) {
		if o == nil {
			return []interface{}{value}, nil
		}

		// keep only last n number
		var queue = o.([]interface{})
		if len(queue) > limit {
			queue[0] = nil    // assign to zero value to free memory
			queue = queue[1:] // Dequeue
		}

		queue = append(queue, value) // Enqueue
		return queue, nil
	}, expireAt)
}

func (s *Service) Fetch(key string, updater Updater, expireAt string) error {
	if !s.Has(key) {
		return s.SetFn(key, func() (interface{}, error) {
			return updater(nil)
		}, expireAt)
	}

	fetch, err := s.Get(key).Fetch()
	if fetch {
		s.logger.Debug("fetching '%s' data", key)
	}

	return err
}

func (s *Service) FetchAll() error {
	var errs = errors.New()
	s.mutex.RLock()
	for _, data := range s.caches {
		_, err := data.Fetch()
		errs.And(err)
	}
	s.mutex.RUnlock()

	return errs.Error()
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
