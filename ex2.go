package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p *ListNode
	carry := 0
	for l1.Next != nil || l2.Next != nil || carry != 0 {
		if l1.Val == nil {

		}

	}
}