// structtest project main.go
package main

import (
	"fmt"
)

type Base struct {
	Data string
}

type SubA struct {
	Base
	DataA string
}

type SubB struct {
	Base
	Data  string
	DataB string
}

func main() {
	var a SubA
	var b SubB

	fmt.Println(fmt.Sprintf("%+v", a))
	fmt.Println(fmt.Sprintf("%+v", b))

	a.Data = "Base"
	a.DataA = "DataA"
	b.Data = "Data of B"
	b.Base.Data = "Data of Base"
	b.DataB = "DataB of B"

	fmt.Println(fmt.Sprintf("%+v", a))
	fmt.Println(fmt.Sprintf("%+v", b))
	fmt.Println(fmt.Sprintf("Call b.Data:%s", b.Base.Data))
	fmt.Println(fmt.Sprintf("Ccall b.Base.Data:%s ", b.Base.Data))

}
