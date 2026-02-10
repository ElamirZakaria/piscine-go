package piscine

func ListLast(l *List) interface{} {
	tempNode := &NodeL{}
	tempNode = l.Head
	if l.Head == nil {
		return nil
	}
	if l.Head == l.Tail {
		return l.Head.Data
	}

	for tempNode != l.Tail {
		tempNode = tempNode.Next
	}

	return tempNode.Data
}
