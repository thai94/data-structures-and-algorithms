package main

import "fmt"

func removeDuplicates(nums []int) int {
	num := nums[0]
	idx_num := 0
	k := 1
	for idx, v := range nums {
		if v != num {
			nums[idx_num+1] = v
			if idx > idx_num+1 {
				nums[idx] = -1
			}

			num = v
			idx_num = idx_num + 1
			k++
		} else {
			if idx != 0 {
				nums[idx] = -1
			}
		}
	}

	return k
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := removeDuplicates(nums)
	for i := 0; i < k; i++ {
		fmt.Println(nums[i])
	}
}
