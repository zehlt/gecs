package main

import (
	"fmt"

	"github.com/zehlt/gecs"
)

type MoveSystem struct {
}

func (s *MoveSystem) Init() gecs.Args {
	return gecs.Args{
		Access:   gecs.Access{&Position{}, &Speed{}},
		Exclude:  gecs.Exclude{},
		Resource: gecs.Resource{},
	}
}

func (s *MoveSystem) Exec(cmd gecs.Controller, q gecs.Query) {
	q.Entities(func(e gecs.Entity, comps []interface{}) {
		pos, _ := comps[0].(*Position)
		spd, _ := comps[1].(*Speed)

		pos.X += int(spd.V)
	})
}

// ---

type EnemyBarkSystem struct {
}

func (s *EnemyBarkSystem) Init() gecs.Args {
	return gecs.Args{
		Access:   gecs.Access{&Life{}},
		Exclude:  gecs.Exclude{},
		Resource: gecs.Resource{},
	}
}

func (s *EnemyBarkSystem) Exec(cmd gecs.Controller, q gecs.Query) {
	q.Entities(func(e gecs.Entity, comps []interface{}) {
		life := comps[0].(*Life)
		life.HP -= 10
		fmt.Println("Life Left:", life)
	})
}

// ---

type KillPlayerSystem struct {
}

func (s *KillPlayerSystem) Init() gecs.Args {

	return gecs.Args{
		Access:   gecs.Access{&Life{}, &Player{}},
		Exclude:  gecs.Exclude{},
		Resource: gecs.Resource{},
	}
}

func (s *KillPlayerSystem) Exec(cmd gecs.Controller, q gecs.Query) {
	q.Entities(func(e gecs.Entity, comps []interface{}) {
		player_life := comps[0].(*Life)

		if player_life.HP < 0 {
			cmd.DestroyEntity(e)
		}

		fmt.Println("PLAYER ALIVE")
	})
}
