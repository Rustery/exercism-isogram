package isogram

func IsIsogram(input string) bool {
	var result bool = true

	for _, v := range input {
		for _, vv := range input {
			if v == vv {
				result = false
			}
		}
	}

	return result
}