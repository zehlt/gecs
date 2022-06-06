package main

import (
	"fmt"

	"github.com/zehlt/gecs/component"
	"github.com/zehlt/gecs/entity"
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

	arena := entity.NewArena()
	arena.Create()
	arena.Create()
	e3, _ := arena.Create()

	store := component.NewSparseStore()
	store.Add(e3, Position{x: 10, y: 115})
	store.Add(e3, Movement{vel: 9, acc: 499})
}
