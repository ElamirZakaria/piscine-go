package piscine

func ToLower(s string) string {
	result := []byte(s)

	for i, a := range result {
		if IsUpper(string(a)) {
			result[i] = a + 32
		} else {
		}
	}
	return string(result)
}
