package sorting

func MergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	middle := len(s) / 2
	return merge(MergeSort(s[:middle]), MergeSort(s[middle:]))
}

func merge(a, b []int) []int {
	var (
		i = 0
		j = 0
		c = []int{}
	)
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	if i < len(a) {
		c = append(c, a[i:]...)
	} else {
		c = append(c, b[j:]...)
	}
	return c
}
