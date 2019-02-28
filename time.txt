package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("now=>", now)
	//时间戳
	fmt.Println("unix=>", now.Unix())
	fmt.Println("unixNano=>", now.UnixNano())
}
