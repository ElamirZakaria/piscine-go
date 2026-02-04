package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := 0; j < len(podium)-i-1; j++ {
			if podium[j][0] > (podium[j+1][0]) {
				a := (podium[j][0])
				podium[j][0] = podium[j+1][0]
				podium[j+1][0] = a
			}
		}
	}
	return podium
}
