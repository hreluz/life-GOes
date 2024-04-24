package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

type Employee struct {
	Person
	Salary float64
}

func NewEmployee(name string, age uint, salary float64) Employee {
	return Employee{
		NewPerson(name, age),
		salary,
	}
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.9)
}

func main() {
	e := NewEmployee("Batman", 30, 1000)
	fmt.Println(e.Name, e.Age)
	e.Greet()
	e.Payroll()
}
