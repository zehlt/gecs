package gecs

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zehlt/datt"
)

type signature struct {
	bitset *datt.Bitset
}

func newSignature(size int) signature {
	bit, _ := datt.NewBitset(size)
	return signature{
		bitset: bit,
	}
}

func (s *signature) EmplaceComponent(t ComponentType) {
	s.bitset.Set(int(t), true)
}

func (s *signature) RemoveComponent(t ComponentType) {
	s.bitset.Set(int(t), false)
}

func (s *signature) HasComponent(t ComponentType) bool {
	return s.bitset.Get(int(t))
}

func (s *signature) Contain(sign signature) bool {
	return s.bitset.Contain(sign.bitset)
}

func (s *signature) Crossing(sign signature) bool {
	return s.bitset.Crossing(sign.bitset)
}

type UID [16]byte

type Entity struct {
	uid UID
}

type registry struct {
	entities    map[Entity]signature
	biggestType int
}

func newRegistry() *registry {
	return &registry{
		entities: make(map[Entity]signature),
	}
}

func (r *registry) RegisterComponent(t ComponentType) {
	currentType := int(t)

	if currentType > r.biggestType {
		r.biggestType = currentType
	}
}

func (r *registry) CreateEntity() Entity {
	e := Entity{
		uid: UID(uuid.New()),
	}

	r.entities[e] = newSignature(r.biggestType + 1)
	return e
}

func (r *registry) DestroyEntity(e Entity) error {
	_, ok := r.entities[e]
	if !ok {
		return fmt.Errorf("entity not registered")
	}

	delete(r.entities, e)
	return nil
}

func (r *registry) EmplaceComponent(e Entity, t ComponentType) error {
	signature, ok := r.entities[e]
	if !ok {
		return fmt.Errorf("entity not registered")
	}

	signature.EmplaceComponent(t)
	return nil
}

func (r *registry) RemoveComponent(e Entity, t ComponentType) error {
	signature, ok := r.entities[e]
	if !ok {
		return fmt.Errorf("entity not registered")
	}
	signature.RemoveComponent(t)
	return nil
}

func (r *registry) HasComponent(e Entity, t ComponentType) bool {
	signature, ok := r.entities[e]
	if ok {
		return signature.HasComponent(t)
	}
	return false
}

func (r *registry) getMatchingEntities(access signature, exclude signature) []Entity {
	ents := []Entity{}

	for e, sign := range r.entities {
		if sign.Contain(access) {
			if !sign.Crossing(exclude) {
				ents = append(ents, e)
			}
		}
	}

	return ents
}

func (r *registry) newSignatureFromTypes(types []ComponentType) signature {
	bitset, _ := datt.NewBitset(r.biggestType + 1)

	for _, t := range types {
		bitset.Set(int(t), true)
	}

	return signature{
		bitset: bitset,
	}
}
