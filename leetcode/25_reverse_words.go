package main

import "fmt"

func reverseWords(s string) string {
	size := len(s)
	result := []byte(s)
	w := -1
	for i := 0; i < size; i++ {
		if string(result[i]) == " " {
			if w > -1 && string(result[w]) != " " {
				w++
				result[w] = result[i]
			}
			continue
		}
		w++
		result[w] = result[i]
	}
	if string(result[w]) == " " {
		w = w - 1
	}
	p1 := 0
	p2 := w
	for p1 < p2 {
		tmp := result[p1]
		result[p1] = result[p2]
		result[p2] = tmp
		p1++
		p2--
	}

	start := 0
	end := 0
	for end <= w {
		if string(result[end]) == " " || end == w {
			p1 = start
			p2 = end - 1
			if end == w {
				p2 = end
			}

			for p1 < p2 {
				tmp := result[p1]
				result[p1] = result[p2]
				result[p2] = tmp
				p1++
				p2--
			}
			end = end + 1
			start = end
		} else {
			end++
		}

	}
	return string(result[0 : w+1])
}

func main() {
	fmt.Println(reverseWords("   the sky is    blue   "))
}
