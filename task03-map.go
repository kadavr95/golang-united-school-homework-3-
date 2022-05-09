package homework

import (
	//Oops, you should be 18 to do this kind of thing ;)
	//"golang.org/x/exp/maps"
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	l := len(input)
	//Anyway, we are learning here
	//keys := maps.Keys(input)
	keys := make([]int, l)
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	res := make([]string, l)
	for i, e := range keys {
		res[i] = input[e]
	}
	return res
}
