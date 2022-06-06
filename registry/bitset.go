package registry

type Bitset interface {
	GetBit(pos int) bool
	SetBit(pos int, val bool)
}
