package gecs

type container interface {
	Add(Entity, interface{}) error
	Emplace(Entity, interface{})
	Remove(Entity) error
	Get(Entity) (interface{}, error)
	Has(Entity) bool
}

type ContainerType int

const (
	SPARSE_ARRAY_CONTAINER ContainerType = iota
	NULL_CONTAINER
	HASHMAP_CONTAINER
)

// HASHMAP

type hashmap struct {
	m map[int]interface{}
}

func newHashmap() container {
	return &hashmap{
		m: make(map[int]interface{}),
	}
}

func (h *hashmap) Add(e Entity, c interface{}) error {
	if h.Has(e) {
		return ErrEntityAlreadyHasComponent
	}

	h.m[e.Id()] = c

	return nil
}

func (h *hashmap) Emplace(e Entity, c interface{}) {
	h.m[e.Id()] = c
}

func (h *hashmap) Remove(e Entity) error {
	if !h.Has(e) {
		return ErrEntityDoesNotHaveComponent
	}

	h.m[e.Id()] = nil

	return nil
}

func (h *hashmap) Get(e Entity) (interface{}, error) {
	if !h.Has(e) {
		return nil, ErrEntityDoesNotHaveComponent
	}

	return h.m[e.Id()], nil
}

func (h *hashmap) Has(e Entity) bool {
	_, ok := h.m[e.Id()]

	return ok
}

// NULL

type NullEmptySruct struct {
}

type Null struct {
}

func newNull() container {
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

// SPARSE

type sparseArray struct {
	components []interface{}
}

func newSparseArray() container {
	return &sparseArray{
		components: make([]interface{}, 0),
	}
}

func (s *sparseArray) Add(e Entity, c interface{}) error {
	if e.Id() >= len(s.components) {
		for i := e.Id() - len(s.components); i >= 0; i-- {
			s.components = append(s.components, nil)
		}
	}

	if s.components[e.Id()] != nil {
		return ErrEntityAlreadyHasComponent
	}

	s.components[e.Id()] = c

	return nil
}

func (s *sparseArray) Emplace(e Entity, c interface{}) {
	if e.Id() >= len(s.components) {
		for i := e.Id() - len(s.components); i >= 0; i-- {
			s.components = append(s.components, nil)
		}
	}

	s.components[e.Id()] = c
}

func (s *sparseArray) Remove(e Entity) error {
	if !s.Has(e) {
		return ErrEntityDoesNotHaveComponent
	}

	s.components[e.Id()] = nil

	return nil
}

func (s *sparseArray) Get(e Entity) (interface{}, error) {
	if !s.Has(e) {
		return nil, ErrEntityDoesNotHaveComponent
	}

	return s.components[e.Id()], nil
}

func (s *sparseArray) Has(e Entity) bool {
	if e.Id() >= len(s.components) {
		return false
	}

	if s.components[e.Id()] == nil {
		return false
	}

	return true
}
