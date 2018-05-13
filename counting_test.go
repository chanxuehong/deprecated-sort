package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := make([]int, 10000) // [-100,100)
	for i := range arr {
		arr[i] = rand.Intn(200) - 100
	}
	CountingSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Error("function CountingSort did not work correctly")
		return
	}
}

func TestGetMinMax(t *testing.T) {
	tests := []struct {
		arr []int
		min int
		max int
	}{
		{
			[]int{1},
			1,
			1,
		},
		{
			[]int{1, 2},
			1,
			2,
		},
		{
			[]int{2, 1},
			1,
			2,
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			5,
		},
		{
			[]int{1, 2, 3, 4, 5, 4},
			1,
			5,
		},
		{
			[]int{2, 1, 2, 3, 4, 5},
			1,
			5,
		},
		{
			[]int{2, 1, 2, 3, 4, 5, 4},
			1,
			5,
		},
	}
	for _, test := range tests {
		min, max := getMinMax(test.arr)
		if min != test.min || max != test.max {
			t.Errorf("arr=%v, want=(%d, %d), have=(%d, %d)", test.arr, test.min, test.max, min, max)
			return
		}
	}
}

func BenchmarkCountingSort(b *testing.B) {
	arr := make([]int, 1e6)
	for i := range arr {
		arr[i] = rand.Intn(2400)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountingSort(arr)
	}
}
