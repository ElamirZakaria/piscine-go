package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}
	temp := f(a[0], a[1])
	if temp > 0 {
		for i := 1; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	} else if temp < 0 {
		for i := 1; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	} else if temp == 0 {
		for i := 1; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) != 0 {
				return false
			}
		}
	}
	return true
}
