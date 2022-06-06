package main

import (
	"fmt"

	"github.com/zehlt/gecs/registry"
)

type Position struct {
	x int
	y int
}

type Movement struct {
	vel int
	acc int
}

func main() {
	fmt.Println("--- GECS: Sandbox ---")

	registry := registry.NewSparceRegistry()

	e1, err := registry.CreateEntity()
	if err != nil {
		panic(err)
	}

	registry.AddComponent(e1, Position{x: 10, y: 125})

	fmt.Println(registry.EntityExists(e1))
	fmt.Println(registry.HasComponent(e1, Position{}))
	fmt.Println(registry.GetComponent(e1, Position{}))
}
