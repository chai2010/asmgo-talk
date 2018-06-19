// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/golang/go/blob/master/src/runtime/runtime2.go
// type g struct { ... }

package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

func getg() interface{}
func getg_type() unsafe.Pointer
func getg_addr() unsafe.Pointer

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			g := getg()
			gid := reflect.ValueOf(g).FieldByName("goid").Int()

			fmt.Printf("goroutine(%d) id: %d\n", idx, gid)
		}(i)
	}
	wg.Wait()
}
