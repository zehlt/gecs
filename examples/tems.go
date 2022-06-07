package main

import (
	"fmt"

	"github.com/zehlt/gecs/command"
	"github.com/zehlt/gecs/query"
)

type MoveSystem struct {
}

func (s *MoveSystem) Init(qm query.QueryMaker) query.Query {
	return qm.Create(query.Access{Position{}, Speed{}}, query.Exclude{})
}

func (s *MoveSystem) Exec(cmd command.Controller, q query.Query) {

	q.ForEach(func(data query.QueryData) {
		fmt.Println("MOVINGGG !! entity:", data.E)
	})
}

// ---

type EnemyBarkSystem struct {
}

func (s *EnemyBarkSystem) Init(qm query.QueryMaker) query.Query {
	return qm.Create(query.Access{Enemy{}}, query.Exclude{})
}

func (s *EnemyBarkSystem) Exec(cmd command.Controller, q query.Query) {

	q.ForEach(func(data query.QueryData) {
		fmt.Println("BARK!! entity:", data.E)
	})
}
