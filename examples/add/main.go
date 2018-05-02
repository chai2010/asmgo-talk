package main

import "unsafe"

func Add(a, b int) int

type Int int

func (a Int) Add(b int) int
func (p *Int) Ptr() uintptr

//go:noinline
func (p *Int) Ptr2() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func main() {
	println(Add(1, 2))
	println(Int(3).Add(4))

	p := new(Int)
	//println("ptr:", p.Ptr())
	println("ptr2:", p.Ptr2())
}
