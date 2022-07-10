package gecs

import (
	"errors"
)

var (
	ErrEntityDoesNotExist           = errors.New("entity does not exist")
	ErrInternalUnableToCreateEntity = errors.New("internal error: unable to create an entity")
)

type entityCellType int

const (
	EMPTY_CELL entityCellType = iota
	END_CELL
	OCCUPIED_CELL
)

type entityCell struct {
	t          entityCellType
	generation uint64
	next       int
}

type arena interface {
	Create() (Entity, error)
	Destroy(Entity) error
	Exists(Entity) bool
	GetAll() []Entity
}

type defaultArena struct {
	cells       []entityCell
	current_gen uint64
	first_free  int
}

func newArena() arena {
	c := make([]entityCell, 1)

	c[0] = entityCell{
		t: END_CELL,
	}

	return &defaultArena{
		cells:       c,
		current_gen: 0,
		first_free:  0,
	}
}

func (a *defaultArena) Create() (Entity, error) {
	e := a.cells[a.first_free]

	switch e.t {
	case EMPTY_CELL:
		a.cells[a.first_free] = entityCell{
			t:          OCCUPIED_CELL,
			generation: a.current_gen,
		}
		newEntity := Entity{id: a.first_free, generation: a.current_gen}
		a.first_free = e.next

		return newEntity, nil

	case END_CELL:
		size := len(a.cells)

		a.cells[size-1] = entityCell{
			t:          OCCUPIED_CELL,
			generation: a.current_gen,
			next:       size,
		}

		a.cells = append(a.cells, entityCell{t: END_CELL})
		a.first_free = size

		return Entity{id: size - 1, generation: a.current_gen}, nil

	default:
		return Entity{}, ErrInternalUnableToCreateEntity
	}
}

func (a *defaultArena) Destroy(e Entity) error {
	if !a.Exists(e) {
		return ErrEntityDoesNotExist
	}

	a.cells[e.id] = entityCell{
		t:    EMPTY_CELL,
		next: a.first_free,
	}
	a.current_gen++
	a.first_free = e.id

	return nil
}

func (a *defaultArena) Exists(e Entity) bool {
	if e.id >= len(a.cells) {
		return false
	}

	cell := a.cells[e.id]

	if cell.t == END_CELL || cell.t == EMPTY_CELL {
		return false
	}

	if cell.generation != e.generation {
		return false
	}

	return true
}

func (a *defaultArena) GetAll() []Entity {

	entities := make([]Entity, 0)

	for i := 0; i < len(a.cells); i++ {
		cell := a.cells[i]

		if cell.t == OCCUPIED_CELL {
			entities = append(entities, Entity{id: i, generation: cell.generation})
		}
	}

	return entities
}
