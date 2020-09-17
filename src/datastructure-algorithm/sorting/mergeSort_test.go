package sorting_test

import (
	"github.com/mrdulin/go-datastructure-algorithm/sorting"
	"github.com/mrdulin/go-datastructure-algorithm/util"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("should sort int slice correctly", func(t *testing.T) {
		a := util.GenerateSeries(5)
		got := sorting.MergeSort(a)
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
		if !reflect.DeepEqual(a, []int{5, 4, 3, 2, 1}) {
			t.Errorf("shoule not mutate input slice, got: %v", a)
		}
	})
}

// cd sorting &&
// go test -bench=MergeSort -cpu=1 -count=5
//goos: darwin
//goarch: amd64
//pkg: github.com/mrdulin/go-datastructure-algorithm/sorting
//BenchmarkMergeSort          8228            167722 ns/op
//BenchmarkMergeSort          8144            163918 ns/op
//BenchmarkMergeSort          8353            164401 ns/op
//BenchmarkMergeSort          8174            173985 ns/op
//BenchmarkMergeSort          7459            201798 ns/op
//PASS
//ok      github.com/mrdulin/go-datastructure-algorithm/sorting   9.218s
func BenchmarkMergeSort(b *testing.B) {
	a := util.GenerateSeries(1000)
	for i := 0; i < b.N; i++ {
		sorting.MergeSort(a)
	}
}
