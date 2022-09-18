package gecs

type World struct {
	ship     *ship
	registry *registry
}

func NewWorld() *World {
	return &World{
		ship:     newShip(),
		registry: newRegistry(),
	}
}

func (w *World) CreateEntity() Entity {
	return w.registry.CreateEntity()
}

func (w *World) DestroyEntity(e Entity) error {
	if err := w.registry.DestroyEntity(e); err != nil {
		return err
	}

	w.ship.RemoveAllComponents(e)
	return nil
}

func (w *World) RegisterComponent(t ComponentType, c ContainerType) error {
	if err := w.ship.RegisterComponent(t, c); err != nil {
		return err
	}

	w.registry.RegisterComponent(t)
	return nil
}

func (w *World) EmplaceComponent(e Entity, c Component) error {
	if err := w.registry.EmplaceComponent(e, c.GetType()); err != nil {
		return err
	}

	w.ship.EmplaceComponent(e, c)
	return nil
}

func (w *World) RemoveComponent(e Entity, t ComponentType) error {
	if err := w.registry.RemoveComponent(e, t); err != nil {
		return err
	}

	w.ship.RemoveComponent(e, t)
	return nil
}

func (w *World) GetComponent(e Entity, t ComponentType) (Component, error) {
	if ok := w.registry.HasComponent(e, t); !ok {
		return nil, ErrContainerEntityDoesNotHaveComponent
	}

	return w.ship.GetComponent(e, t), nil
}

func (w *World) HasComponent(e Entity, t ComponentType) bool {
	return w.registry.HasComponent(e, t)
}
