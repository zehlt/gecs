package gecs

import (
	"reflect"
)

type System interface {
	Init(qm QueryMaker) Query
	Exec(cmd Controller, q Query)
}

type Receiver interface {
	Init() interface{}
	Exec(cmd Controller, signal interface{})
}

type Scheduler interface {
	AddSystem(system System)
	Run(w World)

	AddReceiver(re Receiver)
	Signal(signal interface{}, w World)
}

type scheduler struct {
	systems []System
	signals map[reflect.Type]Receiver
}

func NewScheduler() Scheduler {
	return &scheduler{
		systems: make([]System, 0),
		signals: make(map[reflect.Type]Receiver),
	}
}

func (s *scheduler) AddSystem(system System) {
	s.systems = append(s.systems, system)
}

func (s *scheduler) AddReceiver(signal Receiver) {
	c := signal.Init()
	t := reflect.TypeOf(c)
	s.signals[t] = signal
}

func (s *scheduler) Signal(signal interface{}, w World) {
	sign, ok := s.signals[reflect.TypeOf(signal)]
	ctl := newController(w)

	if ok {
		sign.Exec(ctl, signal)
	}
}

// TODO: should optimize that
func (s *scheduler) Run(w World) {
	qm := NewQueryMaker(w)
	ctl := newController(w)

	for _, sys := range s.systems {
		query := sys.Init(qm)
		sys.Exec(ctl, query)
		ctl.Execute()
	}
}
