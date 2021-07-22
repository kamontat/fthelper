package caches

import (
	"errors"
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

func (d *Data) Update() (interface{}, error) {
	d.Data, d.Error = d.updater(d.Data)
	d.Extend()
	return d.Data, d.Error
}

// Fetch will update only if data is missing or expired
func (d *Data) Fetch() (interface{}, error) {
	if !d.IsExist() || d.IsExpired() {
		return d.Update()
	}
	return d.Data, d.Error
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

	d.Fetch()
	return d
}
