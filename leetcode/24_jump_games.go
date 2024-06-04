package main

import "fmt"

// https://leetcode.com/problems/jump-game/?envType=study-plan-v2&envId=top-interview-150

func canJump(nums []int) bool {
	return jump(nums, 0)
}

func jump(nums []int, jumpIdx int) bool {
	if jumpIdx >= len(nums)-1 {
		return true
	}

	if nums[jumpIdx] == 0 {
		return false
	}

	for i := nums[jumpIdx]; i >= 1; i-- {
		reached := jump(nums, jumpIdx+i)
		if reached {
			return true
		}
	}
	return false
}

func canJump2(nums []int) bool {
	furthest_reach := 0
	for i, num := range nums {
		if i > furthest_reach {
			return false
		}
		if i+num > furthest_reach {
			furthest_reach = i + num
		}
	}
	return furthest_reach >= len(nums)-1
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump2(nums))
}
