package piscine

func ListForEach(l *List, f func(*NodeL)) {
	if l == nil {
		return
	}

	for node := l.Head; node != nil; node = node.Next {
		f(node)
	}
}
