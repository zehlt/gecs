package query

import (
	"errors"

	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/component"
	"github.com/zehlt/gecs/entity"
	"github.com/zehlt/gecs/signature"
)

var (
	ErrQueryComponentPositionTooBig = errors.New("asked component position too big")
)

type Query interface {
	Entities(fn func(e entity.Entity))
	GetComponent(e entity.Entity, n int) interface{}
}

type query struct {
	w             gecs.World
	component_ids []component.ComponentId

	access_sign  signature.Signature
	exclude_sign signature.Signature
}

func (q *query) Entities(fn func(e entity.Entity)) {
	entities := q.w.FindMatchingEntities(q.access_sign)

	for _, entity := range entities {
		fn(entity)
	}
}

func (q *query) GetComponent(e entity.Entity, n int) interface{} {
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
