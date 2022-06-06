package component

import "github.com/zehlt/gecs/entity"

type Container interface {
	Add(entity.Entity, interface{}) error
	Remove(entity.Entity) error
	Get(entity.Entity) (interface{}, error)
	Has(entity.Entity) bool
}
