package scheduler

import (
	"github.com/zehlt/gecs/command"
	"github.com/zehlt/gecs/query"
)

type System interface {
	Init(qm query.QueryMaker) query.Query
	Exec(cmd command.Controller, q query.Query)
}
