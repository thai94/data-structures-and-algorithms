package main

/*
https://leetcode.com/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var first *ListNode
	var pre *ListNode
	node1 := l1
	node2 := l2
	mem := 0
	for {

		v1 := 0
		v2 := 0
		if node1 != nil {
			v1 = node1.Val
		}

		if node2 != nil {
			v2 = node2.Val
		}

		val_tmp := v1 + v2 + mem
		val := val_tmp % 10
		mem = val_tmp / 10
		if first == nil {
			first = &ListNode{}
			first.Val = val
			pre = first
		} else {
			pre.Next = &ListNode{
				Val: val,
			}
			pre = pre.Next
		}

		if node1 != nil && node1.Next != nil {
			node1 = node1.Next
		} else {
			node1 = nil
		}

		if node2 != nil && node2.Next != nil {
			node2 = node2.Next
		} else {
			node2 = nil
		}

		if node1 == nil && node2 == nil {
			if mem > 0 {
				pre.Next = &ListNode{
					Val: mem,
				}
			}
			break
		}
	}

	return first
}

func main() {}
