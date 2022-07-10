package gecs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUniqueEntityWithoutError(t *testing.T) {
	arena := NewArena()

	entity, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, entity.id)
	require.Equal(t, uint64(0), entity.generation)
}

func TestCreateEntityAlreadyDestroyed(t *testing.T) {
	arena := NewArena()

	e1, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e1.id)
	require.Equal(t, uint64(0), e1.generation)

	arena.Destroy(e1)

	e2, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e2.id)
	require.Equal(t, uint64(1), e2.generation)
}

func TestSequenceCreateDestroyCreateCreateCreateEntity(t *testing.T) {
	arena := NewArena()

	e1, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e1.id)
	require.Equal(t, uint64(0), e1.generation)

	arena.Destroy(e1)

	e2, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e2.id)
	require.Equal(t, uint64(1), e2.generation)

	e3, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 1, e3.id)
	require.Equal(t, uint64(1), e3.generation)

	e4, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 2, e4.id)
	require.Equal(t, uint64(1), e4.generation)
}

func TestCreateThreeEntitiesWithoutError(t *testing.T) {
	arena := NewArena()

	e1, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e1.id)
	require.Equal(t, uint64(0), e1.generation)

	e2, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, int(1), e2.id)
	require.Equal(t, uint64(0), e2.generation)

	e3, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, int(2), e3.id)
	require.Equal(t, uint64(0), e3.generation)
}

func TestExistsOneEntityWithoutError(t *testing.T) {
	arena := NewArena()

	e1, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e1.id)
	require.Equal(t, uint64(0), e1.generation)

	require.True(t, arena.Exists(e1))
	require.False(t, arena.Exists(Entity{id: 1, generation: 4}))
}

func TestDestroyOneEntityWithoutError(t *testing.T) {

	arena := NewArena()

	e1, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e1.id)
	require.Equal(t, uint64(0), e1.generation)

	err = arena.Destroy(e1)
	require.NoError(t, err)
	require.False(t, arena.Exists(e1))
}

func TestDestroyEntitySameTwice(t *testing.T) {

	arena := NewArena()

	e1, err := arena.Create()
	require.NoError(t, err)
	require.Equal(t, 0, e1.id)
	require.Equal(t, uint64(0), e1.generation)

	err = arena.Destroy(e1)
	require.NoError(t, err)
	require.False(t, arena.Exists(e1))

	err = arena.Destroy(e1)
	require.Error(t, err)
	require.Equal(t, err, ErrEntityDoesNotExist)
}

func TestDestroyAllEntities(t *testing.T) {
	arena := NewArena()

	e1, _ := arena.Create()
	e2, _ := arena.Create()
	e3, _ := arena.Create()
	e4, _ := arena.Create()

	err := arena.Destroy(e1)
	require.NoError(t, err)

	err = arena.Destroy(e2)
	require.NoError(t, err)

	err = arena.Destroy(e3)
	require.NoError(t, err)

	err = arena.Destroy(e4)
	require.NoError(t, err)

	err = arena.Destroy(e1)
	require.Error(t, err)

	require.False(t, arena.Exists(e1))
	require.False(t, arena.Exists(e2))
	require.False(t, arena.Exists(e3))
	require.False(t, arena.Exists(e4))
}
