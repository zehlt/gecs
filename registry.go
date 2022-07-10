package gecs

import (
	"errors"
	"reflect"
)

var (
	ErrEntityAlreadyHasSignature  = errors.New("entity already has a signature")
	ErrEntityDoesNotHaveSignature = errors.New("entity does not have a signature")
)

type Registry interface {
	CreateEntitySignature(Entity) error
	GetEntitySignature(Entity) (Signature, error)
	DestroyEntitySignature(Entity) error

	AddComponent(Entity, ComponentId) error
	RemoveComponent(Entity, ComponentId) error
	HasComponent(Entity, ComponentId) bool
	GetComponentId(c interface{}) ComponentId

	GetSignatureFromTypes([]interface{}) Signature
	FindMatchingEntities(Signature) []Entity
}

type defaultRegistry struct {
	signatures map[Entity]Signature

	types   map[reflect.Type]ComponentId
	next_id ComponentId
}

func NewRegistry() Registry {
	return &defaultRegistry{
		signatures: make(map[Entity]Signature),
		types:      make(map[reflect.Type]ComponentId),
		next_id:    0,
	}
}

func (r *defaultRegistry) CreateEntitySignature(e Entity) error {
	_, ok := r.signatures[e]
	if ok {
		return ErrEntityAlreadyHasSignature
	}

	r.signatures[e] = NewSignature()

	return nil
}

func (r *defaultRegistry) GetEntitySignature(e Entity) (Signature, error) {
	sign, ok := r.signatures[e]
	if !ok {
		return nil, ErrEntityDoesNotHaveSignature
	}

	return sign, nil
}

func (r *defaultRegistry) DestroyEntitySignature(e Entity) error {
	_, ok := r.signatures[e]
	if !ok {
		return ErrEntityDoesNotHaveSignature
	}
	delete(r.signatures, e)

	return nil
}

func (r *defaultRegistry) AddComponent(e Entity, id ComponentId) error {
	sign, ok := r.signatures[e]
	if !ok {
		return ErrEntityDoesNotHaveSignature
	}

	sign.AddComponent(id)
	return nil
}

func (r *defaultRegistry) RemoveComponent(e Entity, id ComponentId) error {
	sign, ok := r.signatures[e]
	if !ok {
		return ErrEntityDoesNotHaveSignature
	}

	sign.RemoveComponent(id)

	return nil
}

func (r *defaultRegistry) HasComponent(e Entity, id ComponentId) bool {
	sign, ok := r.signatures[e]
	if !ok {
		return false
	}

	return sign.HasComponent(id)
}

func (r *defaultRegistry) GetComponentId(c interface{}) ComponentId {
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

func (r *defaultRegistry) FindMatchingEntities(matcher Signature) []Entity {
	entities := make([]Entity, 0)

	for e, s := range r.signatures {
		if s.Contains(matcher) {
			entities = append(entities, e)
		}
	}

	return entities
}
