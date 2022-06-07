package query

import "github.com/zehlt/gecs"

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

	return &query{w: qm.w, access_sign: access_sign, exclude_sign: exclude_sign}
}
