package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(123))
	fmt.Println(b)

	i64 := binary.LittleEndian.Uint64(b)
	fmt.Println(i64)
}
