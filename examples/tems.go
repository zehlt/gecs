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
	return qm.Create(query.Access{Position{}, Speed{}}, query.Exclude{})
}

func (s *MoveSystem) Exec(cmd command.Controller, q query.Query) {
	q.Entities(func(e entity.Entity) {
		pos, _ := q.GetComponent(e, 0).(Position)
		spd, _ := q.GetComponent(e, 1).(Speed)

		fmt.Println(pos.x)
		pos.x += int(spd.v)
		fmt.Println(pos.x)

		fmt.Println("MOVINGGG !! compo:", pos, spd)
	})
}

// ---

type EnemyBarkSystem struct {
}

func (s *EnemyBarkSystem) Init(qm query.QueryMaker) query.Query {
	return qm.Create(query.Access{Life{}, Enemy{}}, query.Exclude{})
}

func (s *EnemyBarkSystem) Exec(cmd command.Controller, q query.Query) {

	q.Entities(func(e entity.Entity) {
		life := q.GetComponent(e, 0).(Life)
		fmt.Println("Life Left:", life)
	})
}
