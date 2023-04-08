package main

import (
	"fmt"
	"math"
)

func maxSum(arr []int32, n int, k int) int32 {
	maxSum := int32(math.MinInt32)
	for i := 0; i < n-k+1; i++ {
		currentSum := int32(0)
		for j := 0; j < k; j++ {
			currentSum += arr[i+j]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	arr := []int32{1, 2, 3, 4, 5, 6}
	fmt.Println(maxSum(arr, len(arr), 2))
}
