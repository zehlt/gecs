package gecs

type System interface {
	Init() Args
	Exec(cmd Controller, q Query)
}

type Service interface {
	Init(w World, d Dispatch[SignalType, interface{}])
}

type Scheduler interface {
	AddService(service Service)
	AddSystem(system System)
	Run()

	Emit(t SignalType, data interface{})
}

type SignalType int

type scheduler struct {
	w          World
	systems    []System
	dispatcher Dispatch[SignalType, interface{}]
}

func NewScheduler(w World) Scheduler {
	var dispatcher Dispatch[SignalType, interface{}]
	dispatcher.Init()

	return &scheduler{
		systems:    make([]System, 0),
		w:          w,
		dispatcher: dispatcher,
	}
}

func (s *scheduler) AddSystem(system System) {
	s.systems = append(s.systems, system)
}

func (s *scheduler) AddService(service Service) {
	service.Init(s.w, s.dispatcher)
}

func (s *scheduler) Emit(t SignalType, data interface{}) {
	s.dispatcher.Disp(t, data)
}

// TODO: should optimize that
func (s *scheduler) Run() {
	ctl := newController(s.w, s)

	for _, sys := range s.systems {
		args := sys.Init()
		query := s.w.Query(args)
		sys.Exec(ctl, query)
		ctl.Execute()
	}
}
