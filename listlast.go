package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	new := l.Head
	for new.Next != nil {
		new = new.Next
	}
	return new.Data
}
