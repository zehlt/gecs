package component

import (
	"errors"

	"github.com/zehlt/gecs/entity"
)

var (
	ErrEntityAlreadyHasComponent  = errors.New("component is already own by the entity")
	ErrEntityDoesNotHaveComponent = errors.New("entity does not have component")
)

type ComponentId int

type Store interface {
	Add(entity.Entity, ComponentId, interface{}) error
	Remove(entity.Entity, ComponentId) error
	RemoveAll(entity.Entity) error
	Get(entity.Entity, ComponentId) (interface{}, error)
	Has(entity.Entity, ComponentId) bool
}

type defaultStore struct {
	// containers []Container
}

func NewStore() Store {
	return &defaultStore{}
}

func (s *defaultStore) Add(e entity.Entity, id ComponentId, c interface{}) error {
	return nil
}

func (s *defaultStore) Remove(e entity.Entity, id ComponentId) error {
	return nil
}

func (s *defaultStore) RemoveAll(e entity.Entity) error {
	return nil
}

func (s *defaultStore) Get(e entity.Entity, id ComponentId) (interface{}, error) {
	return nil, nil
}

func (s *defaultStore) Has(e entity.Entity, id ComponentId) bool {
	return false
}
