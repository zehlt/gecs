package gecs

import (
	"bytes"
	"encoding/gob"
)

type Encoder interface {
	Register(interface{})
	Encode(Snap) ([]byte, error)
	Decode([]byte) (Snap, error)
}

func NewGobEncoder() Encoder {
	return &encoder{}
}

type encoder struct {
}

func (e *encoder) Register(c interface{}) {
	gob.Register(c)
}

func (e *encoder) Encode(s Snap) ([]byte, error) {
	b := bytes.Buffer{}

	ego := gob.NewEncoder(&b)
	err := ego.Encode(s)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (e *encoder) Decode(data []byte) (Snap, error) {
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
