package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListSort(n1)
	n2 = ListSort(n2)

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data <= n2.Data {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	} else {
		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	}
}

type NodeIn struct {
	Data int
	Next *NodeIn
}

func ListSortlm(l *NodeI) *NodeI {
	Head := l
	if Head == nil {
		return nil
	}

	Head.Next = ListSortlm(Head.Next)

	if Head.Next != nil && Head.Data > Head.Next.Data {
		Head = movelm(Head)
	}
	return Head
}

func movelm(l *NodeI) *NodeI {
	p := l
	n := l.Next
	ret := n

	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}

	p.Next = l
	l.Next = n
	return ret
}
