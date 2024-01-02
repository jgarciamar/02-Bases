package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Worker interface {
	work(int, int)
}

func (p Person) work(a int, b int) {
	fmt.Printf("I work %d hours, %d days", a, b)
}

func PrintWork(a Worker) {
	a.work(3, 4)
}

func main() {

	camilo := Person{"Camilo", 20}
	camilo.work(8, 5)
	PrintWork(camilo)

}
