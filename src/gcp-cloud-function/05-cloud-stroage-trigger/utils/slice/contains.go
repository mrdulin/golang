package slice

func Contains(slice []string, valueToFind string) bool {
	for _, el := range slice {
		if el == valueToFind {
			return true
		}
	}
	return false
}
