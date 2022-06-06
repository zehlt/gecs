// deprecated
package component

// type Position struct {
// 	x int
// 	y int
// }

// type Movement struct {
// 	vel float64
// 	acc float64
// }

// func TestAddComponent(t *testing.T) {
// 	store := NewSparseStore()
// 	e1 := entity.Entity{Id: 10}

// 	err := store.Add(e1, Position{x: 23, y: 25})
// 	require.NoError(t, err)

// 	require.True(t, store.Has(e1, Position{}))
// 	require.False(t, store.Has(e1, Movement{}))
// }

// func TestAddTwoDifferentComponent(t *testing.T) {
// 	store := NewSparseStore()
// 	e1 := entity.Entity{Id: 10}

// 	err := store.Add(e1, Movement{vel: 23, acc: 5})
// 	require.NoError(t, err)

// 	err = store.Add(e1, Position{x: 23, y: 25})
// 	require.NoError(t, err)

// 	require.True(t, store.Has(e1, Position{}))
// 	require.True(t, store.Has(e1, Movement{}))
// }

// func TestAddTwoSameComponent(t *testing.T) {
// 	store := NewSparseStore()
// 	e1 := entity.Entity{Id: 10}

// 	err := store.Add(e1, Position{x: 23, y: 25})
// 	require.NoError(t, err)

// 	err = store.Add(e1, Position{x: 9, y: 87})
// 	require.Error(t, err)
// 	require.Equal(t, ErrComponentAlreadyOwnByEntity, err)
// }
