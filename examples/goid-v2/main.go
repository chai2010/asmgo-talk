// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/golang/go/blob/master/src/runtime/runtime2.go
// type g struct { ... }, go1.10 180 bytes

// +build go1.9

package main

import (
	"fmt"
	"sync"
	"unsafe"
)

const g_goid_offset = 152

func getg_addr() unsafe.Pointer

func get_g_id() int64 {
	g := getg_addr()
	p := (*int64)(unsafe.Pointer(uintptr(g) + g_goid_offset))
	return *p
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			fmt.Printf("goroutine(%d): id = %d\n", idx, get_g_id())
		}(i)
	}
	wg.Wait()
}
