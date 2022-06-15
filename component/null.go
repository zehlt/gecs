package component

import "github.com/zehlt/gecs/entity"

type NullEmptySruct struct {
}

type Null struct {
}

func newNull() Container {
	return &Null{}
}

func (n *Null) Add(e entity.Entity, c interface{}) error {
	return nil
}

func (n *Null) Emplace(e entity.Entity, c interface{}) {
}

func (n *Null) Remove(e entity.Entity) error {
	return nil
}

func (n *Null) Get(e entity.Entity) (interface{}, error) {
	return NullEmptySruct{}, nil
}

func (n *Null) Has(e entity.Entity) bool {
	return false
}
