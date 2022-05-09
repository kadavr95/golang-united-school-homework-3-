package homework

func average(input [15]float32) (result float32) {
	return sum(input) / 15
}

func sum(array [15]float32) float32 {
	var res float32 = 0.0
	for _, e := range array {
		res += e
	}
	return res
}
