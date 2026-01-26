package piscine

func ToUpper(s string) string {
	result := []byte(s)

	for i, a := range result {
		if IsLower(string(a)) {
			result[i] = a - 32
		} else {
		}
	}
	return string(result)
}
