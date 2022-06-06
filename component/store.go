package component

import (
	"github.com/zehlt/gecs/entity"
)

type Store interface {
	Add(entity.Entity, interface{}) error
	Remove(entity.Entity, interface{}) error
	Get(entity.Entity, interface{}) (interface{}, error)
	Has(entity.Entity, interface{}) bool
}
