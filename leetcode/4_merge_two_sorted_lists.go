package main

import "fmt"

//https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	ret := list1
	for pNode2 := list2; pNode2 != nil; pNode2 = pNode2.Next {
		var prev *ListNode
		for pRet := ret; pRet != nil; pRet = pRet.Next {
			if pRet.Val >= pNode2.Val {
				break
			}
			prev = pRet
		}

		if prev != nil {
			tmp := prev.Next
			prev.Next = &ListNode{
				Val:  pNode2.Val,
				Next: tmp,
			}
		} else {
			tmp := ret
			ret = &ListNode{
				Val:  pNode2.Val,
				Next: tmp,
			}
		}

	}

	return ret
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

	list3 := mergeTwoLists(list1, list2)
	for v := list3; v != nil; v = v.Next {
		fmt.Println(v.Val)
	}
}
