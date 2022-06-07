package signature

import (
	"errors"
	"reflect"

	"github.com/zehlt/gecs/component"
	"github.com/zehlt/gecs/entity"
)

var (
	ErrEntityAlreadyHasSignature  = errors.New("entity already has a signature")
	ErrEntityDoesNotHaveSignature = errors.New("entity does not have a signature")
)

type Registry interface {
	CreateEntitySignature(entity.Entity) error
	DestroyEntitySignature(entity.Entity) error

	AddComponent(entity.Entity, component.ComponentId) error
	RemoveComponent(entity.Entity, component.ComponentId) error
	HasComponent(entity.Entity, component.ComponentId) bool
	GetComponentId(c interface{}) component.ComponentId

	GetSignatureFromTypes([]interface{}) Signature
	FindMatchingEntities(Signature) []entity.Entity
}

type defaultRegistry struct {
	signatures map[entity.Entity]Signature

	types   map[reflect.Type]component.ComponentId
	next_id component.ComponentId
}

func NewRegistry() Registry {
	return &defaultRegistry{
		signatures: make(map[entity.Entity]Signature),
		types:      make(map[reflect.Type]component.ComponentId),
		next_id:    0,
	}
}

func (r *defaultRegistry) CreateEntitySignature(e entity.Entity) error {
	_, ok := r.signatures[e]
	if ok {
		return ErrEntityAlreadyHasSignature
	}

	r.signatures[e] = NewSignature()

	return nil
}

func (r *defaultRegistry) DestroyEntitySignature(e entity.Entity) error {
	_, ok := r.signatures[e]
	if !ok {
		return ErrEntityDoesNotHaveSignature
	}
	delete(r.signatures, e)

	return nil
}

func (r *defaultRegistry) AddComponent(e entity.Entity, id component.ComponentId) error {
	sign, ok := r.signatures[e]
	if !ok {
		return ErrEntityDoesNotHaveSignature
	}

	sign.AddComponent(id)
	return nil
}

func (r *defaultRegistry) RemoveComponent(e entity.Entity, id component.ComponentId) error {
	sign, ok := r.signatures[e]
	if !ok {
		return ErrEntityDoesNotHaveSignature
	}

	sign.RemoveComponent(id)

	return nil
}

func (r *defaultRegistry) HasComponent(e entity.Entity, id component.ComponentId) bool {
	sign, ok := r.signatures[e]
	if !ok {
		return false
	}

	return sign.HasComponent(id)
}

func (r *defaultRegistry) GetComponentId(c interface{}) component.ComponentId {
	t := reflect.TypeOf(c)

	id, ok := r.types[t]
	if !ok {
		r.types[t] = r.next_id
		r.next_id++
		id = r.next_id - 1
	}

	return id
}

func (r *defaultRegistry) GetSignatureFromTypes(types []interface{}) Signature {
	sign := NewSignature()

	for _, t := range types {
		id := r.GetComponentId(t)
		sign.AddComponent(id)
	}

	return sign
}

func (r *defaultRegistry) FindMatchingEntities(matcher Signature) []entity.Entity {
	entities := make([]entity.Entity, 0)

	for e, s := range r.signatures {
		if s.Contains(matcher) {
			entities = append(entities, e)
		}
	}

	return entities
}
