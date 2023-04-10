package main

import "fmt"

func genBase() string {
	return "1"
}

func genNext(str string) string {
	var ret string

	preC := str[0]
	count := 1
	len := len(str)
	for i := 1; i < len; i++ {
		c := str[i]
		if c != preC {
			ret = ret + fmt.Sprintf("%d", count) + string(preC)
			count = 1
			preC = c
		} else {
			count++
		}
	}
	ret = ret + fmt.Sprintf("%d", count) + string(preC)
	return ret
}

func countAndSay(n int) string {
	pre := genBase()
	for i := 1; i < n; i++ {
		pre = genNext(pre)
	}
	return pre
}

func main() {
	fmt.Println(countAndSay(4))
}
