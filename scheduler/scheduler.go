package scheduler

type Scheduler struct {
	systems []System
}

func (s *Scheduler) AddSystem(system System) {
	s.systems = append(s.systems, system)
}

func (s Scheduler) Run() {
	for _, system := range s.systems {
		system.Exec()
	}
}
