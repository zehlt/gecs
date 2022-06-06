package registry

import (
	"reflect"

	"github.com/zehlt/gecs/component"
)

type Registry interface {
	GetComponentId(c interface{}) component.ComponentId
}

type defaultRegistry struct {
	types   map[reflect.Type]component.ComponentId
	next_id component.ComponentId
}

func NewRegistry() Registry {
	return &defaultRegistry{
		types:   make(map[reflect.Type]component.ComponentId),
		next_id: 0,
	}
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
