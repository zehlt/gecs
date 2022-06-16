package main

import (
	"fmt"

	"github.com/zehlt/gecs/command"
)

type MovePlayerSignal struct {
	X int
	Y int
}

type NothingSignal struct {
	Nothing string
}

type HealUserSignal struct {
	Amount float64
}

//

type InputReceiver struct {
}

func (s *InputReceiver) Init() interface{} {
	return MovePlayerSignal{}
}

func (s *InputReceiver) Exec(cmd command.Controller, d interface{}) {
	fmt.Println("SIGNAL CALLED!!", d)
}

//

type HealUserReceiver struct {
}

func (s *HealUserReceiver) Init() interface{} {
	return HealUserSignal{}
}

func (s *HealUserReceiver) Exec(cmd command.Controller, d interface{}) {
	fmt.Println("HEAL CALLEd!!", d)
}
