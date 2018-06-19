package main

import (
	"fmt"
)

type MyInterface interface {
	Print()
}

type MyStruct struct{}

func (ms MyStruct) Print() {}

type BigInt struct {
	a int
	b int
	c int
}

//go:noinline
func main() {
	var x = BigInt{a: 12}
	var y interface{} = &x
	var s MyStruct
	var t MyInterface = s
	fmt.Println(y, t)
}
