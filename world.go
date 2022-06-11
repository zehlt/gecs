package gecs

import (
	"github.com/zehlt/gecs/component"
	"github.com/zehlt/gecs/entity"
	"github.com/zehlt/gecs/resource"

	"github.com/zehlt/gecs/signature"
)

type World interface {
	CreateEntity() (entity.Entity, error)
	DestroyEntity(entity.Entity) error
	EntityExists(entity.Entity) bool
	GetAllEntities() []entity.Entity

	RegisterComponent(interface{}, component.ContainerType) error
	AddComponent(entity.Entity, interface{}) error
	RemoveComponent(entity.Entity, interface{}) error
	GetComponent(entity.Entity, interface{}) (interface{}, error)
	GetComponentById(entity.Entity, component.ComponentId) (interface{}, error)
	GetAllComponentsFromEntity(entity.Entity) ([]interface{}, error)
	GetComponentId(c interface{}) component.ComponentId
	HasComponent(entity.Entity, interface{}) bool

	AddResource(interface{}) error
	GetResource(interface{}) (interface{}, error)
	HasResource(interface{}) bool

	GetSignatureFromTypes(types []interface{}) signature.Signature
	FindMatchingEntities(signature.Signature) []entity.Entity
	GetEntitySignature(e entity.Entity) (signature.Signature, error)
}

type world struct {
	locker   resource.Locker
	arena    entity.Arena
	store    component.Store
	registry signature.Registry
}

func DefaultWorld() World {
	return &world{
		locker:   resource.NewLocker(),
		arena:    entity.NewArena(),
		store:    component.NewStore(),
		registry: signature.NewRegistry(),
	}
}

func (w *world) CreateEntity() (entity.Entity, error) {
	e, err := w.arena.Create()
	if err != nil {
		return entity.Entity{}, err
	}

	err = w.registry.CreateEntitySignature(e)
	if err != nil {
		return entity.Entity{}, err
	}

	return e, nil
}

func (w *world) GetAllEntities() []entity.Entity {
	return w.arena.GetAll()
}

func (w *world) DestroyEntity(e entity.Entity) error {
	err := w.arena.Destroy(e)
	if err != nil {
		// TODO: layer more error
		return err
	}

	w.registry.DestroyEntitySignature(e)

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

func (w *world) GetAllComponentsFromEntity(e entity.Entity) ([]interface{}, error) {
	if !w.arena.Exists(e) {
		return nil, entity.ErrEntityDoesNotExist
	}

	return w.store.GetAll(e), nil
}

func (w *world) GetComponentById(e entity.Entity, id component.ComponentId) (interface{}, error) {
	if !w.arena.Exists(e) {
		return nil, entity.ErrEntityDoesNotExist
	}

	if !w.registry.HasComponent(e, id) {
		return nil, component.ErrEntityDoesNotHaveComponent
	}

	return w.store.Get(e, id)
}

func (w *world) GetComponentId(c interface{}) component.ComponentId {
	return w.registry.GetComponentId(c)
}

func (w *world) HasComponent(e entity.Entity, c interface{}) bool {
	if !w.arena.Exists(e) {
		return false
	}

	id := w.registry.GetComponentId(c)
	return w.registry.HasComponent(e, id)
}

func (w *world) AddResource(c interface{}) error {
	return w.locker.Add(c)
}

func (w *world) GetResource(t interface{}) (interface{}, error) {
	return w.locker.Get(t)
}

func (w *world) HasResource(t interface{}) bool {
	return w.locker.Has(t)
}

func (w *world) GetSignatureFromTypes(types []interface{}) signature.Signature {
	return w.registry.GetSignatureFromTypes(types)
}

func (w *world) FindMatchingEntities(s signature.Signature) []entity.Entity {
	return w.registry.FindMatchingEntities(s)
}

func (w *world) GetEntitySignature(e entity.Entity) (signature.Signature, error) {
	return w.registry.GetEntitySignature(e)
}
