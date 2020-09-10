package go_promise_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	go_promise "github.com/mrdulin/golang/src/go-promise"
)

func workloadGen(i interface{}, options ...interface{}) go_promise.Workload {
	return func(args ...interface{}) interface{} {
		var sec time.Duration = 1
		if len(options) != 0 {
			sec = options[0].(time.Duration)
		}
		time.Sleep(time.Second * sec)
		return i
	}
}

func TestPromise_AllSettled(t *testing.T) {
	t.Run("should settle all workloads and return results", func(t *testing.T) {
		p := go_promise.NewPromise()
		w1 := workloadGen(1)
		w2 := workloadGen(2)
		w3 := workloadGen(3)
		ws := []go_promise.Workload{w1, w2, w3}
		got := p.AllSettled(ws)
		want := []go_promise.Result{{Idx: 0, R: 1}, {Idx: 1, R: 2}, {Idx: 2, R: 3}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %+v, want: %+v", got, want)
		}
	})
	t.Run("should settle empty workloads and return nil", func(t *testing.T) {
		p := go_promise.NewPromise()
		ws := []go_promise.Workload{}
		got := p.AllSettled(ws)
		var want []go_promise.Result
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %#v, want: %#v", got, want)
		}
	})
}

func TestPromise_All(t *testing.T) {
	t.Run("should handle all workloads", func(t *testing.T) {
		p := go_promise.NewPromise()
		w1 := workloadGen(1)
		w2 := workloadGen(2)
		w3 := workloadGen(3)
		ws := []go_promise.Workload{w1, w2, w3}
		got := p.All(ws)
		want := []go_promise.Result{{Idx: 0, R: 1}, {Idx: 1, R: 2}, {Idx: 2, R: 3}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %#v, want: %#v", got, want)
		}
	})

	t.Run("should reject the promise when any workload is rejected", func(t *testing.T) {
		p := go_promise.NewPromise()
		err := errors.New("network")
		w1 := workloadGen(1)
		w2 := workloadGen(2)
		w3 := workloadGen(err)
		ws := []go_promise.Workload{w1, w2, w3}
		got := p.All(ws)
		want := []go_promise.Result{{Idx: 2, R: err}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %+v, want: %+v", got, want)
		}
	})
}

func TestPromise_Race(t *testing.T) {
	t.Run("should race all resolved promise", func(t *testing.T) {
		p := go_promise.NewPromise()
		var t1 time.Duration = 2
		var t2 time.Duration = 1
		w1 := workloadGen(1, t1)
		w2 := workloadGen(2, t1)
		w3 := workloadGen(3, t2)
		ws := []go_promise.Workload{w1, w2, w3}
		got := p.Race(ws)
		want := &go_promise.Result{Idx: 0, R: 3}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %+v, want: %+v", got, want)
		}
	})

	t.Run("should race promises with a rejected promise", func(t *testing.T) {
		p := go_promise.NewPromise()
		err := errors.New("network")
		var t1 time.Duration = 2
		var t2 time.Duration = 1
		w1 := workloadGen(1, t1)
		w2 := workloadGen(2, t1)
		w3 := workloadGen(err, t2)
		ws := []go_promise.Workload{w1, w2, w3}
		got := p.Race(ws)
		want := &go_promise.Result{Idx: 0, R: err}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got: %+v, want: %+v", got, want)
		}
	})
}
