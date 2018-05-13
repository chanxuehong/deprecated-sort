package sort

func CountingSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	min, max := getMinMax(arr)
	count := make([]int, max-min+1)
	for _, v := range arr {
		count[v-min]++
	}
	var j int
	for i := range count {
		v := i + min
		for count[i] > 0 {
			arr[j] = v
			j++
			count[i]--
		}
	}
}

func getMinMax(arr []int) (min, max int) {
	min, max = arr[0], arr[0]
	for _, v := range arr {
		switch {
		case v > max:
			max = v
		case v < min:
			min = v
		}
	}
	return
}
