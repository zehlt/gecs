package world

import (
	"github.com/zehlt/gecs/entity"
)

type World interface {
	CreateEntity() (entity.Entity, error)
	DestroyEntity(entity.Entity) error
	EntityExists(entity.Entity) bool
	AddComponent(entity.Entity, interface{}) error
	RemoveComponent(entity.Entity, interface{}) error
	GetComponent(entity.Entity, interface{}) (interface{}, error)
	HasComponent(entity.Entity, interface{}) bool
}

// type SparseRegistry struct {
// 	arena entity.Arena
// 	store component.Store
// }

// func NewSparceRegistry() SparseRegistry {
// 	return SparseRegistry{
// 		arena: entity.NewArena(),
// 		store: component.NewSparseStore(),
// 	}
// }

// func (r *SparseRegistry) CreateEntity() (entity.Entity, error) {
// 	return r.arena.Create()
// }

// func (r *SparseRegistry) DestroyEntity(e entity.Entity) error {
// 	return r.arena.Destroy(e)
// }

// func (r *SparseRegistry) EntityExists(e entity.Entity) bool {
// 	return r.arena.Exists(e)
// }

// func (r *SparseRegistry) AddComponent(e entity.Entity, c interface{}) error {
// 	if !r.arena.Exists(e) {
// 		return entity.ErrEntityDoesNotExist
// 	}

// 	return r.store.Add(e, c)
// }

// func (r *SparseRegistry) RemoveComponent(e entity.Entity, c interface{}) error {
// 	if !r.arena.Exists(e) {
// 		return entity.ErrEntityDoesNotExist
// 	}

// 	return r.store.Remove(e, c)
// }

// func (r *SparseRegistry) GetComponent(e entity.Entity, c interface{}) (interface{}, error) {
// 	if !r.arena.Exists(e) {
// 		return nil, entity.ErrEntityDoesNotExist
// 	}

// 	return r.store.Get(e, c)
// }

// func (r *SparseRegistry) HasComponent(e entity.Entity, c interface{}) bool {
// 	if !r.arena.Exists(e) {
// 		return false
// 	}

// 	return r.store.Has(e, c)
// }
