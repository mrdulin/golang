package util

func GenerateSeries(size int) []int {
	var a []int
	for i := size; i > 0; i-- {
		a = append(a, i)
	}
	return a
}