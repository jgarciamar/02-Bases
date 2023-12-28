package main

import "fmt"

func main() {
	/*
	   var 1firstName string Incorrect
	   var lastName string correct
	   var int age	incorrect
	   1lastName := 6 incorrect
	   var driver_license = true correct but not idiomatic
	   var person height int incorrect
	   childsNumber := 2 correct
	*/

	var firstName string
	var lastName string
	var age int
	driverLicense := true
	var personHeight int
	childsNumber := 2

	fmt.Println(firstName, lastName, age, driverLicense, personHeight, childsNumber)

}
