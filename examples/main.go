package main

import (
	"fmt"

	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/component"
	"github.com/zehlt/gecs/snapshot"
	"github.com/zehlt/gecs/system"
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

func MyPrintln(x interface{}) {
	fmt.Printf("t: %T, v: %v\n", x, x)
}

func main() {

	// // Testing
	world := gecs.DefaultWorld()
	world.RegisterComponent(&Position{}, component.SPARSE_ARRAY_CONTAINER)
	world.RegisterComponent(&Speed{}, component.HASHMAP_CONTAINER)
	world.RegisterComponent(&Life{}, component.HASHMAP_CONTAINER)
	world.RegisterComponent(&Enemy{}, component.NULL_CONTAINER)
	world.RegisterComponent(&Player{}, component.NULL_CONTAINER)

	e1, err := world.CreateEntity()
	if err != nil {
		panic(err)
	}

	e2, err := world.CreateEntity()
	if err != nil {
		panic(err)
	}

	e3, err := world.CreateEntity()
	if err != nil {
		panic(err)
	}

	world.AddComponent(e1, &Position{X: 10, Y: 10})
	world.AddComponent(e1, &Speed{V: 100, A: 1000})
	world.AddComponent(e1, &Life{HP: 100})
	world.AddComponent(e1, &Player{})

	world.AddComponent(e2, &Enemy{})
	world.AddComponent(e2, &Position{X: 20, Y: 20})
	world.AddComponent(e2, &Life{HP: 200})
	world.AddComponent(e2, &Speed{V: 200, A: 2000})

	world.AddComponent(e3, &Enemy{})
	world.AddComponent(e3, &Position{X: 30, Y: 30})
	world.AddComponent(e3, &Life{HP: 300})
	world.AddComponent(e3, &Speed{V: 300, A: 3000})

	// SCHEDULER EXAMPLE
	// sc := system.NewScheduler(world)
	// sc.AddSystem(&MoveSystem{})
	// sc.AddSystem(&EnemyBarkSystem{})
	// sc.AddSystem(&KillPlayerSystem{})

	fmt.Println("--- SERIALIZE ---")
	serial := snapshot.NewSerializer()
	serial.Register(&Position{})
	serial.Register(&Speed{})
	serial.Register(&Life{})

	bytes, err := serial.Serialize(world)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)

	fmt.Println("--- DESERIALIZE ---")
	w2, err := serial.Deserialize(bytes)
	if err != nil {
		panic(err)
	}

	ents := w2.GetAllEntities()
	for _, e := range ents {
		cs, _ := w2.GetAllComponentsFromEntity(e)
		fmt.Println("c", cs)
	}

	sc2 := system.NewScheduler(w2)
	sc2.AddSystem(&MoveSystem{})
	sc2.AddSystem(&EnemyBarkSystem{})
	sc2.AddSystem(&KillPlayerSystem{})

	for i := 0; i < 10; i++ {
		sc2.Run()
	}
}
