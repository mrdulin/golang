package main_test

import (
  "testing"
  "testing/concurrency"
  "sync"
)

func TestUser_DeleteTask(t *testing.T) {
  tasks := map[string]string{
    "1": "a",
    "2": "b",
  }
  u := &main.User{
    Tasks: tasks,
  }
  u.DeleteTask("1")
  if u.Tasks["1"] != "" {
    t.Errorf("task with taskId 1 should be deleted, got: %s", u.Tasks["1"])
  }
}

func TestSafeCounter_Inc(t *testing.T) {
  var (
    nOps = 100
    wg sync.WaitGroup
  )
  counter := main.SafeCounter{V: make(map[string]int)}
  wg.Add(nOps)
  for i := 0; i < nOps; i++ {
    go func() {
      defer wg.Done()
      counter.Inc("a")
    }()
  }
  wg.Wait()
  
  if counter.Value("a") != nOps {
    t.Errorf("the value of key a should be %d, got: %d", nOps, counter.Value("a"))
  }
}