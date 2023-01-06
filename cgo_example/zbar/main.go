package main

//#include <zbar.h>
import "C"

func main() {
	var zbarScanner * C.zbar_image_scanner_t
	println(zbarScanner)
}
