package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := map[[3]int]struct{}{}

	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			break
		}

		j := i + 1
		k := size - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				triplets[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	ret := [][]int{}
	for k := range triplets {
		ret = append(ret, []int{k[0], k[1], k[2]})
	}

	return ret
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	rets := threeSum(nums)
	for _, triplet := range rets {
		fmt.Println(triplet)
	}
}
