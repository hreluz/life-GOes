package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {
}

func (p Paypal) Pay() {
	fmt.Println("Paid by Paypal")
}

type Cash struct {
}

func (c Cash) Pay() {
	fmt.Println("Paid by Cash")
}

type CreditCard struct {
}

func (c CreditCard) Pay() {
	fmt.Println("Paid by credit card")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {

	var method uint

	fmt.Println("Enter a number: ")
	fmt.Println("\t 1:Paypal 2:Cash 3:Credit Card")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("You must eneter a valid method")
	}

	if method > 3 {
		panic("You must enter a valid method")
	}

	payMethod := Factory(method)
	payMethod.Pay()

}
