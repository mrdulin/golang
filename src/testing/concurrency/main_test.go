package main_test

import (
  "testing"
  "testing/concurrency"
  "sync"
)

func TestUser_DeleteTask(t *testing.T) {
  t.Skip()
  var (
    nOps = 100
    wg sync.WaitGroup
  )
  tasks := map[string]string{
    "1": "a",
    "2": "b",
  }
  u := &main.User{
    Tasks: tasks,
  }
  wg.Add(nOps)
  for i := 0; i < nOps; i++ {
    go func() {
      defer wg.Done()
      u.DeleteTask("1")
    }()
  }
  wg.Wait();
  
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


func TestSafeCounter_Value(t *testing.T) {
  var (
    nOps = 100
    wg sync.WaitGroup
  )
  counter := main.SafeCounter{V: make(map[string]int)}
  wg.Add(nOps)
  for i := 0; i < nOps; i++ {
    go func() {
      defer wg.Done()
      if counter.Value("a") == 0 {
        counter.Inc("a")
      }
    }()
  }
  wg.Wait()
  if counter.Value("a") != 1 {
    t.Errorf("the value of key a should be %d, got: %d", 1, counter.Value("a"))
  }
}