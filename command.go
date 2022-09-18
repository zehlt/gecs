package gecs

import "log"

type worldCommand interface {
	execute(w *World, systems []*systemCached)
}

type Command struct {
	creators []worldCommand
}

func (cmd *Command) CreateEntity() *entityCreator {
	creator := &entityCreator{}
	cmd.creators = append(cmd.creators, creator)
	return creator
}

func (cmd *Command) DestroyEntity(e Entity) {
	cmd.creators = append(cmd.creators, &entityDestructor{
		entityToDestroy: e,
	})
}

func (cmd *Command) EmplaceComponent(e Entity, c Component) {
	cmd.creators = append(cmd.creators, &componentEmplacer{
		e: e,
		c: c,
	})
}

func (cmd *Command) RemoveComponent(e Entity, t ComponentType) {
	cmd.creators = append(cmd.creators, &componentRemover{
		e: e,
		t: t,
	})
}

func (cmd *Command) execute(w *World, systems []*systemCached) {
	for _, creator := range cmd.creators {
		creator.execute(w, systems)
	}
}

type entityDestructor struct {
	entityToDestroy Entity
}

func (ec *entityDestructor) execute(w *World, systems []*systemCached) {
	w.DestroyEntity(ec.entityToDestroy)

	for _, system := range systems {
		_, ok := system.entities[ec.entityToDestroy]
		if ok {
			delete(system.entities, ec.entityToDestroy)
		}
	}
}

type entityCreator struct {
	components []Component
}

func (ec *entityCreator) EmplaceComponent(c Component) *entityCreator {
	ec.components = append(ec.components, c)
	return ec
}

func (ec *entityCreator) execute(w *World, systems []*systemCached) {
	e := w.CreateEntity()

	for _, component := range ec.components {
		if err := w.EmplaceComponent(e, component); err != nil {
			log.Fatalln("error trying to execute the entity creator cmd", err)
		}
	}

	sign := w.registry.getSignature(e)
	for _, system := range systems {
		if matchSignature(sign, system.rw, system.exclude) {
			system.entities[e] = nil
		}
	}
}

type componentEmplacer struct {
	e Entity
	c Component
}

func (ec *componentEmplacer) execute(w *World, systems []*systemCached) {
	w.EmplaceComponent(ec.e, ec.c)
}

type componentRemover struct {
	e Entity
	t ComponentType
}

func (ec *componentRemover) execute(w *World, systems []*systemCached) {
	w.RemoveComponent(ec.e, ec.t)
}
