package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}

	for i := 0; l != nil; i++ {
		if i == pos {
			return l
		}
		l = l.Next
	}

	return nil
}
