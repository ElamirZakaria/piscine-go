package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i+1], a[i]) < 0 {
			return false
		}
	}
	return true
}
