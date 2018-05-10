package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1,nil}
	l2 := &ListNode{2,nil}
	l1.Next = &ListNode{9,nil}
	l2.Next = &ListNode{8,nil}
	l1.Next.Next = &ListNode{5,nil}
	l2.Next.Next = &ListNode{9,nil}

	p := addTwoNumbers(l1, l2)
	fmt.Println(p.Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1,s2,sum,out int
	var head *ListNode
	var rear *ListNode
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			s1 = 0
		}else {
			s1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			s2 = 0
		}else {
			s2 = l2.Val
			l2 = l2.Next
		}
		sum = s1 + s2 + carry

		carry = sum / 10
		out = sum % 10


		p := &ListNode{out,nil}

		if head == nil {
			head = p
			rear = p
		} else {
			rear.Next = p
			rear = p
		}
	}

	return head
}