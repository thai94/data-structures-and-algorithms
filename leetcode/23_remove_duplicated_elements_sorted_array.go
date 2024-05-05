package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

func removeDuplicates2(nums []int) int {
	p1 := 0
	p2 := 0
	for p2 < len(nums) {
		if nums[p1] == nums[p2] {
			p2++
			continue
		}
		p1++
		nums[p1] = nums[p2]
		p2++
	}
	return p1 + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates2(nums)
	fmt.Println(k)
}
