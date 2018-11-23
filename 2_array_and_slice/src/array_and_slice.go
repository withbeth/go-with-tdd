package main

// [5]int and 4[int] are different type as string and int are.
// So, use "slice"
func Sum(numbers []int)(result int) {
	// Index and value
	for _, val := range numbers {
		result += val
	}
	return
}

// Go can let you write variadic functions than can have variable numbers of params.
func SumAll(slices ...[]int)(result []int) {

	// Better solution : use append() for avoiding capacity access problem(will cause run time err)
	// append function takes a slice and new value, returning
	// a new slice with all the items in it.
	for _, slice := range slices {
		result = append(result, Sum(slice))
	}
	return
}

func SumAllTail(slices ...[]int)(result []int) {
	for _, slice := range slices {
		sliceSize := len(slice)
		if sliceSize > 0 {
			last := slice[sliceSize - 1]
			result = append(result, last)
		}
	}
	return
}