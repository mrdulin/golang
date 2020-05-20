package foo

//go:generate mockgen -destination mock_foo/mock_foo.go github.com/mrdulin/golang/src/testing/mock/foo Foo
 
// cd foo && go generate
// cd foo && go test -test.v

type Foo interface {
  Bar(x int) int
}

func SUT(f Foo) int {
  return f.Bar(99)
}