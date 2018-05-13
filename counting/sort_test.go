package counting

import (
	"math/rand"
	"sort"
	"testing"
)

type intSlice []int

func (p intSlice) Len() int                    { return len(p) }
func (p intSlice) ComparedField(index int) int { return p[index] }
func (p intSlice) CopyElement(dst Interface, dstIndex int, srcIndex int) {
	dst.(intSlice)[dstIndex] = p[srcIndex]
}

func TestSort(t *testing.T) {
	arr := make([]int, 10000) // [-100,100)
	for i := range arr {
		arr[i] = rand.Intn(200) - 100
	}
	dst := make([]int, 10000)
	Sort(intSlice(dst), intSlice(arr))
	if !sort.IntsAreSorted(dst) {
		t.Error("function Sort did not work correctly")
		return
	}
}

func BenchmarkSort(b *testing.B) {
	arr := make([]int, 1e6)
	for i := range arr {
		arr[i] = rand.Intn(2400)
	}
	dst := make([]int, 1e6)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(intSlice(dst), intSlice(arr))
	}
}
