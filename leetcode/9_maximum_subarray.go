// https://leetcode.com/problems/maximum-subarray/description/
package main

import "fmt"

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	var curSum int
	for i := 0; i < len(nums); i++ {
		if curSum < 0 {
			curSum = 0
		}
		curSum += nums[i]
		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
