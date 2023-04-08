package main

import "fmt"

func maxSum(arr []int32, n int, k int) int32 {

	var maxSum int32
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	for i := 1; i < n-k+1; i++ {
		currentSum := maxSum - arr[i-1] + arr[i+k-1]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	arr := []int32{1, 2, 3, 4, 5, 6}
	fmt.Println(maxSum(arr, len(arr), 3))
}
