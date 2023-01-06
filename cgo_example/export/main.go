package main

//#include <hello.h>
import "C"

// 用Go重新实现C函数
func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
