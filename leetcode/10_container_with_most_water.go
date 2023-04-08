// https://leetcode.com/problems/container-with-most-water/
package main

import "fmt"

func min(h1 int, h2 int) int {
	if h1 < h2 {
		return h1
	}
	return h2
}

func maxArea(height []int) int {
	var maxArea int
	left := 0
	right := len(height) - 1
	for left < right {

		curArea := (right - left) * min(height[left], height[right])
		if curArea > maxArea {
			maxArea = curArea
		}
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h))
}
