package counting

func Ints(arr []int) {
	if len(arr) < 2 {
		return
	}
	min, max := getIntsMinMax(arr)
	count := make([]int, max-min+1)
	for _, v := range arr {
		count[v-min]++
	}
	var j int
	for i, c := range count {
		v := i + min
		for c > 0 {
			arr[j] = v
			j++
			c--
		}
	}
}

func getIntsMinMax(arr []int) (min, max int) {
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
