package gecs

type EntityCellType int

const (
	EMPTY_CELL EntityCellType = iota
	END_CELL
	OCCUPIED_CELL
)

type EntityCell struct {
	t          EntityCellType
	generation uint64
	next       int
}
