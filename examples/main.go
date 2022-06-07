package main

import (
	"fmt"

	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/component"
)

type Position struct {
	x int
	y int
}

type Speed struct {
	v float64
	a float64
}

func main() {
	world := gecs.DefaultWorld()
	e1, err := world.CreateEntity()
	if err != nil {
		panic(err)
	}
	// e2, err := world.CreateEntity()
	// if err != nil {
	// 	panic(err)
	// }
	e3, err := world.CreateEntity()
	if err != nil {
		panic(err)
	}

	world.RegisterComponent(Position{}, component.SPARSE_ARRAY_CONTAINER)
	world.RegisterComponent(Speed{}, component.HASHMAP_CONTAINER)

	world.AddComponent(e1, Position{x: 10, y: 25})
	world.AddComponent(e3, Position{x: 10, y: 25})
	world.AddComponent(e3, Speed{v: 999, a: 111111})

	fmt.Println(world.HasComponent(e1, Speed{}))
	fmt.Println(world.HasComponent(e1, Position{}))

	fmt.Println(world.GetComponent(e3, Position{}))
	fmt.Println(world.GetComponent(e3, Speed{}))

	world.DestroyEntity(e3)

	fmt.Println(e3)
	fmt.Println(world.GetComponent(e3, Position{}))
	fmt.Println(world.GetComponent(e3, Speed{}))

	// e4, err := world.CreateEntity()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(e4)
	fmt.Println(world.GetComponent(e3, Position{}))
	fmt.Println(world.GetComponent(e3, Speed{}))

	// e1, err := world.CreateEntity()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(e1)

	// fmt.Println(world.EntityExists(e1))
}

// import (
// 	"fmt"

// 	"github.com/zehlt/gecs/registry"
// )

// type Player struct {
// }

// func Move() {
// 	fmt.Println("Move")
// }

// func Speak() {
// 	fmt.Println("Speak")
// }

// func Eat() {
// 	fmt.Println("Eat")
// }

// func main() {
// 	fmt.Println("--- GECS: Sandbox ---")

// 	// store := component.NewStore()
// 	// store.Register(1, component.SPARSE_ARRAY_CONTAINER)
// 	// store.Register(2, component.NULL_CONTAINER)

// 	r := registry.NewRegistry()
// 	fmt.Println(r.GetComponentId(Position{}))
// }
