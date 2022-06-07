package main

import (
	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/component"
	"github.com/zehlt/gecs/scheduler"
)

type Position struct {
	x int
	y int
}

type Speed struct {
	v float64
	a float64
}

type Life struct {
	hp int
}

type Enemy struct {
}

func main() {
	world := gecs.DefaultWorld()
	world.RegisterComponent(Position{}, component.SPARSE_ARRAY_CONTAINER)
	world.RegisterComponent(Speed{}, component.HASHMAP_CONTAINER)
	world.RegisterComponent(Life{}, component.HASHMAP_CONTAINER)
	world.RegisterComponent(Enemy{}, component.NULL_CONTAINER)

	e1, err := world.CreateEntity()
	if err != nil {
		panic(err)
	}

	e2, err := world.CreateEntity()
	if err != nil {
		panic(err)
	}

	world.AddComponent(e1, Position{x: 10, y: 10})
	world.AddComponent(e1, Speed{v: 100, a: 1000})
	world.AddComponent(e1, Life{hp: 100})
	world.AddComponent(e1, Enemy{})

	world.AddComponent(e2, Enemy{})
	world.AddComponent(e2, Position{x: 20, y: 20})
	world.AddComponent(e2, Life{hp: 200})
	world.AddComponent(e2, Speed{v: 200, a: 2000})

	// SCHEDULER EXAMPLE
	sc := scheduler.NewScheduler(world)
	sc.AddSystem(&MoveSystem{})
	sc.AddSystem(&EnemyBarkSystem{})
	sc.Run()
}
