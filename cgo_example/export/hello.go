// hello.go
package main

import "C"

import "fmt"

// 通过CGO的//export SayHello指令将Go语言实现的函数SayHello导出为C语言函数

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}
