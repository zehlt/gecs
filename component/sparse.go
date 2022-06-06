package component

import (
	"errors"
	"reflect"

	"github.com/zehlt/gecs/entity"
)

var (
	ErrComponentAlreadyOwnByEntity = errors.New("component is already own by the entity")
	ErrEntityDoesNotHaveComponent  = errors.New("entity does not have component")
)

type sparseStore struct {
	store map[reflect.Type][]interface{}
}

func NewSparseStore() Store {
	return &sparseStore{
		store: make(map[reflect.Type][]interface{}),
	}
}

func (s *sparseStore) Add(e entity.Entity, c interface{}) error {
	t := reflect.TypeOf(c)

	_, ok := s.store[t]
	if !ok {
		s.store[t] = make([]interface{}, 0)
	}

	if e.Id >= len(s.store[t]) {
		for i := e.Id - len(s.store[t]); i >= 0; i-- {
			s.store[t] = append(s.store[t], nil)
		}
	}

	if s.store[t][e.Id] != nil {
		return ErrComponentAlreadyOwnByEntity
	}

	s.store[t][e.Id] = c

	return nil
}

func (s *sparseStore) Remove(e entity.Entity, c interface{}) error {
	if !s.Has(e, c) {
		return ErrEntityDoesNotHaveComponent
	}

	t := reflect.TypeOf(c)
	s.store[t][e.Id] = nil
	return nil
}

func (s *sparseStore) Get(e entity.Entity, c interface{}) (interface{}, error) {
	if !s.Has(e, c) {
		return nil, ErrEntityDoesNotHaveComponent
	}

	t := reflect.TypeOf(c)

	return s.store[t][e.Id], nil
}

func (s *sparseStore) Has(e entity.Entity, c interface{}) bool {
	t := reflect.TypeOf(c)

	_, ok := s.store[t]
	if !ok {
		return false
	}

	if e.Id > len(s.store[t])-1 {
		return false
	}

	if s.store[t][e.Id] == nil {
		return false
	}

	return true
}
