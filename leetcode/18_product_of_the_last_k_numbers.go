package main

import "fmt"

type ProductOfNumbers struct {
	A []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		A: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num > 0 {
		this.A = append(this.A, num*this.A[len(this.A)-1])
	} else {
		this.A = []int{1}
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.A)
	if k >= n {
		return 0
	} else {
		return this.A[n-1] / this.A[n-k-1]
	}
}

func main() {
	obj := Constructor()
	obj.Add(2)
	obj.Add(1)
	param_2 := obj.GetProduct(2)
	fmt.Println(param_2)
}
