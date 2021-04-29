package main

import (
	"fmt"
)

func main() {

	a := []int{0, 1, 2}
	b := a

	b = append(b, 3)
	a[0] = 100
	// a = append(a, 3)
	b[1] = 101
	a[2] = 102

	fmt.Printf("[a]=>%p:%v,len:%v,cap:%v\n", &a, a, len(a), cap(a))
	fmt.Printf("[b]=>%p:%v,len:%v,cap:%v\n", &b, b, len(b), cap(b))
}
