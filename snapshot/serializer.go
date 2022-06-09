package snapshot

import (
	"fmt"
	"reflect"

	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/component"
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

type Serializer interface {
	Serialize(w gecs.World) ([]byte, error)
	Deserialize([]byte) (gecs.World, error)
}

func NewSerializer() Serializer {
	return &serializer{
		encoder: NewGobEncoder(),
	}
}

type serializer struct {
	encoder Encoder
}

func (s *serializer) Serialize(w gecs.World) ([]byte, error) {
	entities := w.GetAllEntities()

	var snap Snap

	snap.Comps = make([][]interface{}, 0)

	for _, entity := range entities {
		components, _ := w.GetAllComponentsFromEntity(entity)
		snap.Comps = append(snap.Comps, components)
	}

	return s.encoder.Encode(snap)
}

func (s *serializer) Deserialize(data []byte) (gecs.World, error) {
	snap, err := s.encoder.Decode(data)
	if err != nil {
		panic(err)
	}
	w := gecs.DefaultWorld()

	for i := 0; i < len(snap.Comps); i++ {
		e, err := w.CreateEntity()
		if err != nil {
			panic(err)
		}
		fmt.Println(e)

		components := snap.Comps[i]
		for j := 0; j < len(components); j++ {
			p := reflect.New(reflect.TypeOf(components[j]))
			p.Elem().Set(reflect.ValueOf(components[j]))
			comp := p.Interface()

			// TODO: made the container not specific
			w.RegisterComponent(comp, component.SPARSE_ARRAY_CONTAINER)
			w.AddComponent(e, comp)
		}
	}

	return w, nil
}
