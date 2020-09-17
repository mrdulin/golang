package sorting_test

import (
	"github.com/mrdulin/go-datastructure-algorithm/sorting"
	"github.com/mrdulin/go-datastructure-algorithm/util"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Run("should sort int slice correctly", func(t *testing.T) {
		a := util.GenerateSeries(5)
		got := sorting.SelectionSort(a)
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
		if !reflect.DeepEqual(a, []int{5, 4, 3, 2, 1}) {
			t.Errorf("shoule not mutate input slice, got: %v", a)
		}
	})
}

//cd sorting
//go test -bench=SelectionSort -cpu=1 -count=5
//goos: darwin
//goarch: amd64
//pkg: github.com/mrdulin/go-datastructure-algorithm/sorting
//BenchmarkSelectionSort      2424            444651 ns/op
//BenchmarkSelectionSort      2714            480189 ns/op
//BenchmarkSelectionSort      2358            474588 ns/op
//BenchmarkSelectionSort      2601            520246 ns/op
//BenchmarkSelectionSort      2613            631061 ns/op
//PASS
//ok      github.com/mrdulin/go-datastructure-algorithm/sorting   11.379s
func BenchmarkSelectionSort(b *testing.B) {
	a := util.GenerateSeries(1000)
	for i := 0; i < b.N; i++ {
		sorting.SelectionSort(a)
	}
}
