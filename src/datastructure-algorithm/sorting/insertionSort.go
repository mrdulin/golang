package sorting

func InsertionSort(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	l := len(s)
	for i := 1; i < l; i++ {
		j := i
		temp := c[i]
		for j > 0 && (c[j-1] > temp) {
			c[j] = c[j-1]
			j--
		}
		c[j] = temp
	}
	return c
}
