// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/golang/go/blob/master/src/runtime/runtime2.go
// type g struct { ... }, go1.10 180 bytes

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
func get_tls() (tls, g unsafe.Pointer)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			g := getg()
			gid := reflect.ValueOf(g).FieldByName("goid").Int()

			tls, gp := get_tls()

			fmt.Printf("goroutine(%d): id = %d, tls = %d, g = %p\n", idx, gid, uintptr(tls), gp)
		}(i)
	}
	wg.Wait()
}

//#ifdef GOARCH_amd64
//#define	get_tls(r)	MOVQ TLS, r // 取 TLS 指向内存的值, 和 goid 有相关性, 类似 x
//#define	g(r)	0(r)(TLS*1)     // 这里的 TLS 是TLS符号本身的地址, 类似 &x
//#endif
