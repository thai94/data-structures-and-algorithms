package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	var ret *ListNode
	var pre *ListNode
	var prepre *ListNode
	var idx int
	for p := head; p != nil; p = p.Next {
		if idx%2 != 0 {
			temp := p
			pre.Next = temp.Next
			temp.Next = pre
			if prepre != nil {
				prepre.Next = temp
			}

			if idx == 1 {
				ret = temp
			}
			p = pre
			prepre = pre
		}
		pre = p
		idx++
	}
	if ret != nil {
		return ret
	}
	return head

}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	swapPairs(head)
}
