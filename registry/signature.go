package registry

import "github.com/zehlt/gecs/component"

// footprint
type Signature interface {
	AddComponent(id component.ComponentId)
	RemoveComponent(id component.ComponentId)
	HasComponent(id component.ComponentId) bool
}

type signature struct {
	bitset *Bitset
}

func NewSignature() Signature {
	return &signature{
		bitset: NewBitset(),
	}
}

func (s *signature) AddComponent(id component.ComponentId) {
	s.bitset.Set(int(id), true)
}

func (s *signature) RemoveComponent(id component.ComponentId) {
	s.bitset.Set(int(id), false)
}

func (s *signature) HasComponent(id component.ComponentId) bool {
	return s.bitset.Get(int(id))
}
