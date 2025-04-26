package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTime
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	qty := make([]float64, 0)
	var onePortion float64

	for _, v := range quantities {
		onePortion = v / 2
		qty = append(qty, onePortion*float64(numPortions))
	}
	return qty
}
