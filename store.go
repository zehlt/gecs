package gecs

import (
	"errors"
)

var (
	ErrEntityAlreadyHasComponent  = errors.New("component is already own by the entity")
	ErrEntityDoesNotHaveComponent = errors.New("entity does not have component")
	ErrContainerTypeDoesNotExist  = errors.New("container type does not exist")
)

type ComponentId int

type store interface {
	Register(ComponentId, ContainerType) error
	Add(Entity, ComponentId, interface{}) error
	Emplace(Entity, ComponentId, interface{})
	Remove(Entity, ComponentId) error
	RemoveAll(Entity) error
	Get(Entity, ComponentId) (interface{}, error)
	GetAll(Entity) []interface{}
	Has(Entity, ComponentId) bool
}

type defaultStore struct {
	containers map[ComponentId]container
}

func newStore() store {
	return &defaultStore{
		containers: make(map[ComponentId]container),
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

func (s *defaultStore) Add(e Entity, id ComponentId, c interface{}) error {
	return s.containers[id].Add(e, c)
}

func (s *defaultStore) Emplace(e Entity, id ComponentId, c interface{}) {
	s.containers[id].Emplace(e, c)
}

func (s *defaultStore) Remove(e Entity, id ComponentId) error {
	return s.containers[id].Remove(e)
}

func (s *defaultStore) RemoveAll(e Entity) error {
	for _, container := range s.containers {
		err := container.Remove(e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *defaultStore) Get(e Entity, id ComponentId) (interface{}, error) {
	return s.containers[id].Get(e)
}

func (s *defaultStore) GetAll(e Entity) []interface{} {
	components := make([]interface{}, 0)

	for _, container := range s.containers {
		if container.Has(e) {
			component, _ := container.Get(e)
			components = append(components, component)
		}
	}

	return components
}

func (s *defaultStore) Has(e Entity, id ComponentId) bool {
	return s.containers[id].Has(e)
}
