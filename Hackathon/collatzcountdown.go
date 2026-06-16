package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	count := 0
	num := start
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
		count++
	}
	return count
}
