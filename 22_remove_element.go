package main

import "fmt"

// https://leetcode.com/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150

func removeElement(nums []int, val int) int {
	p1 := 0
	p2 := len(nums) - 1
	for p1 <= p2 {
		if nums[p1] != val {
			p1++
			continue
		}
		if nums[p2] == val {
			p2--
			continue
		}
		k := nums[p1]
		nums[p1] = nums[p2]
		nums[p2] = k
	}
	return p1
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(nums, 2)
	fmt.Println(k)
}
