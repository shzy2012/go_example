package main

import "fmt"

type ByteCounter int

func main() {
	bc := ByteCounter(3)
	i := int(10)
	s := string("34")
	fmt.Printf("%v - %v -%s  \n", bc, i, s)
}
