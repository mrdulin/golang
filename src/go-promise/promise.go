package go_promise

import (
	"sort"
)

type Promiser interface {
	AllSettled(iterable []Workload) []Result
	All(iterable []Workload) []Result
	Race(iterable []Workload) *Result
}

type promise struct {
}

type Workload func(args ...interface{}) interface{}
type Result struct {
	Idx int
	R   interface{}
}

func NewPromise() *promise {
	return &promise{}
}

func (p *promise) AllSettled(iterable []Workload) []Result {
	if len(iterable) == 0 {
		return nil
	}
	c := make(chan struct{}, len(iterable))
	r := make([]Result, len(iterable))
	for idx, iter := range iterable {
		iter := iter
		idx := idx
		go func() {
			r[idx].Idx = idx
			r[idx].R = iter()
			c <- struct{}{}
		}()
	}
	for i := 0; i < len(iterable); i++ {
		<-c
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Idx < r[j].Idx
	})
	close(c)
	return r
}

func (p *promise) All(iterable []Workload) []Result {
	if len(iterable) == 0 {
		return nil
	}
	c := make(chan Result, len(iterable))
	errc := make(chan Result, len(iterable))
	r := []Result{}
	for idx, iter := range iterable {
		iter := iter
		idx := idx
		go func() {
			result := iter()
			if err, ok := result.(error); ok {
				errc <- Result{Idx: idx, R: err}
			} else {
				c <- Result{Idx: idx, R: result}
			}
		}()
	}
	for {
		select {
		case result := <-errc:
			return []Result{result}
		case result := <-c:
			r = append(r, result)
			if len(r) == len(iterable) {
				sort.Slice(r, func(i, j int) bool {
					return r[i].Idx < r[j].Idx
				})
				return r
			}
		}
	}
}

func (p *promise) Race(iterable []Workload) *Result {
	if len(iterable) == 0 {
		return nil
	}
	c := make(chan *Result, len(iterable))
	for _, iter := range iterable {
		iter := iter
		go func() {
			r := Result{R: iter()}
			c <- &r
		}()
	}
	return <-c
}
