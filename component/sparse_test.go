package component

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zehlt/gecs/entity"
)

type Position struct {
	x int
	y int
}

type Movement struct {
	vel float64
	acc float64
}

func TestAddComponent(t *testing.T) {
	store := NewSparseStore()
	e1 := entity.Entity{Id: 10}

	err := store.Add(e1, Position{x: 23, y: 25})
	require.NoError(t, err)

	require.True(t, store.Has(e1, Position{}))
	require.False(t, store.Has(e1, Movement{}))
}
