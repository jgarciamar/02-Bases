package main

import "fmt"

var temp int8
var humidity float32
var pressure float32

func main() {
	temp = 10
	humidity = 3.2
	pressure = 0.4
	fmt.Printf("Temp: %d Humidity: %.3f Pressure: %.3f \n", temp, humidity, pressure)

}
