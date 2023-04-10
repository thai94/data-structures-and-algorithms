// https://leetcode.com/problems/valid-parentheses/
package main

import "fmt"

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str rune) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isPair(c1, c2 rune) bool {
	if c1 == '(' && c2 == ')' {
		return true
	}

	if c1 == '{' && c2 == '}' {
		return true
	}

	if c1 == '[' && c2 == ']' {
		return true
	}
	return false
}

func isValid(s string) bool {

	stack := new(Stack)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.Push(v)
		} else {
			c, ok := stack.Pop()
			if !ok {
				return false
			}
			if !isPair(c, v) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func main() {
	fmt.Println(isValid("()[]{})"))
}
