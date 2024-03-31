package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	maxIdx := m - 1
	if maxIdx < 0 {
		maxIdx = 0
	}
	max := nums1[maxIdx]
	for i2 := n - 1; i2 >= 0; i2-- {
		num2 := nums2[i2]
		if num2 < max {
			for i1 := 0; i1 <= len(nums1)-1; i1++ {
				if num2 <= nums1[i1] {
					for k := len(nums1) - (n - i2); k > i1; k-- {
						nums1[k] = nums1[k-1]
					}
					nums1[i1] = num2
					break
				}
			}
		} else {
			nums1[len(nums1)-(n-i2)] = num2
		}

	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for ; k >= 0; k-- {
		if j < 0 {
			return
		}
		if nums2[j] > nums1[i] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge2(nums1, 3, nums2, 3)
	fmt.Sprintln(nums1)
}
