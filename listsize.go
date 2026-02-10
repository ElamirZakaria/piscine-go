package piscine

func ListSize(l *List) int {
	tempNode := &NodeL{}
	tempNode = l.Head
	count := 0
	if l.Head == nil {
		return count
	} else {
		for tempNode != nil {
			tempNode = tempNode.Next
			count++
		}
	}
	return count
}
