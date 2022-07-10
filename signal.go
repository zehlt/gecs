package gecs

type Receiver interface {
	Init() interface{}
	Exec(cmd Controller, signal interface{})
}
