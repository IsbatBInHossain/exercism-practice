package lasagna

func PreparationTime (layers []string, avgTime int) int {
    if (avgTime == 0){
        return 2 * len(layers)
    }
    return avgTime * len(layers)
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } 
        if layer == "sauce" {
            sauce += 0.2
        }
    }
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    var scaledQuantities []float64 
    
    for _, val := range quantities {
        scaledVal := (val * float64(portions))/2.0
        scaledQuantities = append(scaledQuantities, scaledVal)
    }
    return scaledQuantities
}


