// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	HelloWorld()
	fmt.Println(helloworld)

	fmt.Println("ptr0:", uintptr(unsafe.Pointer(&helloworld)))

	//fmt.Println("ptr:", helloworld.Ptr())

	// asm lex: (*)
	// https://zh.wikipedia.org/wiki/%E6%98%9F%E8%99%9F
	// https://zh.wikipedia.org/wiki/%E6%8B%AC%E5%8F%B7
	//
	// http://xahlee.info/comp/unicode_full-width_chars.html
	//
	// cmd/link/internal/objfile/objfile.go
	// readSymName reads a symbol name, replacing all "". with pkg.
	// func (r *objReader) readSymName() string
	//
	// 2217
	for i, r := range "（∗*）（）" {
		fmt.Printf("%d: 0x%X, %d\n", i, int(r), int(r))
	}
}

/*
switch ch {
  	case '_': // Underscore; traditional.
  		return true
  	case '\u00B7': // Represents the period in runtime.exit. U+00B7 '·' middle dot
  		return true
  	case '\u2215': // Represents the slash in runtime/debug.setGCPercent. U+2215 '∕' division slash
  		return true
  	}
}
*/
