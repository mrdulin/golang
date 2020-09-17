package sorting_test

import (
	"github.com/mrdulin/go-datastructure-algorithm/sorting"
	"github.com/mrdulin/go-datastructure-algorithm/util"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("should sort int slice correctly", func(t *testing.T) {
		a := util.GenerateSeries(5)
		got := sorting.QuickSort(a)
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
		if !reflect.DeepEqual(a, []int{5, 4, 3, 2, 1}) {
			t.Errorf("shoule not mutate input slice, got: %v", a)
		}
	})
}

//goos: darwin
//goarch: amd64
//pkg: github.com/mrdulin/go-datastructure-algorithm/sorting
//BenchmarkQuickSort         86059             14383 ns/op
//BenchmarkQuickSort         73132             17112 ns/op
//BenchmarkQuickSort         72796             14469 ns/op
//BenchmarkQuickSort         74786             14009 ns/op
//BenchmarkQuickSort         78560             13882 ns/op
//PASS
//ok      github.com/mrdulin/go-datastructure-algorithm/sorting   8.620s
func BenchmarkQuickSort(b *testing.B) {
	a := util.GenerateSeries(1000)
	for i := 0; i < b.N; i++ {
		sorting.QuickSort(a)
	}
}
