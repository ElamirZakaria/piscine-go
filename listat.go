package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	tempNode := l
	count := 0
	if l == nil {
		return nil
	}
	for tempNode != nil {
		if count == pos {
			return tempNode
		}

		count++
		tempNode = tempNode.Next
	}
	return nil
}
