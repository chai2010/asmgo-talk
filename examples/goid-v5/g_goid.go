// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"reflect"
	"unsafe"
)

var g_goid_offset uintptr = func() uintptr {
	g := GetGroutine()
	if f, ok := reflect.TypeOf(g).FieldByName("goid"); ok {
		return f.Offset
	}
	panic("can not find g.goid field")
}()

func GetGroutineId() int64 {
	g := GetGroutinePointer()
	p := (*int64)(unsafe.Pointer(g + g_goid_offset))
	return *p
}
