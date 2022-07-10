package main

import (
	"fmt"

	"github.com/zehlt/gecs"
)

type MoveSystem struct {
}

func (s *MoveSystem) Init(qm gecs.QueryMaker) gecs.Query {
	return qm.Create(gecs.Resource{}, gecs.Access{&Position{}, &Speed{}}, gecs.Exclude{})
}

func (s *MoveSystem) Exec(cmd gecs.Controller, q gecs.Query) {
	q.Entities(func(e gecs.Entity) {
		// pos, _ := q.GetComponent(e, 0).(*Position)
		// spd, _ := q.GetComponent(e, 1).(*Speed)

		// pos.X += int(spd.V)
	})
}

// ---

type EnemyBarkSystem struct {
}

func (s *EnemyBarkSystem) Init(qm gecs.QueryMaker) gecs.Query {
	return qm.Create(gecs.Resource{}, gecs.Access{&Life{}}, gecs.Exclude{})
}

func (s *EnemyBarkSystem) Exec(cmd gecs.Controller, q gecs.Query) {

	q.Entities(func(e gecs.Entity) {
		life := q.GetComponent(e, 0).(*Life)
		life.HP -= 10
		fmt.Println("Life Left:", life)
	})
}

// ---

type KillPlayerSystem struct {
}

func (s *KillPlayerSystem) Init(qm gecs.QueryMaker) gecs.Query {
	return qm.Create(gecs.Resource{}, gecs.Access{&Life{}, &Player{}}, gecs.Exclude{})
}

func (s *KillPlayerSystem) Exec(cmd gecs.Controller, q gecs.Query) {

	q.Entities(func(e gecs.Entity) {
		player_life := q.GetComponent(e, 0).(*Life)

		if player_life.HP < 0 {
			cmd.DestroyEntity(e)
		}

		fmt.Println("PLAYER ALIVE")
	})
}
