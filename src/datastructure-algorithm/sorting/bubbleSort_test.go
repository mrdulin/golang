package sorting_test

import (
	"github.com/mrdulin/go-datastructure-algorithm/util"
	"reflect"
	"testing"
	"github.com/mrdulin/go-datastructure-algorithm/sorting"
)

func TestBubbleSort(t *testing.T) {
	t.Run("should sort int slice correctly", func(t *testing.T) {
		a := util.GenerateSeries(5)
		got := sorting.BubbleSort(a)
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestBubbleSortOptimize(t *testing.T) {
	t.Run("should sort int slice correctly", func(t *testing.T) {
		a := util.GenerateSeries(5)
		got := sorting.BubbleSortOptimize(a)
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// cd sorting
// go test -bench=. -cpu=1 -count=5
// Output:
//goos: darwin
//goarch: amd64
//pkg: github.com/mrdulin/go-datastructure-algorithm/sorting
//BenchmarkBubbleSort          613           1656886 ns/op
//BenchmarkBubbleSort          750           1444395 ns/op
//BenchmarkBubbleSort          750           1425196 ns/op
//BenchmarkBubbleSort          762           1465209 ns/op
//BenchmarkBubbleSort          784           1477291 ns/op
//PASS
//ok      github.com/mrdulin/go-datastructure-algorithm/sorting   6.318s
func BenchmarkBubbleSort(b *testing.B) {
	a := util.GenerateSeries(1000)
	for i := 0; i < b.N; i++ {
		sorting.BubbleSort(a)
	}
}

//go test -bench=. -cpu=1 -count=5
//BenchmarkBubbleSortOptimize         2130            567584 ns/op
//BenchmarkBubbleSortOptimize         2157            561928 ns/op
//BenchmarkBubbleSortOptimize         1864            586374 ns/op
//BenchmarkBubbleSortOptimize         1803            580165 ns/op
//BenchmarkBubbleSortOptimize         1812            566454 ns/op
func BenchmarkBubbleSortOptimize(b *testing.B) {
	a := util.GenerateSeries(1000)
	for i := 0; i < b.N; i++ {
		sorting.BubbleSortOptimize(a)
	}
}


