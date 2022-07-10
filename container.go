package gecs

type Container interface {
	Add(Entity, interface{}) error
	Emplace(Entity, interface{})
	Remove(Entity) error
	Get(Entity) (interface{}, error)
	Has(Entity) bool
}

type ContainerType int

const (
	SPARSE_ARRAY_CONTAINER ContainerType = iota
	NULL_CONTAINER
	HASHMAP_CONTAINER
)
