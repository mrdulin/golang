package sorting

func SelectionSort(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	l := len(s)
	for i := 0; i < l; i++ {
		minIdx := i
		for j := i; j < l; j++ {
			if c[minIdx] > c[j] {
				minIdx = j
			}
		}
		if i != minIdx {
			t := c[i]
			c[i] = c[minIdx]
			c[minIdx] = t
		}
	}
	return c
}
