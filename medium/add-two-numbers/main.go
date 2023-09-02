package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	startNode := &ListNode{}
	currentNode := startNode
	front := 0
	val1 := l1.Val
	val2 := l2.Val
	for {
		currentNode.Val = (val1 + val2 + front) % 10
		front = (val1 + val2 + front) / 10
		if l1 != nil {
			l1 = l1.Next
			if l1 != nil {
				val1 = l1.Val
			} else {
				val1 = 0
			}
		} else {
			l1 = nil
			val1 = 0
		}
		if l2 != nil {
			l2 = l2.Next
			if l2 != nil {
				val2 = l2.Val
			} else {
				val2 = 0
			}
		} else {
			l2 = nil
			val2 = 0
		}
		if l1 == nil && l2 == nil {
			break
		}
		new_node := &ListNode{}
		currentNode.Next = new_node
		currentNode = new_node
	}
	if front > 0 {
		currentNode.Val = currentNode.Val % 10
		new_node := &ListNode{}
		currentNode.Next = new_node
		new_node.Val = 1
	}
	return startNode
}

func main() {

}
