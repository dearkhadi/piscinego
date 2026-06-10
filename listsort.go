package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}

	for current := l; current != nil; current = current.Next {
		for next := current.Next; next != nil; next = next.Next {
			if current.Data > next.Data {
				current.Data, next.Data = next.Data, current.Data
			}
		}
	}

	return l
}
