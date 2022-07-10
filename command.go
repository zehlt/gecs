package gecs

import (
	"fmt"
)

type Command interface {
	Execute(w World)
}

type DestroyEntityCommand struct {
	e Entity
}

func (c *DestroyEntityCommand) Execute(w World) {
	fmt.Println("DESTROY ENTITY!!")
	w.DestroyEntity(c.e)
}
