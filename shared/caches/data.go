package caches

import (
	"errors"
	"fmt"
	"time"
)

type Data struct {
	Key     string
	Data    interface{}
	Error   error
	updater Updater

	createAt time.Time
	updateAt time.Time
	// time when it expired < 0 mean never expire
	expireAt time.Duration
}

func (d *Data) CreateAt() time.Time {
	return d.createAt
}
func (d *Data) UpdateAt() time.Time {
	return d.updateAt
}
func (d *Data) ExpireAt() *time.Time {
	if d.expireAt <= 0 {
		return nil
	} else {
		var expireAt = d.updateAt.Add(d.expireAt)
		return &expireAt
	}
}

func (d *Data) IsExist() bool {
	return d.Error == nil && d.Data != nil
}
func (d *Data) IsExpired() bool {
	if d.expireAt <= 0 {
		return false
	}
	var duration = time.Since(d.updateAt)
	return duration > d.expireAt
}

// Extend will update updateAt to now()
func (d *Data) Extend() *Data {
	d.updateAt = time.Now()
	return d
}

// Update will force call updater and update updateAt value
func (d *Data) Update() error {
	d.Data, d.Error = d.updater(d.Data)
	d.Extend()
	return d.Error
}

// UpdateData same with Update but return data out
func (d *Data) UpdateData() (interface{}, error) {
	var err = d.Update()
	return d.Data, err
}

// Fetch will update only if data is missing or expired
// return true if data got updated
func (d *Data) Fetch() (bool, error) {
	var needFetch = !d.IsExist() || d.IsExpired()
	if needFetch {
		_ = d.Update()
	}

	return needFetch, d.Error
}

// FetchData same with Fetch but return data out
func (d *Data) FetchData() (interface{}, error) {
	var _, err = d.Fetch()
	return d.Data, err
}

func (d *Data) String() string {
	var expireAt = d.ExpireAt()

	var cteTS = d.CreateAt().Unix()
	var updTS = d.UpdateAt().Unix()
	var expTS = int64(0)
	if expireAt != nil {
		expTS = expireAt.Unix()
	}

	return fmt.Sprintf(
		"%s: %v (%v) | C: %d | U: %d | E: %d",
		d.Key,
		d.Data,
		d.Error,
		cteTS,
		updTS,
		expTS,
	)
}

func NewData(key string, updater Updater, expireAt time.Duration) *Data {
	if expireAt < 0 {
		expireAt = -1 // force negative expire to be only -1
	}

	return &Data{
		Key:      key,
		Data:     nil,
		Error:    errors.New("empty data"),
		updater:  updater,
		expireAt: expireAt,
		createAt: time.Now(),
		updateAt: time.Now(),
	}
}

// SData is constants data
func SData(key string, data interface{}) *Data {
	var d = NewData(key, func(o interface{}) (interface{}, error) {
		return data, nil
	}, -1)

	_, _ = d.Fetch()
	return d
}
