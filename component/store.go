package component

import (
	"errors"

	"github.com/zehlt/gecs/entity"
)

var (
	ErrEntityAlreadyHasComponent  = errors.New("component is already own by the entity")
	ErrEntityDoesNotHaveComponent = errors.New("entity does not have component")
	ErrContainerTypeDoesNotExist  = errors.New("container type does not exist")
)

type ComponentId int

type Store interface {
	Register(ComponentId, ContainerType) error
	Add(entity.Entity, ComponentId, interface{}) error
	Remove(entity.Entity, ComponentId) error
	RemoveAll(entity.Entity) error
	Get(entity.Entity, ComponentId) (interface{}, error)
	GetAll(entity.Entity) []interface{}
	Has(entity.Entity, ComponentId) bool
}

type defaultStore struct {
	containers map[ComponentId]Container
}

func NewStore() Store {
	return &defaultStore{
		containers: make(map[ComponentId]Container),
	}
}

func (s *defaultStore) Register(id ComponentId, t ContainerType) error {
	_, ok := s.containers[id]
	if ok {
		return nil
	}

	switch t {
	case SPARSE_ARRAY_CONTAINER:
		s.containers[id] = newSparseArray()
		return nil

	case NULL_CONTAINER:
		s.containers[id] = newNull()
		return nil

	case HASHMAP_CONTAINER:
		s.containers[id] = newHashmap()
		return nil

	default:
		return ErrContainerTypeDoesNotExist
	}
}

func (s *defaultStore) Add(e entity.Entity, id ComponentId, c interface{}) error {
	return s.containers[id].Add(e, c)
}

func (s *defaultStore) Remove(e entity.Entity, id ComponentId) error {
	return s.containers[id].Remove(e)
}

func (s *defaultStore) RemoveAll(e entity.Entity) error {
	for _, container := range s.containers {
		err := container.Remove(e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *defaultStore) Get(e entity.Entity, id ComponentId) (interface{}, error) {
	return s.containers[id].Get(e)
}

func (s *defaultStore) GetAll(e entity.Entity) []interface{} {
	components := make([]interface{}, 0)

	for _, container := range s.containers {
		if container.Has(e) {
			component, _ := container.Get(e)
			components = append(components, component)
		}
	}

	return components
}

func (s *defaultStore) Has(e entity.Entity, id ComponentId) bool {
	return s.containers[id].Has(e)
}
