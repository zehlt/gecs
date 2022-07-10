package gecs

import (
	"fmt"
	"reflect"
)

type Parser struct {
}

func (p *Parser) Parse(fn interface{}) {

	t := reflect.TypeOf(fn)
	fmt.Println("Function arguments:")
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf(" %d. %v\n", i, t.In(i))
	}
	fmt.Println("Function return values:")
	for i := 0; i < t.NumOut(); i++ {
		fmt.Printf(" %d. %v\n", i, t.Out(i))
	}
}
