package entity

import "fmt"

type Entity struct {
	id         int
	generation uint64
}

func (e Entity) Id() int {
	return e.id
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity -> {ID: %d GEN: %d}", e.id, e.generation)
}
