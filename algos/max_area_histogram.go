package algos

func FindMax(arr ...int) int {
	max := -1
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func FindMaxAreaHistogram(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min_index := -1
	MaxUint := ^uint(0)
	min_val := int(MaxUint >> 1)
	for index, val := range arr {
		if val < min_val {
			min_index = index
			min_val = val
		}
	}
	return FindMax(min_val*len(arr), FindMaxAreaHistogram(arr[:min_index]), FindMaxAreaHistogram(arr[min_index+1:]))
}
