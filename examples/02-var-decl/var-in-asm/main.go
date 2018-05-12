// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

var (
	debug         bool
	helloworld    string
	helloworld_zh string
	listen_port   int
	point         struct{ X, Y int }
	point_slice   []struct{ X, Y int }
	const_id      int // read only
	m             map[string]int
	ch            chan int
)

func main() {
	fmt.Println(debug)
	fmt.Println(helloworld)
	fmt.Println(helloworld_zh)
	fmt.Println(listen_port)
	fmt.Println(point)
	fmt.Println(point_slice)
	fmt.Println(const_id)
	fmt.Println(m)
	fmt.Printf("%#v\n", ch)

	if false {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic:", err)
			}
		}()

		fmt.Println("set const_id begin")
		const_id = 123 // should panic, can't recover!!!
		fmt.Println("set const_id end")
	}
}
