package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	str := ""
	flag := 0
	for _, a := range s {
		if ((a == ' ') || (a == '\n') || (a == '\t')) && (flag == 1) {
			answer = append(answer, str)
			str = ""
			flag = 0

		} else {
			str = str + string(a)
			flag = 1
		}
	}
	answer = append(answer, str)
	return answer
}

// SplitWhiteSpaces("Hello how are you?") == []string{" Hello", "how", "are", "you?"} instead of []string{"Hello", "how", "are", "you?"}
// SplitWhiteSpaces("&4/[%$e{X&V   @1LE9-I)spO4X ($eg)^0%l\\X;i 7MiyU`d_o[YBm (L}7:m1Ln=^Cf iOJH~?G{NrKVn oRs+F.\\|ce]>t NWgZ#@D[NV/O.") == []string{"&4/[%$e{X&V", "", "", "@1LE9-I)spO4X", "($eg)^0%l\\X;i", "7MiyU`d_o[YBm", "(L}7:m1Ln=^Cf", "iOJH~?G{NrKVn", "oRs+F.\\|ce]>t", "NWgZ#@D[NV/O."} instead of []string{"&4/[%$e{X&V", "@1LE9-I)spO4X", "($eg)^0%l\\X;i", "7MiyU`d_o[YBm", "(L}7:m1Ln=^Cf", "iOJH~?G{NrKVn", "oRs+F.\\|ce]>t", "NWgZ#@D[NV/O."}
