package command

import (
	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/entity"
)

type Controller interface {
	CreateEntity()
	DestroyEntity(e entity.Entity)
	Execute()
}

type controller struct {
	w gecs.World
	// TODO: maybe switch from slice to queue
	commands []Command
}

func NewController(w gecs.World) Controller {
	return &controller{
		w:        w,
		commands: make([]Command, 0),
	}
}

func (c *controller) CreateEntity() {

}

func (c *controller) DestroyEntity(e entity.Entity) {
	c.commands = append(c.commands, &DestroyEntityCommand{e: e})
}

func (c *controller) Execute() {
	if len(c.commands) <= 0 {
		return
	}

	for _, cmd := range c.commands {
		cmd.Execute(c.w)
	}

	c.commands = make([]Command, 0)
}
