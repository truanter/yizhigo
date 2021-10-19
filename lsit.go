package main

type Node struct {
	Val  int
	Next *Node
}

func (n Node) construct(l []int) Node {
	var head Node
	if len(l) <= 0 {
		return head
	}
	head = Node{
		Val:  l[0],
		Next: nil,
	}
	cur := head
	for _, v := range l[1:] {
		c := Node{
			Val:  v,
			Next: nil,
		}
		cur.Next = &c
		cur = c
	}
	return head
}

func (n Node) Remove(val int) Node {
	if n.Next == nil {
		return n
	}
	head := n
	for head.Val == val {
		head = *head.Next
	}
	pre, cur := head, head.Next
	for cur.Next != nil {
		if cur.Val != val {
			pre = *cur
		}
		cur = cur.Next
		pre.Next = cur
	}
	return head
}
