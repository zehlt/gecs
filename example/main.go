package main

import (
	"log"

	"github.com/zehlt/gecs"
)

const (
	PLAYER_COMPONENT gecs.ComponentType = iota
	VELOCITY_COMPONENT
	POSITION_COMPONENT
)

type Player struct {
}

func (p Player) GetType() gecs.ComponentType {
	return PLAYER_COMPONENT
}

type Position struct {
	X int
	Y int
}

func (p Position) GetType() gecs.ComponentType {
	return POSITION_COMPONENT
}

type Velocity struct {
	D int
}

func (p Velocity) GetType() gecs.ComponentType {
	return VELOCITY_COMPONENT
}

func main() {
	world := gecs.NewWorld()
	err := world.RegisterComponent(PLAYER_COMPONENT, gecs.TAG_CONTAINER)
	log.Println(err)
	err = world.RegisterComponent(POSITION_COMPONENT, gecs.HASHMAP_CONTAINER)
	log.Println(err)
	err = world.RegisterComponent(VELOCITY_COMPONENT, gecs.HASHMAP_CONTAINER)
	log.Println(err)

	// e1 := world.CreateEntity()
	// world.EmplaceComponent(e1, &Position{X: 100, Y: 1000})

	for i := 0; i < 100; i++ {
		e2 := world.CreateEntity()
		world.EmplaceComponent(e2, &Position{X: 200, Y: 2000})
		world.EmplaceComponent(e2, &Velocity{D: 2222})
	}

	// scheduler := gecs.NewScheduler(world)
	// scheduler.AddSystem(&MovementSystem{})

	// scheduler.Init()
	// scheduler.Step()
	// scheduler.Step()
	// scheduler.Step()
	// scheduler.Dispose()
}

type MovementSystem struct {
}

func (s *MovementSystem) Init() gecs.Args {
	log.Println("INIT MOVEMENT")

	return gecs.Args{
		Access:  []gecs.ComponentType{POSITION_COMPONENT, VELOCITY_COMPONENT},
		Exclude: []gecs.ComponentType{},
	}
}

func (s *MovementSystem) Execute(cmd gecs.Command, q gecs.Query) {

	q.ForEach(func(e gecs.Entity) bool {
		// pos := q.GetComponent(e, POSITION_COMPONENT).(*Position)
		// log.Println(pos)
		// vel := q.GetComponent(e, VELOCITY_COMPONENT).(*Velocity)
		// log.Println(vel)

		return false
	})
}

func (s *MovementSystem) Dispose() {
	log.Println("DIPOSE MOVEMENT")
}
