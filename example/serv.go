package main

import (
	"fmt"

	"github.com/zehlt/gecs"
)

type UserService struct {
	w gecs.World
}

type UserHealingSignal struct {
	Heal int
}

func (s *UserService) Init(w gecs.World, d gecs.Dispatch[gecs.SignalType, interface{}]) {
	s.w = w
	d.Register(1, s.OnUserHealingSignal)
}

func (s *UserService) OnUserHealingSignal(data interface{}) {
	sig := data.(UserHealingSignal)
	fmt.Println("SINGAL 1 RECEIVED", sig)
}
