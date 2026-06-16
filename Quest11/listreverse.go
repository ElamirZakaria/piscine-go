package piscine

func ListReverse(l *List) {
	var previousNode *NodeL = nil
	currentNode := l.Head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	l.Tail = l.Head
	l.Head = previousNode
}
