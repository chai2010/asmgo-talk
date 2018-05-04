package main

import (
	"fmt"
)

var (
	helloworld    string
	helloworld_zh string
	listen_port   int
	point         struct{ X, Y int }
	const_id      int // read only
)

func main() {
	fmt.Println(helloworld)
	fmt.Println(helloworld_zh)
	fmt.Println(listen_port)
	fmt.Println(point)
	fmt.Println(const_id)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()

	fmt.Println("set const_id begin")
	const_id = 123 // should panic, can't recover!!!
	fmt.Println("set const_id end")
}
