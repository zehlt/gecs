package gecs

import (
	"errors"
)

var (
	ErrQueryComponentPositionTooBig = errors.New("asked component position too big")
	ErrResourcePositionTooBig       = errors.New("asked resource position too big")
)

type Query interface {
	Entities(fn func(e Entity))
	GetComponent(e Entity, n int) interface{}
	GetResource(n int) interface{}
}

type query struct {
	w             World
	component_ids []ComponentId
	resources     []interface{}

	access_sign  Signature
	exclude_sign Signature
}

func (q *query) Entities(fn func(e Entity)) {
	entities := q.w.FindMatchingEntities(q.access_sign)

	for _, entity := range entities {
		fn(entity)
	}
}

func (q *query) GetComponent(e Entity, n int) interface{} {
	if n >= len(q.component_ids) {
		panic(ErrQueryComponentPositionTooBig)
	}

	componentId := q.component_ids[n]
	comp, err := q.w.GetComponentById(e, componentId)
	if err != nil {
		panic(err)
	}

	return comp
}

func (q *query) GetResource(n int) interface{} {
	if n >= len(q.resources) {
		panic(ErrQueryComponentPositionTooBig)
	}

	return q.resources[n]
}
