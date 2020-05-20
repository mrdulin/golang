package foo

import (
  "testing"
  gomock "github.com/golang/mock/gomock"
  mock_foo "github.com/mrdulin/golang/src/testing/mock/foo/mock_foo"
)

func TestFoo(t *testing.T) {
  ctrl := gomock.NewController(t)
  defer ctrl.Finish()
  m := mock_foo.NewMockFoo(ctrl)
  
  m.EXPECT().Bar(gomock.Eq(99)).Return(101)
  got := SUT(m)
  t.Logf("got: %d", got)
  if got != 101 {
		t.Errorf("Expected 101, got %d", got)
	}
}