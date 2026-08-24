
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("os:", runtime.GOOS)
	fmt.Println("arch:", runtime.GOARCH)
	fmt.Println("version:", runtime.Version())
}
