package snapshot

import (
	"github.com/zehlt/gecs/entity"
)

// Add insert entity
// Add Tag component
// Add reference with types of containers
// Add update world instead of rebuild again
// Make sheduler working on new world without requiery

type Couple struct {
	E          entity.Entity
	Components []interface{}
}

type Reference struct {
}

type Snap struct {
	// Couples []Couple
	Comps [][]interface{}
}
