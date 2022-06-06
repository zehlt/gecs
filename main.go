package main

import (
	"fmt"

	"github.com/zehlt/gecs/registry"
)

type Position struct {
	x int
	y int
}

type Speed struct {
	v float64
	a float64
}

type Player struct {
}

func Move() {
	fmt.Println("Move")
}

func Speak() {
	fmt.Println("Speak")
}

func Eat() {
	fmt.Println("Eat")
}

func main() {
	fmt.Println("--- GECS: Sandbox ---")

	// store := component.NewStore()
	// store.Register(1, component.SPARSE_ARRAY_CONTAINER)
	// store.Register(2, component.NULL_CONTAINER)

	r := registry.NewRegistry()
	fmt.Println(r.GetComponentId(Position{}))
	fmt.Println(r.GetComponentId(Speed{}))
	fmt.Println(r.GetComponentId(Speed{}))
	fmt.Println(r.GetComponentId(Position{}))
	fmt.Println(r.GetComponentId(Speed{}))
	fmt.Println(r.GetComponentId(Speed{}))
	fmt.Println(r.GetComponentId(Speed{}))
	fmt.Println(r.GetComponentId(Position{}))
	fmt.Println(r.GetComponentId(Speed{}))
	fmt.Println(r.GetComponentId(Speed{}))
	fmt.Println(r.GetComponentId(15))
}
