package system

import (
	"github.com/zehlt/gecs"
	"github.com/zehlt/gecs/command"
	"github.com/zehlt/gecs/query"
)

type QSystem struct {
	q query.Query
	s System
}

type Scheduler struct {
	ctl     command.Controller
	qm      query.QueryMaker
	w       gecs.World
	systems []QSystem
}

func NewScheduler(w gecs.World) Scheduler {
	return Scheduler{
		w:       w,
		qm:      query.NewQueryMaker(w),
		systems: make([]QSystem, 0),
		ctl:     command.NewController(w),
	}
}

func (s *Scheduler) AddSystem(system System) {
	query := system.Init(s.qm)

	s.systems = append(s.systems, QSystem{q: query, s: system})
}

func (s Scheduler) Run() {
	for _, system := range s.systems {
		system.s.Exec(s.ctl, system.q)
		s.ctl.Execute()
	}
}
