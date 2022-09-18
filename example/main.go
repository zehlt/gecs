package main

import (
	"log"

	"github.com/zehlt/gecs"
)

const (
	PLAYER_COMPONENT gecs.ComponentType = iota
	ENEMY_COMPONENT
	VELOCITY_COMPONENT
	POSITION_COMPONENT
)

type Enemy struct {
}

func (p Enemy) GetType() gecs.ComponentType {
	return ENEMY_COMPONENT
}

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
	world.RegisterComponent(PLAYER_COMPONENT, gecs.TAG_CONTAINER)
	world.RegisterComponent(ENEMY_COMPONENT, gecs.TAG_CONTAINER)
	world.RegisterComponent(POSITION_COMPONENT, gecs.HASHMAP_CONTAINER)
	world.RegisterComponent(VELOCITY_COMPONENT, gecs.HASHMAP_CONTAINER)

	scheduler := gecs.NewScheduler(world)
	scheduler.AddStartupSystem(&CreateWorldStartup{})
	scheduler.AddSystem(&MovementSystem{})
	scheduler.AddSystem(&EnemySystem{})
	scheduler.Build()

	scheduler.Step()
	scheduler.Step()
	scheduler.Step()
	scheduler.Dispose()
}

type MovementSystem struct {
}

func (s *MovementSystem) Init() gecs.Args {
	return gecs.Args{
		Read:    []gecs.ComponentType{POSITION_COMPONENT},
		Write:   []gecs.ComponentType{},
		Exclude: []gecs.ComponentType{},
	}
}

func (s *MovementSystem) Execute(cmd *gecs.Command, q gecs.Query) {
	q.ForEach(func(e gecs.Entity) bool {
		pos := q.GetComponent(e, POSITION_COMPONENT).(*Position)
		log.Println("POS: ", pos)

		return false
	})
}

func (s *MovementSystem) Dispose() {
}

type PlayerSystem struct {
}

func (s *PlayerSystem) Init() gecs.Args {
	return gecs.Args{
		Read:    []gecs.ComponentType{PLAYER_COMPONENT},
		Write:   []gecs.ComponentType{},
		Exclude: []gecs.ComponentType{},
	}
}

func (s *PlayerSystem) Execute(cmd *gecs.Command, q gecs.Query) {
	q.ForEach(func(e gecs.Entity) bool {

		return false
	})
}

type EnemySystem struct {
	once bool
}

func (s *EnemySystem) Init() gecs.Args {
	return gecs.Args{
		Read:    []gecs.ComponentType{ENEMY_COMPONENT},
		Write:   []gecs.ComponentType{},
		Exclude: []gecs.ComponentType{},
	}
}

func (s *EnemySystem) Execute(cmd *gecs.Command, q gecs.Query) {
	q.ForEach(func(e gecs.Entity) bool {
		if !s.once {
			cmd.DestroyEntity(e)
			s.once = true
			log.Println("destroy enemy called")
		} else {
			log.Println("ENEMY e:", e)
		}
		return false
	})
}

func (s *EnemySystem) Dispose() {
}

type CreateWorldStartup struct {
}

func (s *CreateWorldStartup) Execute(cmd *gecs.Command) {
	enemy := cmd.CreateEntity()
	enemy.EmplaceComponent(Enemy{})
	enemy.EmplaceComponent(&Position{X: 999, Y: 9999})

	player := cmd.CreateEntity()
	player.EmplaceComponent(Player{})
	player.EmplaceComponent(&Position{X: 111, Y: 1111})
}
