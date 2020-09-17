package sorting

func QuickSort(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	return quick(c, 0, len(c)-1)
}

func quick(s []int, left, right int) []int {
	if len(s) > 1 {
		index := partition(s, left, right)
		if left < index-1 {
			quick(s, left, index-1)
		}
		if index < right {
			quick(s, index, right)
		}
	}
	return s
}

func partition(s []int, left, right int) int {
	pivot := s[(left+right)/2]
	i := left
	j := right
	for i <= j {
		for s[i] < pivot {
			i++
		}
		for s[j] > pivot {
			j--
		}
		if i <= j {
			t := s[j]
			s[j] = s[i]
			s[i] = t
			i++
			j--
		}
	}
	return i
}
