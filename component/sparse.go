package component

import (
	"reflect"

	"github.com/zehlt/gecs/entity"
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

	comps, ok := s.store[t]
	if !ok {
		s.store[t] = make([]interface{}, 0)
		comps = s.store[t]
	}

	if e.Id >= len(comps) {
		for i := e.Id - len(comps); i >= 0; i-- {
			comps = append(comps, nil)
		}
	}

	comps[e.Id] = c

	return nil
}

func (s *sparseStore) Remove(e entity.Entity, c interface{}) error {
	return nil
}

func (s *sparseStore) Get(e entity.Entity, c interface{}) (interface{}, error) {
	return nil, nil
}

func (s *sparseStore) Has(e entity.Entity, c interface{}) bool {
	t := reflect.TypeOf(c)

	comps, ok := s.store[t]
	if !ok {
		return false
	}

	if e.Id > len(comps)-1 {
		return false
	}

	if comps[e.Id] == nil {
		return false
	}

	return true
}
