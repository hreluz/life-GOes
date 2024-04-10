package main

import "fmt"

const (
	inputAttack = iota
	inputSpecialAttack
	inputHeal
	inputHeal2 = iota + 10
	inputHeal3
)

func main() {

	fmt.Println(inputAttack)
	fmt.Println(inputSpecialAttack)
	fmt.Println(inputHeal)
	fmt.Println(inputHeal2)
	fmt.Println(inputHeal3)

}
