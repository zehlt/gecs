package component

import (
	"github.com/zehlt/gecs/entity"
)

type SparseArray struct {
	components []interface{}
}

func (s *SparseArray) Add(e entity.Entity, c interface{}) error {
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

func (s *SparseArray) Remove(e entity.Entity) error {
	if !s.Has(e) {
		return ErrEntityDoesNotHaveComponent
	}

	s.components[e.Id()] = nil

	return nil
}

func (s *SparseArray) Get(e entity.Entity) (interface{}, error) {
	if !s.Has(e) {
		return nil, ErrEntityDoesNotHaveComponent
	}

	return s.components[e.Id()], nil
}

func (s *SparseArray) Has(e entity.Entity) bool {
	if e.Id() >= len(s.components) {
		return false
	}

	if s.components[e.Id()] == nil {
		return false
	}

	return true
}
