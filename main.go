package main

import (
	"encoding/binary"
	"fmt"
)

func main(){
	abyte := byte(144)
	asint := int(abyte)
	fmt.Println(asint)


	asint64 :=int64(abyte)
	fmt.Println(asint64)

	int63:= binary.BigEndian.Uint64([]byte("1234344"))
	fmt.Println(int63)
}