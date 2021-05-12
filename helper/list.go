package helper

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateList(nums []int) *ListNode {
	start := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	temp := start
	for _, v := range nums[1:] {
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		temp.Next = node
		temp = temp.Next
	}
	return start
}

func ShowList(node *ListNode) {
	for {
		if node != nil {
			fmt.Print(node.Val, "\t")
			node = node.Next
			continue
		}
		break
	}
	fmt.Print("\n")
}
