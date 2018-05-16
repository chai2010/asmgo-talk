package main

import (
	"fmt"
	"math/rand"
)

func foo() {
	var a = bar()
	println(a)

	var b = 2
	fmt.Println(b)

	fmt.Sprint()

	fmt.Sprintf("")

}

func bar() int {
	return rand.Int()
}
