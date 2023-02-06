package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	longest := 0
	m := map[string]int{}
	start := 0
	for end, vi := range s {
		idx, exist := m[string(vi)]
		if exist && idx >= start {
			start = idx + 1
		}

		count := end - start + 1
		if count > longest {
			longest = count
		}
		m[string(vi)] = end
	}

	return longest
}

func main() {
	test1 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(test1))
}
