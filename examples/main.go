package main

import (
	"fmt"

	"github.com/zehlt/gecs"
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

type TryInterface interface {
	Try() bool
}

type Renderer struct {
	x int
}

func main() {
	// // Testing
	world := gecs.DefaultWorld()
	world.RegisterComponent(&Position{}, gecs.SPARSE_ARRAY_CONTAINER)
	world.RegisterComponent(&Speed{}, gecs.HASHMAP_CONTAINER)
	world.RegisterComponent(&Life{}, gecs.HASHMAP_CONTAINER)
	world.RegisterComponent(&Enemy{}, gecs.NULL_CONTAINER)
	world.RegisterComponent(&Player{}, gecs.NULL_CONTAINER)

	err := world.AddResource(Renderer{x: 12})
	if err != nil {
		panic(err)
	}

	re, err := world.GetResource(Renderer{})
	if err != nil {
		panic(err)
	}
	fmt.Println("ress", re)

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

	fmt.Println("--- SERIALIZE ---")
	serial := gecs.NewSerializer()
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
	fmt.Println(w2)

	ents := w2.GetAllEntities()
	for _, e := range ents {
		cs, _ := w2.GetAllComponentsFromEntity(e)
		fmt.Println("c", cs)
	}

	sc2 := gecs.NewScheduler()
	sc2.AddSystem(&MoveSystem{})
	sc2.AddSystem(&EnemyBarkSystem{})
	sc2.AddSystem(&KillPlayerSystem{})

	sc2.AddReceiver(&InputReceiver{})
	sc2.AddReceiver(&HealUserReceiver{})

	for i := 0; i < 10; i++ {
		sc2.Run(w2)
	}

	sc2.Signal(MovePlayerSignal{}, w2)
	sc2.Signal(NothingSignal{}, w2)
	sc2.Signal(MovePlayerSignal{}, w2)
	sc2.Signal(MovePlayerSignal{}, w2)
	sc2.Signal(HealUserSignal{}, w2)
	sc2.Signal(MovePlayerSignal{}, w2)
}
