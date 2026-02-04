package piscine

func ShoppingListSort(slice []string) []string {
	length := len(slice)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if len(slice[j+1]) < len(slice[j]) {
				temp := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}
