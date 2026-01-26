package piscine

func Index(s string, toFind string) int {
	runes := []rune(toFind)
	lengths := 0
	indexrunes := 0
	firstOccurence := 0
	flag := 0
	lengthstring := 0
	for range runes {
		lengths++
	}
	for range s {
		lengthstring++
	}
	if lengths == 0 || lengths > lengthstring {
		return -1
	}
	for i, a := range s {
		if a == runes[indexrunes] {
			if flag == 0 {
				firstOccurence = i
				flag = 1
			}
			indexrunes++
		} else {
			continue
		}
	}
	if indexrunes == lengths {
		return firstOccurence
	} else {
		return -1
	}
}
