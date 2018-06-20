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
func getg_addr() unsafe.Pointer

var g_goid_offset uintptr = func() uintptr {
	g := getg()
	if f, ok := reflect.TypeOf(g).FieldByName("goid"); ok {
		return f.Offset
	}
	panic("can not find g.goid field")
}()

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
