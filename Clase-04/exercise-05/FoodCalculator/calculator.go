package FoodCalculator

type AnimalFoodCalculator func(int) float64

func Animal(a string) (AnimalFoodCalculator, string) {

	FoodPerAnimal := map[string]float64{
		"dog":     250,
		"cat":     5,
		"hamster": 250,
		"spider":  150,
	}

	val, ok := FoodPerAnimal[a]

	if ok {
		return func(petAmount int) float64 {
			return val * float64(petAmount)
		}, ""
	}
	return nil, "Animal doesn't exist"
}
