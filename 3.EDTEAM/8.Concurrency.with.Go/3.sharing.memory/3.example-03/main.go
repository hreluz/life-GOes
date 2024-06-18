package main

import (
	"fmt"
	"time"
)

type account struct {
	name    string
	balance int
}

func (a *account) String() string {
	return fmt.Sprintf("(%s: Balance: %d)", a.name, a.balance)
}

type bankOperation struct {
	amount int
	done   chan struct{}
}

func main() {
	signal := make(chan struct{})
	transaction := make(chan *bankOperation)

	ironman := account{name: "Ironman", balance: 500}
	batman := account{name: "Batman", balance: 900}

	go func() {
		for {
			request := <-transaction
			transfer(request.amount, &ironman, &batman)
			request.done <- struct{}{}
		}
	}()

	for _, value := range []int{300, 300} {
		go func(amount int) {
			requestTransaction := bankOperation{amount: amount, done: make(chan struct{})}
			transaction <- &requestTransaction

			<-requestTransaction.done
			signal <- struct{}{}
		}(value)
	}

	<-signal
	<-signal
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
