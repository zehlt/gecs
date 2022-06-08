package command

import (
	"fmt"

	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/entity"
)

type Command interface {
	Execute(w gecs.World)
}

type DestroyEntityCommand struct {
	e entity.Entity
}

func (c *DestroyEntityCommand) Execute(w gecs.World) {
	fmt.Println("DESTROY ENTITY!!")
	w.DestroyEntity(c.e)
}
