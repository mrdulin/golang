package array_test

import (
	"github.com/mrdulin/go-datastructure-algorithm/array"
	"github.com/mrdulin/go-datastructure-algorithm/util"
	"reflect"
	"testing"
)

func TestFindDupByHash(t *testing.T) {
	t.Run("should find duplicated element", func(t *testing.T) {
		s := []int{1,3,4,2,5,3}
		got := array.FindDupByHash(s)
		want := 3
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %d, want: %d", got, want)
		}
	})

	t.Run("should find duplicated element for big slice", func(t *testing.T) {
		s := util.GenerateSeries(1000)
		s = append(s, 123)
		got := array.FindDupByHash(s)
		want := 123
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %d, want: %d", got, want)
		}
	})
}