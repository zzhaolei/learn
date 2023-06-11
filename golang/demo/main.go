package main

// #include <stdio.h>
// #include <stdlib.h>
//
// char *myprint(char* s) {
//   return s;
// }
import "C"

import (
	"fmt"
	"unsafe"
)

func test(s string) string {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	ss := C.myprint(cs)

	str := C.GoString(ss)
	return str
}

func main() {
	str1 := test("s1")
	fmt.Println(str1)
	str2 := test("s2")
	fmt.Println("--", str1)
	fmt.Println(str2)
}
