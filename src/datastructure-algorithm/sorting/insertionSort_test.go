package sorting_test

import (
	"github.com/mrdulin/go-datastructure-algorithm/sorting"
	"github.com/mrdulin/go-datastructure-algorithm/util"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	t.Run("should sort int slice correctly", func(t *testing.T) {
		a := util.GenerateSeries(5)
		got := sorting.InsertionSort(a)
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
//BenchmarkInsertionSort      3795            290617 ns/op
//BenchmarkInsertionSort      4162            284867 ns/op
//BenchmarkInsertionSort      3670            285245 ns/op
//BenchmarkInsertionSort      4232            292365 ns/op
//BenchmarkInsertionSort      4210            285274 ns/op
//PASS
//ok      github.com/mrdulin/go-datastructure-algorithm/sorting   9.828s
func BenchmarkInsertionSort(b *testing.B) {
	a := util.GenerateSeries(1000)
	for i := 0; i < b.N; i++ {
		sorting.InsertionSort(a)
	}
}
