package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	// fmt.Printf("Value: %v, Type: %T\n", i, i)

	v, ok := i.(string)

	if ok {
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	}

	switch x := i.(type) {

	case string:
		fmt.Printf("\t%s\n", strings.ToUpper(x))
	case int:
		fmt.Printf("\t%v\n", (x * 2))
	default:
		fmt.Printf("Value: %v, Type: %T\n", i, i)
	}
}

func main() {
	// null interface
	// var e exampler
	// e.x()

	// empty interface
	wrapper(12)
	wrapper(14.32)
	wrapper(false)
	wrapper("Batman")
}
