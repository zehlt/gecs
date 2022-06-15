package component

import (
	"github.com/zehlt/gecs/entity"
)

type sparseArray struct {
	components []interface{}
}

func newSparseArray() Container {
	return &sparseArray{
		components: make([]interface{}, 0),
	}
}

func (s *sparseArray) Add(e entity.Entity, c interface{}) error {
	if e.Id() >= len(s.components) {
		for i := e.Id() - len(s.components); i >= 0; i-- {
			s.components = append(s.components, nil)
		}
	}

	if s.components[e.Id()] != nil {
		return ErrEntityAlreadyHasComponent
	}

	s.components[e.Id()] = c

	return nil
}

func (s *sparseArray) Emplace(e entity.Entity, c interface{}) {
	if e.Id() >= len(s.components) {
		for i := e.Id() - len(s.components); i >= 0; i-- {
			s.components = append(s.components, nil)
		}
	}

	s.components[e.Id()] = c
}

func (s *sparseArray) Remove(e entity.Entity) error {
	if !s.Has(e) {
		return ErrEntityDoesNotHaveComponent
	}

	s.components[e.Id()] = nil

	return nil
}

func (s *sparseArray) Get(e entity.Entity) (interface{}, error) {
	if !s.Has(e) {
		return nil, ErrEntityDoesNotHaveComponent
	}

	return s.components[e.Id()], nil
}

func (s *sparseArray) Has(e entity.Entity) bool {
	if e.Id() >= len(s.components) {
		return false
	}

	if s.components[e.Id()] == nil {
		return false
	}

	return true
}
