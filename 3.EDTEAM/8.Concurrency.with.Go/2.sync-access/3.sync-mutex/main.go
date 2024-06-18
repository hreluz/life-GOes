package main

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	name    string
	balance int
}

func (a *account) String() string {
	return fmt.Sprintf("(%s: Balance: %d)", a.name, a.balance)
}

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(2)

	ironman := account{name: "Ironman", balance: 500}
	batman := account{name: "Batman", balance: 900}

	for _, value := range []int{300, 300} {
		go func(amount int) {
			mu.Lock()
			transfer(amount, &ironman, &batman)
			mu.Unlock()
			wg.Done()
		}(value)
	}
	wg.Wait()
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("❌: %s\n", fmt.Sprintf("%s %s", source, dest))
		return
	}

	time.Sleep(time.Second)

	dest.balance += amount
	source.balance -= amount

	fmt.Printf("✅: %s\n", fmt.Sprintf("%s %s", source, dest))
}
