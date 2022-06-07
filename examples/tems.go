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

		fmt.Println("The entity that matches:", data.E)
	})
}
