package main

import (
	"fmt"

	"github.com/zehlt/gecs/query"
	"github.com/zehlt/gecs/registry"
	"github.com/zehlt/gecs/scheduler"
)

type Position struct {
	x int
	y int
}

type Movement struct {
	vel int
	acc int
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

	registry := registry.NewSparceRegistry()
	q := query.Make(&registry, query.Access{Position{}, Movement{}}, query.Exclude{})
	fmt.Println(q)

	s := scheduler.Scheduler{}
	s.AddSystem(scheduler.System{Exec: Move})
	s.AddSystem(scheduler.System{Exec: Speak})
	s.AddSystem(scheduler.System{Exec: Eat})
	s.Run()
}
