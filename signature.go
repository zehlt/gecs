package gecs

type Signature interface {
	AddComponent(id ComponentId)
	RemoveComponent(id ComponentId)
	HasComponent(id ComponentId) bool
	Contains(Signature) bool
	GetData() interface{}
	String() string
}

type signature struct {
	bitset *Bitset
}

func NewSignature() Signature {
	return &signature{
		bitset: NewBitset(),
	}
}

func (s *signature) AddComponent(id ComponentId) {
	s.bitset.Set(int(id), true)
}

func (s *signature) RemoveComponent(id ComponentId) {
	s.bitset.Set(int(id), false)
}

func (s *signature) HasComponent(id ComponentId) bool {
	return s.bitset.Get(int(id))
}

func (s *signature) Contains(other Signature) bool {
	other_data := other.GetData()
	other_bitset := other_data.(*Bitset)

	return s.bitset.Contains(other_bitset)
}

func (s *signature) GetData() interface{} {
	return s.bitset
}

func (s *signature) String() string {
	return s.bitset.String()
}
