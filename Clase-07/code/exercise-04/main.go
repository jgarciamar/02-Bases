package main

import (
	"fmt"
)

func main() {
	salary := 10000
	if salary < 150000 {
		err := fmt.Errorf("The minimum taxable amount is 150.000 and the entered salary is: %d", salary)
		fmt.Println(err.Error())
	}
}
