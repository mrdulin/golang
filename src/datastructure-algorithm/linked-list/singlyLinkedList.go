package linked_list

type node struct {
	Next *node
	Item interface{}
}

func NewNode(item interface{}, next *node) *node {
	return &node{Item: item, Next: next}
}

type singlyLinkedList struct {
	head *node
	length int
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{length: 0, head: nil}
}

func (s *singlyLinkedList) Append(n *node) {
	if s.head == nil {
		s.head = n
	} else {
		next := s.head.Next
		for next.Next != nil {
			next = next.Next
		}
		next.Next = n
	}
	s.length++
}

func (s *singlyLinkedList) Length() int {
	return s.length
}

func Length(head *node) int {
	if head == nil {
		return 0
	}
	l := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		l ++
	}
	return l
}

func GetKthLastElement(head *node, k int) *node {
	var r []*node
	for tail := head; tail != nil; tail = tail.Next {
		r = append([]*node{tail}, r...)
	}
	if r != nil {
		return r[k]
	}
	return nil
}