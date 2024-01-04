package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateofBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("%d, %s born in %s, his position is %s\n",
		e.ID, e.Name, e.DateofBirth, e.Position)
}

func main() {
	person := Person{
		ID:          0,
		Name:        "Juan",
		DateofBirth: "date"}

	employee := Employee{0, "Developer", person}

	employee.PrintEmployee()

}
