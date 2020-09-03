package caller_test

import (
	"reflect"
	"testing"

	caller "github.com/mrdulin/golang/src/stackoverflow/51760447"
)

func TestInvoke(t *testing.T) {
	oPerfom := caller.Perfom
	perfomCallCount := 0
	caller.Perfom = func(url string) (string, error) {
		perfomCallCount++
		return "fake data", nil
	}
	defer func() {
		caller.Perfom = oPerfom
	}()
	got := caller.Invoke("localhost")
	want := "fake data"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("should return fake data, got:%v, want: %v", got, want)
	}
	if !reflect.DeepEqual(perfomCallCount, 1) {
		t.Errorf("Perfom method should be called once, got:%d", perfomCallCount)
	}
}
