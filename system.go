package gecs

type System interface {
	Init(qm QueryMaker) Query
	Exec(cmd Controller, q Query)
}
