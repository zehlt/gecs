package resource

import (
	"errors"
	"reflect"
)

var (
	ErrResourceAlreadyExists = errors.New("resource already exists")
	ErrResourceDoesNotExist  = errors.New("resource does not exist")
)

type Locker interface {
	Add(c interface{}) error
	Get(t interface{}) (interface{}, error)
	Has(t interface{}) bool
}

type locker struct {
	resources map[reflect.Type]interface{}
}

func NewLocker() Locker {
	return &locker{
		resources: make(map[reflect.Type]interface{}),
	}
}

func (r *locker) Add(c interface{}) error {
	t := reflect.TypeOf(c)

	_, ok := r.resources[t]
	if ok {
		return ErrResourceAlreadyExists
	}

	r.resources[t] = c
	return nil
}

func (r *locker) Get(c interface{}) (interface{}, error) {
	t := reflect.TypeOf(c)

	data, ok := r.resources[t]
	if !ok {
		return nil, ErrResourceDoesNotExist
	}

	return data, nil
}

func (r *locker) Has(c interface{}) bool {
	t := reflect.TypeOf(c)
	_, ok := r.resources[t]
	return ok
}
