package gecs

type World interface {
	// entity
	CreateEntity() (Entity, error)
	DestroyEntity(Entity) error
	EntityExists(Entity) bool
	GetAllEntities() []Entity

	// component
	RegisterComponent(interface{}, ContainerType) error
	AddComponent(Entity, interface{}) error
	EmplaceComponent(Entity, interface{}) error
	RemoveComponent(Entity, interface{}) error
	GetComponent(Entity, interface{}) (interface{}, error)
	HasComponent(Entity, interface{}) bool
	GetAllComponentsFromEntity(Entity) ([]interface{}, error)

	// resource
	AddResource(interface{}) error
	GetResource(interface{}) (interface{}, error)
	HasResource(interface{}) bool

	// query
	Query(args Args) Query
}

type world struct {
	locker   locker
	arena    arena
	store    store
	registry registry
}

func DefaultWorld() World {
	return &world{
		locker:   newLocker(),
		arena:    newArena(),
		store:    newStore(),
		registry: newRegistry(),
	}
}

func (w *world) CreateEntity() (Entity, error) {
	e, err := w.arena.Create()
	if err != nil {
		return Entity{}, err
	}

	err = w.registry.CreateEntitySignature(e)
	if err != nil {
		return Entity{}, err
	}

	return e, nil
}

func (w *world) GetAllEntities() []Entity {
	return w.arena.GetAll()
}

func (w *world) DestroyEntity(e Entity) error {
	err := w.arena.Destroy(e)
	if err != nil {
		// TODO: layer more error
		return err
	}

	w.registry.DestroyEntitySignature(e)

	return w.store.RemoveAll(e)
}

func (w *world) EntityExists(e Entity) bool {
	return w.arena.Exists(e)
}

func (w *world) RegisterComponent(c interface{}, t ContainerType) error {
	id := w.registry.GetComponentId(c)
	return w.store.Register(id, t)
}

func (w *world) AddComponent(e Entity, c interface{}) error {
	if !w.arena.Exists(e) {
		return ErrEntityDoesNotExist
	}

	componenId := w.registry.GetComponentId(c)
	err := w.registry.AddComponent(e, componenId)
	if err != nil {
		// TODO: add layer of info in error
		return err
	}

	return w.store.Add(e, componenId, c)
}

func (w *world) EmplaceComponent(e Entity, c interface{}) error {
	if !w.arena.Exists(e) {
		return ErrEntityDoesNotExist
	}

	componenId := w.registry.GetComponentId(c)

	err := w.registry.AddComponent(e, componenId)
	if err != nil {
		// TODO: add layer of info in error
		return err
	}

	w.store.Emplace(e, componenId, c)
	return nil
}

func (w *world) RemoveComponent(e Entity, c interface{}) error {
	if !w.arena.Exists(e) {
		return ErrEntityDoesNotExist
	}

	id := w.registry.GetComponentId(c)
	if !w.registry.HasComponent(e, id) {
		return ErrEntityDoesNotHaveComponent
	}
	w.registry.RemoveComponent(e, id)

	return w.store.Remove(e, id)
}

func (w *world) GetComponent(e Entity, c interface{}) (interface{}, error) {
	if !w.arena.Exists(e) {
		return nil, ErrEntityDoesNotExist
	}

	id := w.registry.GetComponentId(c)
	if !w.registry.HasComponent(e, id) {
		return nil, ErrEntityDoesNotHaveComponent
	}

	return w.store.Get(e, id)
}

func (w *world) GetAllComponentsFromEntity(e Entity) ([]interface{}, error) {
	if !w.arena.Exists(e) {
		return nil, ErrEntityDoesNotExist
	}

	return w.store.GetAll(e), nil
}

func (w *world) HasComponent(e Entity, c interface{}) bool {
	if !w.arena.Exists(e) {
		return false
	}

	id := w.registry.GetComponentId(c)
	return w.registry.HasComponent(e, id)
}

func (w *world) AddResource(c interface{}) error {
	return w.locker.Add(c)
}

func (w *world) GetResource(t interface{}) (interface{}, error) {
	return w.locker.Get(t)
}

func (w *world) HasResource(t interface{}) bool {
	return w.locker.Has(t)
}

func (w *world) Query(args Args) Query {
	access_sign := w.registry.GetSignatureFromTypes(args.Access)
	// exclude_sign := w.registry.GetSignatureFromTypes(e)

	entities := w.registry.FindMatchingEntities(access_sign)
	components := make([][]interface{}, len(entities))
	resources := make([]interface{}, len(args.Resource))

	componentIds := make([]ComponentId, len(args.Access))
	for i, t := range args.Access {
		componentIds[i] = w.registry.GetComponentId(t)
	}

	for i, e := range entities {
		comps := make([]interface{}, len(componentIds))

		for y, cid := range componentIds {
			c, err := w.store.Get(e, cid)
			if err != nil {
				panic(err)
			}

			comps[y] = c
		}

		components[i] = comps
	}

	for i, t := range args.Resource {
		co, err := w.locker.Get(t)
		if err != nil {
			panic(err)
		}
		resources[i] = co
	}

	return &query{
		entities:   entities,
		components: components,
		resources:  resources,
	}
}
