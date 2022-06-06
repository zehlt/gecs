package gecs

import (
	"github.com/zehlt/gecs/component"
	"github.com/zehlt/gecs/entity"
	"github.com/zehlt/gecs/registry"
)

type World interface {
	CreateEntity() (entity.Entity, error)
	DestroyEntity(entity.Entity) error
	EntityExists(entity.Entity) bool

	RegisterComponent(interface{}, component.ContainerType) error
	AddComponent(entity.Entity, interface{}) error
	RemoveComponent(entity.Entity, interface{}) error
	GetComponent(entity.Entity, interface{}) (interface{}, error)
	HasComponent(entity.Entity, interface{}) bool
}

type world struct {
	arena    entity.Arena
	store    component.Store
	registry registry.Registry
}

func DefaultWorld() World {
	return &world{
		arena:    entity.NewArena(),
		store:    component.NewStore(),
		registry: registry.NewRegistry(),
	}
}

func (w *world) CreateEntity() (entity.Entity, error) {
	return w.arena.Create()
}

func (w *world) DestroyEntity(e entity.Entity) error {
	err := w.arena.Destroy(e)
	if err != nil {
		// TODO: layer more error
		return err
	}

	w.registry.DestroySignature(e)

	return w.store.RemoveAll(e)
}

func (w *world) EntityExists(e entity.Entity) bool {
	return w.arena.Exists(e)
}

func (w *world) RegisterComponent(c interface{}, t component.ContainerType) error {
	id := w.registry.GetComponentId(c)
	return w.store.Register(id, t)
}

func (w *world) AddComponent(e entity.Entity, c interface{}) error {
	if !w.arena.Exists(e) {
		return entity.ErrEntityDoesNotExist
	}

	componenId := w.registry.GetComponentId(c)
	err := w.registry.AddComponent(e, componenId)
	if err != nil {
		// TODO: add layer of info in error
		return err
	}

	return w.store.Add(e, componenId, c)
}

func (w *world) RemoveComponent(e entity.Entity, c interface{}) error {
	if !w.arena.Exists(e) {
		return entity.ErrEntityDoesNotExist
	}

	id := w.registry.GetComponentId(c)
	if !w.registry.HasComponent(e, id) {
		return component.ErrEntityDoesNotHaveComponent
	}
	w.registry.RemoveComponent(e, id)

	return w.store.Remove(e, id)
}

func (w *world) GetComponent(e entity.Entity, c interface{}) (interface{}, error) {
	if !w.arena.Exists(e) {
		return nil, entity.ErrEntityDoesNotExist
	}

	id := w.registry.GetComponentId(c)
	if !w.registry.HasComponent(e, id) {
		return nil, component.ErrEntityDoesNotHaveComponent
	}

	return w.store.Get(e, id)
}

func (w *world) HasComponent(e entity.Entity, c interface{}) bool {
	if !w.arena.Exists(e) {
		return false
	}

	id := w.registry.GetComponentId(c)
	return w.registry.HasComponent(e, id)
}
