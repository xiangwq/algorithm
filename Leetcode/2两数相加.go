package Leetcode

import "algorithm/helper"

func addTwoNumbers(l1 *helper.ListNode, l2 *helper.ListNode) *helper.ListNode {
	node := &helper.ListNode{
		Val:  (l1.Val + l2.Val) % 10,
		Next: nil,
	}
	nextAddNum := (l1.Val + l2.Val) / 10
	l1 = l1.Next
	l2 = l2.Next
	temp := node
	for {
		if l1 == nil && l2 == nil {
			if nextAddNum != 0 {
				tempNode := &helper.ListNode{
					Val:  nextAddNum,
					Next: nil,
				}
				temp.Next = tempNode
			}
			break
		}
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		tempNode := &helper.ListNode{
			Val:  (v1 + v2 + nextAddNum) % 10,
			Next: nil,
		}
		nextAddNum = (v1 + v2 + nextAddNum) / 10

		temp.Next = tempNode
		temp = temp.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return node
}
