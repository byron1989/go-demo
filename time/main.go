// time project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ts := time.Now().Unix()
	tStr := time.Unix(ts, 0).Format("2006-01-02 15")

	fmt.Println("Hello World!", tStr)
}
