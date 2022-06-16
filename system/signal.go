package system

import "github.com/zehlt/gecs/command"

type Receiver interface {
	Init() interface{}
	Exec(cmd command.Controller, signal interface{})
}
