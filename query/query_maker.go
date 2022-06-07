package query

import (
	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/component"
)

type Access []interface{}

type Exclude []interface{}

type QueryMaker interface {
	Create(a Access, e Exclude) Query
}

type queryMaker struct {
	w gecs.World
}

func NewQueryMaker(w gecs.World) QueryMaker {
	return &queryMaker{w: w}
}

func (qm *queryMaker) Create(a Access, e Exclude) Query {
	access_sign := qm.w.GetSignatureFromTypes(a)
	exclude_sign := qm.w.GetSignatureFromTypes(e)

	component_ids := make([]component.ComponentId, len(a))

	for i, t := range a {
		component_ids[i] = qm.w.GetComponentId(t)
	}

	return &query{
		w:             qm.w,
		component_ids: component_ids,
		access_sign:   access_sign,
		exclude_sign:  exclude_sign,
	}
}
