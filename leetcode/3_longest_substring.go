package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	lenS := len(s)
	longest := 0
	for i, vi := range s {
		m := map[string]bool{}
		m[string(vi)] = true

		j := 0
		for j = i + 1; j < lenS; j++ {
			_, exist := m[string(s[j])]
			if exist {
				break
			}
			m[string(s[j])] = true
		}

		if j-i > longest {
			longest = j - i
		}
	}

	return longest
}

func main() {
	test1 := "au"
	fmt.Println(lengthOfLongestSubstring(test1))
}
