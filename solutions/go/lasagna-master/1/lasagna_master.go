package lasagna

func PreparationTime(layers []string, avgPreparationTime int) int {
	if avgPreparationTime == 0 {
		avgPreparationTime = 2
	}

	return len(layers) * avgPreparationTime
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendLayers []string, myLayers []string) {
	myLayers[len(myLayers)-1] = friendLayers[len(friendLayers)-1]
}

func ScaleRecipe(twoPortionsAmount []float64, cookPortions int) []float64 {
	newPortions := make([]float64, len(twoPortionsAmount))
	for i := 0; i < len(twoPortionsAmount); i++ {
		newPortions[i] = twoPortionsAmount[i] / 2 * float64(cookPortions)
	}
	return newPortions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
