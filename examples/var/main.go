package main

import (
	"fmt"
)

var (
	helloworld    string
	helloworld_zh string
	listen_port   int
	point         struct{ X, Y int }
)

func main() {
	fmt.Println(helloworld)
	fmt.Println(helloworld_zh)
	fmt.Println(listen_port)
	fmt.Println(point)
}
