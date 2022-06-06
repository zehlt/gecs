package component

import "github.com/zehlt/gecs/entity"

type hashmap struct {
	m map[int]interface{}
}

func newHashmap() Container {
	return &hashmap{
		m: make(map[int]interface{}),
	}
}

func (h *hashmap) Add(e entity.Entity, c interface{}) error {
	if h.Has(e) {
		return ErrEntityAlreadyHasComponent
	}

	h.m[e.Id()] = c

	return nil
}

func (h *hashmap) Remove(e entity.Entity) error {
	if !h.Has(e) {
		return ErrEntityDoesNotHaveComponent
	}

	h.m[e.Id()] = nil

	return nil
}

func (h *hashmap) Get(e entity.Entity) (interface{}, error) {
	if !h.Has(e) {
		return nil, ErrEntityDoesNotHaveComponent
	}

	return h.m[e.Id()], nil
}

func (h *hashmap) Has(e entity.Entity) bool {
	_, ok := h.m[e.Id()]

	return ok
}
