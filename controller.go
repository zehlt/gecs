package gecs

import (
	"fmt"

	"github.com/zehlt/datt"
)

type Controller interface {
	CreateEntity() Entity
	AddComponent(e Entity, comp interface{})
	EmplaceComponent(e Entity, comp interface{})
	DestroyEntity(e Entity)
	Emit(t SignalType, data interface{})
	Execute()
}

type command interface {
	Execute(w World, s Scheduler)
}

type controller struct {
	w        World
	s        Scheduler
	commands datt.Queue[command]
}

func newController(w World, s Scheduler) Controller {
	return &controller{
		s:        s,
		w:        w,
		commands: datt.NewQueue[command](),
	}
}

func (c *controller) Emit(t SignalType, data interface{}) {
	c.commands.Enqueue(&emitCommand{t: t, data: data})
}

// // TODO: should be done at the end of the stage
func (c *controller) CreateEntity() Entity {
	e, err := c.w.CreateEntity()
	if err != nil {
		panic(err)
	}

	return e
}

func (c *controller) AddComponent(e Entity, comp interface{}) {
	c.commands.Enqueue(&addComponentCommand{e: e, comp: comp})
}

func (c *controller) EmplaceComponent(e Entity, comp interface{}) {
	c.commands.Enqueue(&emplaceComponentCommand{e: e, comp: comp})
}

func (c *controller) RemoveComponent(e Entity, comp interface{}) {
	c.commands.Enqueue(&removeComponentCommand{e: e, comp: comp})
}

func (c *controller) DestroyEntity(e Entity) {
	c.commands.Enqueue(&destroyEntityCommand{e: e})
}

func (c *controller) Execute() {
	if c.commands.IsEmpty() {
		return
	}

	len := c.commands.Length()
	for i := 0; i < len; i++ {
		cmd, err := c.commands.Dequeue()
		if err != nil {
			panic(err)
		}
		cmd.Execute(c.w, c.s)
	}
}

type destroyEntityCommand struct {
	e Entity
}

func (c *destroyEntityCommand) Execute(w World, s Scheduler) {
	fmt.Println("DESTROY ENTITY!!")
	w.DestroyEntity(c.e)
}

type emitCommand struct {
	t    SignalType
	data interface{}
}

func (c *emitCommand) Execute(w World, s Scheduler) {
	fmt.Println("EMIT SIGNAL")
	s.Emit(c.t, c.data)
}

type addComponentCommand struct {
	e    Entity
	comp interface{}
}

func (c *addComponentCommand) Execute(w World, s Scheduler) {
	fmt.Println("ADD COMPONENT")
	w.AddComponent(c.e, c.comp)
}

type emplaceComponentCommand struct {
	e    Entity
	comp interface{}
}

func (c *emplaceComponentCommand) Execute(w World, s Scheduler) {
	fmt.Println("EMPLACE COMPONENT SIGNAL")
	w.EmplaceComponent(c.e, c.comp)
}

type removeComponentCommand struct {
	e    Entity
	comp interface{}
}

func (c *removeComponentCommand) Execute(w World, s Scheduler) {
	fmt.Println("REMOVE COMPONENT SIGNAL")

	w.RemoveComponent(c.e, c.comp)
}
