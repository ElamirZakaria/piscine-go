package piscine

func Compare(a, b string) int {
	lengtha := 0
	lengthb := 0

	for range a {
		lengtha++
	}
	for range b {
		lengthb++
	}

	for i := 0; i < lengtha && i < lengthb; i++ {
		if a[i] == b[i] {
			continue
		} else if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	if lengtha > lengthb {
		return 1
	} else if lengtha < lengthb {
		return -1
	} else {
		return 0
	}
}
