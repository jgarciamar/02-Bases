package FoodCalculator

type AnimalFoodCalculator func(float64) float64

func Animal(a string) (AnimalFoodCalculator, string) {

	FoodPerAnimal := map[string]float64{
		"dog":     250,
		"cat":     5,
		"hamster": 250,
		"spider":  150,
	}

	val, ok := FoodPerAnimal[a]

	if ok {
		return func(petAmount float64) float64 {
			return val * petAmount
		}, ""
	}
	return nil, "Animal doesn't exist"
}
