package main

import (
	"fmt"
	"time"
)

var memo map[int]int

func fibonacciWithMemo(n int) int {
	if n <= 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	if _, ok := memo[n]; ok {
		return memo[n]
	}

	memo[(n - 1)] = fibonacciWithMemo(n - 1)
	memo[(n - 2)] = fibonacciWithMemo(n - 2)

	return memo[(n-1)] + memo[(n-2)]
}

func fibonacci(n int) int {
	if n <= 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {

	n := 30
	memo = make(map[int]int)

	start := time.Now().UnixNano() / int64(time.Millisecond)
	for i := 1; i <= n; i++ {
		fmt.Println(fibonacciWithMemo(i))
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("Fibonacci with memo took %d s \n\n\n", (end - start))

	start = time.Now().UnixNano() / int64(time.Millisecond)
	for i := 1; i <= n; i++ {
		fmt.Println(fibonacci(i))
	}
	end = time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("Fibonacci without memo took %d s \n\n\n", (end - start))
}
