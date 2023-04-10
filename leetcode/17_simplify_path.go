package main

import "fmt"

type StackS []string

func (s *StackS) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StackS) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *StackS) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func simplifyPath(path string) string {
	s := new(StackS)
	tmp := ""
	for _, c := range path {
		if c == '/' {
			if tmp == "." {
				tmp = ""
				continue
			} else if tmp == ".." {
				s.Pop()
			} else if tmp != "" {
				s.Push(tmp)
			}

			tmp = ""
			continue
		} else {
			tmp = tmp + string(c)
		}
	}

	if tmp == ".." {
		s.Pop()
	} else if tmp != "" && tmp != "." {
		s.Push(tmp)
	}

	ret := ""
	for !s.IsEmpty() {
		if name, ok := s.Pop(); ok {
			ret = "/" + name + ret
		}

	}

	if ret == "" {
		return "/"
	}
	return ret
}

func main() {

	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}
