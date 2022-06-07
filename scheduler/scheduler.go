package scheduler

import "github.com/zehlt/gecs"

type Scheduler struct {
	w gecs.World
}

// func (s *Scheduler) AddSystem(system System) {
// 	// qm := system.Init()

// 	// s.qsystems = append(s.qsystems, system)
// 	system.Init()
// }

// func (s Scheduler) Run() {
// 	// for _, system := range s.systems {
// 	// 	system.Exec(command.Controller{}, nil)
// 	// }
// }

// func (s *Scheduler) Build(w gecs.World) {
// 	s.w = w
// }
