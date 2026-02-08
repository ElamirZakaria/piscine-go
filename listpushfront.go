package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	n.Next = l.Head
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Head = n
	}
}
