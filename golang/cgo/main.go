package main

// CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux" CXX="zig c++ -target x86_64-linux" go build -v
// CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="zig cc -target x86_64-windows" CXX="zig c++ -target x86_64-windows" go build -v

// int add(int x, int y){
//     return x + y;
// }
import "C"

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := 0
	y := 0
	if len(os.Args) > 1 {
		x, _ = strconv.Atoi(os.Args[1])
	}
	if len(os.Args) > 2 {
		y, _ = strconv.Atoi(os.Args[2])
	}
	z := C.add(C.int(x), C.int(y))

	fmt.Printf("%d\n", int(z))
}
