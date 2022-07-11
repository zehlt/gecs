package gecs

import "errors"

var (
	ErrQueryComponentPositionTooBig = errors.New("asked component position too big")
	ErrResourcePositionTooBig       = errors.New("asked resource position too big")
)

type Query interface {
	Entities(fn func(e Entity, comps []interface{}))
	// GetComponent(e Entity, n int) interface{}
	GetResource(n int) interface{}
}

type Access []interface{}
type Exclude []interface{}
type Resource []interface{}

type Args struct {
	Access
	Exclude
	Resource
}

type query struct {
	entities   []Entity
	components [][]interface{}
	resources  []interface{}
}

func (q *query) Entities(fn func(e Entity, comps []interface{})) {
	for i, entity := range q.entities {
		fn(entity, q.components[i])
	}
}

// func (q *query) GetComponent(e Entity, n int) interface{} {
// 	if n >= len(q.component_ids) {
// 		panic(ErrQueryComponentPositionTooBig)
// 	}

// 	componentId := q.component_ids[n]
// 	comp, err := q.w.GetComponentById(e, componentId)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return nil
// }

func (q *query) GetResource(n int) interface{} {
	if n >= len(q.resources) {
		panic(ErrQueryComponentPositionTooBig)
	}

	return q.resources[n]
}
