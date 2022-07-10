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

type Access []interface{}
type Exclude []interface{}
type Resource []interface{}

type QueryMaker interface {
	Create(r Resource, a Access, e Exclude) Query
}

type queryMaker struct {
	w World
}

func NewQueryMaker(w World) QueryMaker {
	return &queryMaker{w: w}
}

func (qm *queryMaker) Create(r Resource, a Access, e Exclude) Query {
	access_sign := qm.w.GetSignatureFromTypes(a)
	exclude_sign := qm.w.GetSignatureFromTypes(e)

	component_ids := make([]ComponentId, len(a))
	resources := make([]interface{}, len(r))

	for i, t := range a {
		component_ids[i] = qm.w.GetComponentId(t)
	}

	for i, t := range r {
		co, err := qm.w.GetResource(t)
		if err != nil {
			panic(err)
		}
		resources[i] = co
	}

	return &query{
		w:             qm.w,
		component_ids: component_ids,
		access_sign:   access_sign,
		exclude_sign:  exclude_sign,
		resources:     resources,
	}
}
