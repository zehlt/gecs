package system

import (
	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/command"
	"github.com/zehlt/gecs/query"
)

type Scheduler interface {
	AddSystem(system System)
	Run(w gecs.World)
}

type scheduler struct {
	systems []System
}

func NewScheduler() Scheduler {
	return &scheduler{
		systems: make([]System, 0),
	}
}

func (s *scheduler) AddSystem(system System) {
	s.systems = append(s.systems, system)
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
