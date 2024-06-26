package functions

import (
	"fmt"
	"time"
)

type MyFunction func(string)

func MiddlewareLog(f MyFunction) MyFunction {
	return func(name string) {
		fmt.Println("start", time.Now().Format("2006-01-02 15:04:05"))
		f(name)
		fmt.Println("end", time.Now().Format("2006-01-02 15:04:05"))
	}
}
