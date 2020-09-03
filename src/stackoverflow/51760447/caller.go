package caller

import "github.com/mrdulin/golang/src/stackoverflow/51760447/a"

var s = a.A{}
var (
	Perfom = s.Perfom
)

func Invoke(url string) string {
	out, _ := Perfom(url)
	return out
}
