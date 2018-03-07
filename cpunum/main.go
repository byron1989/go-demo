// cpunum project main.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpu num:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	fmt.Println("Hello World!")
}
