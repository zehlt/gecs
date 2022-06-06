package entity

type Entity struct {
	id         int
	generation uint64
}

func (e Entity) Id() int {
	return e.id
}
