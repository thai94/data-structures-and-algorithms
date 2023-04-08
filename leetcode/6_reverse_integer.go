package main

import "fmt"

// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	sign := 1
	if x < 0 {
		x = x * -1
		sign = -1
	}

	ret := x % 10
	x = x / 10

	for x > 0 {
		dv := x % 10
		x = x / 10
		ret = ret*10 + dv
	}

	if ret*sign > 2147483647 {
		return 0
	}

	if ret*sign < -2147483648 {
		return 0
	}

	return ret * sign
}

func main() {
	fmt.Println(reverse(0))
}
