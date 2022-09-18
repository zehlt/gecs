package gecs

type Args struct {
	Read    []ComponentType
	Write   []ComponentType
	Exclude []ComponentType
}

type StartupSystem interface {
	Execute(cmd *Command)
}

type System interface {
	Init() Args
	Execute(cmd *Command, q Query)
	Dispose()
}

type systemCached struct {
	system   System
	entities map[Entity]any
	write    signature
	read     signature
	rw       signature
	exclude  signature
}

type Scheduler struct {
	world *World

	startups []StartupSystem
	systems  []*systemCached
}

func NewScheduler(w *World) *Scheduler {
	return &Scheduler{
		world: w,
	}
}

func (sc *Scheduler) AddStartupSystem(st StartupSystem) {
	sc.startups = append(sc.startups, st)
}

func (sc *Scheduler) AddSystem(s System) {
	cached := sc.CacheSystem(s)
	sc.systems = append(sc.systems, cached)
}

func (sc *Scheduler) Build() {
	// should build the tree

	// launch statups systems
	cmd := &Command{}
	for _, startup := range sc.startups {
		startup.Execute(cmd)
	}
	cmd.execute(sc.world, sc.systems)
}

func (sc *Scheduler) Step() {

	cmds := make([]*Command, len(sc.systems))
	for i, system := range sc.systems {
		cmd := &Command{}
		cmds[i] = cmd
		system.system.Execute(cmd, Query{
			world:    sc.world,
			entities: system.entities,
		})
	}

	for _, cmd := range cmds {
		cmd.execute(sc.world, sc.systems)
	}
}

func (sc *Scheduler) CacheSystem(s System) *systemCached {
	// get system args
	args := s.Init()

	// parse and cache
	exclude := sc.world.registry.newSignatureFromTypes(args.Exclude)
	read := sc.world.registry.newSignatureFromTypes(args.Read)
	write := sc.world.registry.newSignatureFromTypes(args.Write)

	rw := read.Clone()
	rw.Or(write)
	entities := sc.world.registry.getMatchingEntities(rw, exclude)

	cached := systemCached{
		write:    write,
		read:     read,
		rw:       rw,
		exclude:  exclude,
		system:   s,
		entities: entities,
	}

	return &cached
}

func (sc *Scheduler) Dispose() {
	// for _, stage := range sc.systems {
	// 	for _, system := range stage.systems {
	// 		system.Dispose()
	// 	}
	// }
}

type Query struct {
	entities map[Entity]any
	world    *World
}

func (q *Query) ForEach(fn func(e Entity) bool) {
	for ent := range q.entities {
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

// s := sc.systems[stage]

// wp := WorkerPool{}
// wp.Start(4)

// cmds := make([]*Command, len(s.systems))
// for i, system := range s.systems {
// 	systemCommand := &Command{}
// 	cmds[i] = systemCommand
// 	wp.Send(func() {
// 		system.Execute(systemCommand, Query{entities: system.entities, world: sc.world})
// 	})
// }
// wp.Stop()

// // updating the state of the cached systems entities
// for _, cmd := range cmds {
// 	cmd.execute(sc.world)
// }
