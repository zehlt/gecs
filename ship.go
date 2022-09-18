package gecs

import (
	"errors"
)

var (
	ErrShipContainerAlreadyRegistered = errors.New("container already registered for component")
	ErrShipContainerNotRegistered     = errors.New("no container registered for component")
	ErrShipContainerTypeNotValid      = errors.New("container type not valid")
)

type ComponentType int

type Component interface {
	GetType() ComponentType
}

type ship struct {
	containers map[ComponentType]container
}

func newShip() *ship {
	return &ship{
		containers: make(map[ComponentType]container),
	}
}

func (s *ship) RegisterComponent(c ComponentType, t ContainerType) error {
	_, ok := s.containers[c]
	if ok {
		return ErrShipContainerAlreadyRegistered
	}

	switch t {
	case HASHMAP_CONTAINER:
		s.containers[c] = newMapContainer()
	case TAG_CONTAINER:
		s.containers[c] = newTagContainer()
	default:
		return ErrShipContainerTypeNotValid
	}

	return nil
}

func (s *ship) EmplaceComponent(e Entity, c Component) {
	container := s.containers[c.GetType()]
	container.Emplace(e, c)
}

func (s *ship) GetComponent(e Entity, t ComponentType) Component {
	container := s.containers[t]

	return container.Get(e)
}

func (s *ship) RemoveComponent(e Entity, t ComponentType) {
	container, ok := s.containers[t]
	if ok {
		container.Remove(e)
	}
}

func (s *ship) RemoveAllComponents(e Entity) {
	for _, container := range s.containers {
		container.Remove(e)
	}
}

var (
	ErrContainerEntityAlreadyHasComponent  = errors.New("entity already has component")
	ErrContainerEntityDoesNotHaveComponent = errors.New("entity does not have component")
)

type ContainerType int

const (
	HASHMAP_CONTAINER ContainerType = iota
	TAG_CONTAINER
)

type container interface {
	Add(e Entity, c Component)
	Emplace(e Entity, c Component)
	Get(e Entity) Component
	Remove(e Entity)
}

type mapContainer struct {
	hash map[Entity]Component
}

func newMapContainer() container {
	return &mapContainer{
		hash: make(map[Entity]Component),
	}
}

func (ctn *mapContainer) Add(e Entity, c Component) {
	ctn.hash[e] = c
}

func (ctn *mapContainer) Emplace(e Entity, c Component) {
	ctn.hash[e] = c
}

func (ctn *mapContainer) Remove(e Entity) {
	delete(ctn.hash, e)
}

func (ctn *mapContainer) Get(e Entity) Component {
	c := ctn.hash[e]
	return c
}

type tagContainer struct {
}

func newTagContainer() container {
	return &tagContainer{}
}

func (t *tagContainer) Add(e Entity, c Component) {
}

func (t *tagContainer) Emplace(e Entity, c Component) {

}

func (t *tagContainer) Get(e Entity) Component {
	return nil
}

func (t *tagContainer) Remove(e Entity) {
}
