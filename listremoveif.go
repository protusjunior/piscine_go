package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	var prev *NodeL
	for current != nil {
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}
		} else {
			prev = current
		}
		current = current.Next
	}
	l.Tail = prev
}
