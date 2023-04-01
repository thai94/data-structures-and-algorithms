package main

import "fmt"

//https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	ret := &ListNode{}
	p := ret

	pNode1 := list1
	pNode2 := list2

	for {
		if (pNode1 != nil && pNode2 == nil) || (pNode1 != nil && pNode1.Val <= pNode2.Val) {
			p.Next = &ListNode{
				Val: pNode1.Val,
			}
			p = p.Next
			pNode1 = pNode1.Next
		} else if pNode2 != nil {
			p.Next = &ListNode{
				Val: pNode2.Val,
			}
			p = p.Next
			pNode2 = pNode2.Next
		}

		if pNode1 == nil && pNode2 == nil {
			break
		}
	}

	return ret.Next
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list3 := mergeTwoLists2(list1, list2)
	for v := list3; v != nil; v = v.Next {
		fmt.Println(v.Val)
	}
}
