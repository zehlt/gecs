package registry

type Registry interface {
	CreateEntity()
	DestroyEntity()
	HasEntity()

	AddComponent()
	RemoveComponent()
	HasComponent()
}
