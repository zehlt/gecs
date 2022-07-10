package gecs

type NullEmptySruct struct {
}

type Null struct {
}

func newNull() Container {
	return &Null{}
}

func (n *Null) Add(e Entity, c interface{}) error {
	return nil
}

func (n *Null) Emplace(e Entity, c interface{}) {
}

func (n *Null) Remove(e Entity) error {
	return nil
}

func (n *Null) Get(e Entity) (interface{}, error) {
	return NullEmptySruct{}, nil
}

func (n *Null) Has(e Entity) bool {
	return false
}
