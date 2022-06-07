package query

import (
	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/entity"
	"github.com/zehlt/gecs/signature"
)

type QueryData struct {
	E entity.Entity
}

type Query interface {
	ForEach(fn func(data QueryData))
}

type query struct {
	w            gecs.World
	access_sign  signature.Signature
	exclude_sign signature.Signature
}

func (q *query) ForEach(fn func(QueryData)) {
	entities := q.w.FindMatchingEntities(q.access_sign)

	for _, entity := range entities {
		fn(QueryData{E: entity})
	}
}
