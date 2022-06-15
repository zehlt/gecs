package component

import "github.com/zehlt/gecs/entity"

type Container interface {
	Add(entity.Entity, interface{}) error
	Emplace(entity.Entity, interface{})
	Remove(entity.Entity) error
	Get(entity.Entity) (interface{}, error)
	Has(entity.Entity) bool
}

type ContainerType int

const (
	SPARSE_ARRAY_CONTAINER ContainerType = iota
	NULL_CONTAINER
	HASHMAP_CONTAINER
)
