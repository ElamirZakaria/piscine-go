package piscine

func Split(s, sep string) []string {
	var answer []string
	lengthofseprator := len(sep)
	lengthofstring := len(s)
	var temp string = ""
	// var flag bool = true
	for i := 0; i < lengthofstring; i++ {
		if s[i] == sep[0] {
			temp2 := i
			temp3 := i
			for j := 0; j < lengthofseprator; j++ {
				if sep[j] == s[temp2] && temp2 < lengthofstring {
					temp2++
				} else {
					// flag = true
					i = temp3
					break
				}
				i = temp2
			}
			if temp != "" {
				answer = append(answer, temp)
				temp = ""
			}
		}
		temp = temp + string(s[i])
	}
	if temp != "" {
		answer = append(answer, temp)
		temp = ""
	}
	return answer
}
