package a

type A struct{}

func (a *A) Perfom(url string) (string, error) {
	return "real data", nil
}
