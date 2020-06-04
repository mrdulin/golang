package main

import "sync"

type User struct {
  Tasks map[string]string
}

func (u *User) DeleteTask(taskId string) {
  delete(u.Tasks, taskId)
} 

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	V   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.V[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.V[key]
}