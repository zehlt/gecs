package gecs

type Controller interface {
	CreateEntity() Entity
	AddComponent(e Entity, comp interface{})
	EmplaceComponent(e Entity, comp interface{})
	DestroyEntity(e Entity)
	Execute()
}

type controller struct {
	w World
	// TODO: maybe switch from slice to queue
	commands []Command
}

func NewController(w World) Controller {
	return &controller{
		w:        w,
		commands: make([]Command, 0),
	}
}

func (c *controller) CreateEntity() Entity {
	// TODO: should be done at the end of the stage
	e, err := c.w.CreateEntity()
	if err != nil {
		panic(err)
	}

	return e
}

func (c *controller) AddComponent(e Entity, comp interface{}) {
	err := c.w.AddComponent(e, comp)
	if err != nil {
		panic(err)
	}
}

func (c *controller) EmplaceComponent(e Entity, comp interface{}) {
	err := c.w.EmplaceComponent(e, comp)
	if err != nil {
		panic(err)
	}
}

func (c *controller) DestroyEntity(e Entity) {
	err := c.w.DestroyEntity(e)
	if err != nil {
		panic(err)
	}
	// c.commands = append(c.commands, &DestroyEntityCommand{e: e})
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
