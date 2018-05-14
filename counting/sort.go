package counting

type Interface interface {
	Len() int
	ComparedField(index int) int
	CopyElement(dst Interface, dstIndex int, srcIndex int)
}

func Sort(dst, src Interface) {
	srcLen := src.Len()
	if dst.Len() < srcLen {
		panic("dst.Len() < src.Len()")
	}
	if srcLen < 0 {
		panic("src.Len() < 0")
	}
	if srcLen == 0 {
		return
	}
	if srcLen == 1 {
		src.CopyElement(dst, 0, 0)
		return
	}
	min, max := getMinMax(src)
	count := make([]int, max-min+1)
	for i := 0; i < srcLen; i++ {
		count[src.ComparedField(i)-min]++
	}
	total := 0
	for i, c := range count {
		count[i] = total
		total += c
	}
	for i := 0; i < srcLen; i++ {
		countKey := src.ComparedField(i) - min
		src.CopyElement(dst, count[countKey], i)
		count[countKey]++
	}
}

func getMinMax(data Interface) (min, max int) {
	min = data.ComparedField(0)
	max = min
	for i, l := 1, data.Len(); i < l; i++ {
		switch n := data.ComparedField(i); {
		case n > max:
			max = n
		case n < min:
			min = n
		}
	}
	return
}
