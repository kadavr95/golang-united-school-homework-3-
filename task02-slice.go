package homework

func reverse(input []int64) (result []int64) {
	l := len(input)
	res := make([]int64, l)
	for i, e := range input {
		res[l-i-1] = e
	}
	return res
}
