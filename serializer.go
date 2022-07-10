package gecs

import (
	"bytes"
	"encoding/gob"
)

type Position struct {
	X int
	Y int
}

type Speed struct {
	V float64
	A float64
}

type Life struct {
	HP int
}

type Enemy struct {
}

type Player struct {
}

// Add insert entity
// Add Tag component
// Add reference with types of containers
// Add update world instead of rebuild again
// Make sheduler working on new world without requiery

type Couple struct {
	E          Entity
	Components []interface{}
}

type Reference struct {
}

type Snap struct {
	// Couples []Couple
	Comps [][]interface{}
}

type encoder interface {
	Register(interface{})
	Encode(Snap) ([]byte, error)
	Decode([]byte) (Snap, error)
}

func newGobEncoder() encoder {
	return &defaultEncoder{}
}

type defaultEncoder struct {
}

func (e *defaultEncoder) Register(c interface{}) {
	gob.Register(c)
}

func (e *defaultEncoder) Encode(s Snap) ([]byte, error) {
	b := bytes.Buffer{}

	ego := gob.NewEncoder(&b)
	err := ego.Encode(s)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (e *defaultEncoder) Decode(data []byte) (Snap, error) {
	var snap Snap
	b := bytes.Buffer{}

	_, err := b.Write(data)
	if err != nil {
		return Snap{}, err
	}

	ego := gob.NewDecoder(&b)
	err = ego.Decode(&snap)
	if err != nil {
		return Snap{}, err
	}

	return snap, nil
}

type Serializer interface {
	Register(c interface{})
	Serialize(w World) ([]byte, error)
	Deserialize([]byte) (World, error)
	Merge([]byte, World) error
}

func NewSerializer() Serializer {
	return &serializer{
		encoder: newGobEncoder(),
	}
}

type serializer struct {
	encoder encoder
}

func (s *serializer) Register(c interface{}) {
	s.encoder.Register(c)
}

func (s *serializer) Serialize(w World) ([]byte, error) {

	entities := w.GetAllEntities()

	var snap Snap

	snap.Comps = make([][]interface{}, 0)

	for _, entity := range entities {
		components, _ := w.GetAllComponentsFromEntity(entity)
		snap.Comps = append(snap.Comps, components)
	}

	return s.encoder.Encode(snap)
}

func (s *serializer) Deserialize(data []byte) (World, error) {
	snap, err := s.encoder.Decode(data)
	if err != nil {
		panic(err)
	}
	w := DefaultWorld()

	for _, components := range snap.Comps {
		e, err := w.CreateEntity()
		if err != nil {
			panic(err)
		}

		for _, c := range components {
			w.RegisterComponent(c, SPARSE_ARRAY_CONTAINER)
			w.AddComponent(e, c)
		}
	}

	return w, nil
}

func (s *serializer) Merge(data []byte, w World) error {

	// snap, err := s.encoder.Decode(data)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(snap.Comps); i++ {
	// 	// e, err := w.CreateEntity()
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }

	// 	// components := snap.Comps[i]
	// 	// for j := 0; j < len(components); j++ {
	// 	// 	// p := reflect.New(reflect.TypeOf(components[j]))
	// 	// 	// p.Elem().Set(reflect.ValueOf(components[j]))
	// 	// 	// comp := p.Interface()

	// 	// 	// TODO: made the container not specific
	// 	// 	w.RegisterComponent(components[j], SPARSE_ARRAY_CONTAINER)
	// 	// 	w.AddComponent(e, components[j])
	// 	// }
	// }

	return nil
}

// p := reflect.New(reflect.TypeOf(components[j]))
// p.Elem().Set(reflect.ValueOf(components[j]))
// comp := p.Interface()

// type parser struct {
// }

// func (p *parser) Parse(fn interface{}) {

// 	t := reflect.TypeOf(fn)
// 	fmt.Println("Function arguments:")
// 	for i := 0; i < t.NumIn(); i++ {
// 		fmt.Printf(" %d. %v\n", i, t.In(i))
// 	}
// 	fmt.Println("Function return values:")
// 	for i := 0; i < t.NumOut(); i++ {
// 		fmt.Printf(" %d. %v\n", i, t.Out(i))
// 	}
// }
