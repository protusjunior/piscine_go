package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	var prev *NodeL

	for current != nil && current.Data == data_ref {
		l.Head = current.Next
		current = l.Head
	}

	for current != nil {
		if current.Data == data_ref {
			prev.Next = current.Next
		} else {
			prrent.Next
	}
}