package sorting

func BubbleSort(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if c[j] > c[j+1] {
				t := c[j]
				c[j] = c[j+1]
				c[j+1] = t
			}
		}
	}
	return c
}


func BubbleSortOptimize(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if c[j] > c[j+1] {
				t := c[j]
				c[j] = c[j+1]
				c[j+1] = t
			}
		}
	}
	return c
}