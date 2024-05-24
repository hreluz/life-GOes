package main

import "github.com/hreluz/go-middlewares/functions"

func execute(name string, f functions.MyFunction) {
	f(name)
}

func main() {
	name := "Github community"

	execute(name, functions.MiddlewareLog(functions.Hello))
	execute(name, functions.MiddlewareLog(functions.GoodBye))
}
