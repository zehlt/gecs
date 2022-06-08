package main

import (
	"fmt"

	"github.com/zehlt/gecs/command"
	"github.com/zehlt/gecs/entity"
	"github.com/zehlt/gecs/query"
)

type MoveSystem struct {
}

func (s *MoveSystem) Init(qm query.QueryMaker) query.Query {
	return qm.Create(query.Access{&Position{}, &Speed{}}, query.Exclude{})
}

func (s *MoveSystem) Exec(cmd command.Controller, q query.Query) {
	q.Entities(func(e entity.Entity) {
		pos, _ := q.GetComponent(e, 0).(*Position)
		spd, _ := q.GetComponent(e, 1).(*Speed)

		pos.x += int(spd.v)
	})
}

// ---

type EnemyBarkSystem struct {
}

func (s *EnemyBarkSystem) Init(qm query.QueryMaker) query.Query {
	return qm.Create(query.Access{&Life{}}, query.Exclude{})
}

func (s *EnemyBarkSystem) Exec(cmd command.Controller, q query.Query) {

	q.Entities(func(e entity.Entity) {
		life := q.GetComponent(e, 0).(*Life)
		life.hp -= 10
		fmt.Println("Life Left:", life)
	})
}

// ---

type KillPlayerSystem struct {
}

func (s *KillPlayerSystem) Init(qm query.QueryMaker) query.Query {
	return qm.Create(query.Access{&Life{}, &Player{}}, query.Exclude{})
}

func (s *KillPlayerSystem) Exec(cmd command.Controller, q query.Query) {

	q.Entities(func(e entity.Entity) {
		player_life := q.GetComponent(e, 0).(*Life)

		if player_life.hp < 0 {
			cmd.DestroyEntity(e)
		}

		fmt.Println("PLAYER ALIVE")
	})
}
