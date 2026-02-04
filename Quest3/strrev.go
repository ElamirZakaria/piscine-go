package piscine

func StrRev(s string) string {
	var stringchanged string
	stringChangeable := []byte(s)
	length := 0
	for range s {
		length++
	}
	for i := 0; i < (length)/2; i++ {
		temp := stringChangeable[length-i-1]
		stringChangeable[length-i-1] = s[i]
		stringChangeable[i] = temp
	}
	stringchanged = string(stringChangeable)
	return stringchanged
}
