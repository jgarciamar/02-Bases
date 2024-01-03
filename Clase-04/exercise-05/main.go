package main

import (
	"exercise-05/FoodCalculator"
	"fmt"
)

func main() {

	const (
		dog     = "dog"
		cat     = "cat"
		hamster = "hamster"
		spider  = "spider"
		pig     = "pig"
	)

	var amount float64

	animalDog, msg := FoodCalculator.Animal(dog)
	if msg != "" {
		fmt.Println("Found some error")
		return
	}
	animalCat, msg := FoodCalculator.Animal(cat)
	animalHamster, msg := FoodCalculator.Animal(hamster)
	animalSpider, msg := FoodCalculator.Animal(spider)

	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(10)
	amount += animalSpider(10)

	fmt.Printf("The amount of food required is %.0f\n", amount)

	animalPig, msg := FoodCalculator.Animal(pig)

	if msg != "" {
		fmt.Println("Found some error")
		return
	}
	fmt.Printf("The amount of food required for the pigs is %.0f\n", animalPig(5))

}
