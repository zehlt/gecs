package entity

import (
	"errors"
)

var (
	ErrEntityDoesNotExist           = errors.New("entity does not exist")
	ErrInternalUnableToCreateEntity = errors.New("internal error: unable to create an entity")
)

type Arena interface {
	Create() (Entity, error)
	Destroy(Entity) error
	Exists(Entity) bool
	GetAll() []Entity
}

type arena struct {
	cells       []EntityCell
	current_gen uint64
	first_free  int
}

func NewArena() Arena {
	c := make([]EntityCell, 1)

	c[0] = EntityCell{
		t: END_CELL,
	}

	return &arena{
		cells:       c,
		current_gen: 0,
		first_free:  0,
	}
}

func (arena *arena) Create() (Entity, error) {
	e := arena.cells[arena.first_free]

	switch e.t {
	case EMPTY_CELL:
		arena.cells[arena.first_free] = EntityCell{
			t:          OCCUPIED_CELL,
			generation: arena.current_gen,
		}
		newEntity := Entity{id: arena.first_free, generation: arena.current_gen}
		arena.first_free = e.next

		return newEntity, nil

	case END_CELL:
		size := len(arena.cells)

		arena.cells[size-1] = EntityCell{
			t:          OCCUPIED_CELL,
			generation: arena.current_gen,
			next:       size,
		}

		arena.cells = append(arena.cells, EntityCell{t: END_CELL})
		arena.first_free = size

		return Entity{id: size - 1, generation: arena.current_gen}, nil

	default:
		return Entity{}, ErrInternalUnableToCreateEntity
	}
}

func (arena *arena) Destroy(e Entity) error {
	if !arena.Exists(e) {
		return ErrEntityDoesNotExist
	}

	arena.cells[e.id] = EntityCell{
		t:    EMPTY_CELL,
		next: arena.first_free,
	}
	arena.current_gen++
	arena.first_free = e.id

	return nil
}

func (arena *arena) Exists(e Entity) bool {
	if e.id >= len(arena.cells) {
		return false
	}

	cell := arena.cells[e.id]

	if cell.t == END_CELL || cell.t == EMPTY_CELL {
		return false
	}

	if cell.generation != e.generation {
		return false
	}

	return true
}

func (arena *arena) GetAll() []Entity {

	entities := make([]Entity, 0)

	for i := 0; i < len(arena.cells); i++ {
		cell := arena.cells[i]

		if cell.t == OCCUPIED_CELL {
			entities = append(entities, Entity{id: i, generation: cell.generation})
		}
	}

	return entities
}
