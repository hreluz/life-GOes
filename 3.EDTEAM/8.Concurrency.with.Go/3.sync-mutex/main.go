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

func main() {

	ironman := account{name: "Ironman", balance: 500}
	batman := account{name: "Batman", balance: 900}

	for _, value := range []int{300, 300} {
		transfer(value, &ironman, &batman)
	}
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
