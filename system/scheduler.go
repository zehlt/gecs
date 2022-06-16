package system

import (
	"reflect"

	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/command"
	"github.com/zehlt/gecs/query"
)

type Scheduler interface {
	AddSystem(system System)
	Run(w gecs.World)

	AddReceiver(re Receiver)
	Signal(signal interface{}, w gecs.World)
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

func (s *scheduler) Signal(signal interface{}, w gecs.World) {
	sign, ok := s.signals[reflect.TypeOf(signal)]
	ctl := command.NewController(w)

	if ok {
		sign.Exec(ctl, signal)
	}
}

// TODO: should optimize that
func (s *scheduler) Run(w gecs.World) {
	qm := query.NewQueryMaker(w)
	ctl := command.NewController(w)

	for _, sys := range s.systems {
		query := sys.Init(qm)
		sys.Exec(ctl, query)
		ctl.Execute()
	}
}
