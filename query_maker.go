package gecs

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
