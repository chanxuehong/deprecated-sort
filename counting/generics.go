package counting

// if you care about performance, you can copy this code into your project and modify element to meet your own needs

type element struct {
	ComparedField int
	Others        [1]byte
}

func genericsSort(dst, src []element) {
	if len(dst) < len(src) {
		panic("len(dst) < len(src)")
	}
	if len(src) < 0 {
		panic("len(src) < 0")
	}
	if len(src) == 0 {
		return
	}
	if len(src) == 1 {
		dst[0] = src[0]
		return
	}
	min, max := getGenericsMinMax(src)
	count := make([]int, max-min+1)
	for _, v := range src {
		count[v.ComparedField-min]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for i := len(src) - 1; i >= 0; i-- {
		v := src[i]
		countIndex := v.ComparedField - min
		dstIndex := count[countIndex] - 1
		dst[dstIndex] = v
		count[countIndex]--
	}
}

func getGenericsMinMax(arr []element) (min, max int) {
	min, max = arr[0].ComparedField, arr[0].ComparedField
	for _, v := range arr {
		switch n := v.ComparedField; {
		case n > max:
			max = n
		case n < min:
			min = n
		}
	}
	return
}
