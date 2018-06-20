// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/golang/go/blob/master/src/runtime/runtime2.go
// type g struct { ... }

package main

import (
	"fmt"
	"sync"
)

func PrintValue(idx int) {
	fmt.Println("PrintValue:", idx, GlsGet(idx))
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			defer GlsClear()
			defer PrintValue(idx)

			GlsPut(idx, idx)

			fmt.Printf("goroutine(%d): id = %d\n", idx, GetGroutineId())
		}(i)
	}
	wg.Wait()
}
