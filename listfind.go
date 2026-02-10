package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	tempNode := l.Head

	for tempNode != nil {
		if comp(ref, tempNode.Data) == true {
			return &tempNode.Data
		}
		tempNode = tempNode.Next
	}
	return nil
}
