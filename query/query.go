package query

import (
	"reflect"

	"github.com/zehlt/gecs/registry"
)

type Access []interface{}

type Exclude []interface{}

type Query struct {
	access_types  []reflect.Type
	exclude_types []reflect.Type
}

func Make(r registry.Registry, a Access, e Exclude) Query {
	access := make([]reflect.Type, len(a))
	exclude := make([]reflect.Type, len(e))

	for i, t := range a {
		access[i] = reflect.TypeOf(t)
	}

	for i, t := range e {
		exclude[i] = reflect.TypeOf(t)
	}

	return Query{
		access_types:  access,
		exclude_types: exclude,
	}
}

func (q *Query) Iter() (interface{}, bool) {

	return nil, false
}
