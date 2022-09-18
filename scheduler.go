package gecs

type Access []interface{}
type Exclude []interface{}

type Args struct {
	Access  []ComponentType
	Exclude []ComponentType
}

type System interface {
	Init() Args
	Execute(cmd Command, q Query)
	Dispose()
}

type stage struct {
	systems []*systemCached
}

func (s *stage) AddSystem(system System) {
	s.systems = append(s.systems, &systemCached{
		System: system,
	})
}

type Scheduler struct {
	world *World

	systems map[string]*stage
}

type systemCached struct {
	System
	access   signature
	exclude  signature
	entities []Entity
}

func NewScheduler(w *World) *Scheduler {
	return &Scheduler{
		world:   w,
		systems: make(map[string]*stage),
	}
}

func (sc *Scheduler) AddSystem(st string, s System) {
	sta, ok := sc.systems[st]
	if !ok {
		sta = &stage{}
		sc.systems[st] = sta
	}

	sta.AddSystem(s)
}

func (sc *Scheduler) Init() {
	for _, stage := range sc.systems {
		for _, system := range stage.systems {
			// cache system signature
			args := system.Init()

			access := sc.world.registry.newSignatureFromTypes(args.Access)
			exclude := sc.world.registry.newSignatureFromTypes(args.Exclude)

			// find all matching entities
			system.entities = sc.world.registry.getMatchingEntities(access, exclude)
			system.access = access
			system.exclude = exclude
		}
	}
}

func (sc *Scheduler) Step(stage string) {
	s := sc.systems[stage]

	wp := WorkerPool{}
	wp.Start(4)

	for _, system := range s.systems {
		wp.Send(func() {
			system.Execute(Command{}, Query{entities: system.entities, world: sc.world})
		})
	}

	wp.Stop()
}

func (sc *Scheduler) Dispose() {
	for _, stage := range sc.systems {
		for _, system := range stage.systems {
			system.Dispose()
		}
	}
}

type Command struct {
}

type Query struct {
	entities []Entity
	world    *World
}

func (q *Query) ForEach(fn func(e Entity) bool) {
	for _, ent := range q.entities {
		stop := fn(ent)
		if stop {
			break
		}
	}
}

func (q *Query) GetComponent(e Entity, t ComponentType) Component {
	c, _ := q.world.GetComponent(e, t)
	return c
}
