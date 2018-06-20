// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://github.com/golang/go/blob/master/src/runtime/runtime2.go
// type g struct { ... }

package main

import (
	"sync"
)

var gls struct {
	m map[int64]map[interface{}]interface{}
	sync.Mutex
}

func init() {
	gls.m = make(map[int64]map[interface{}]interface{})
}

func getMap(goid int64) map[interface{}]interface{} {
	gls.Lock()
	defer gls.Unlock()

	if m, _ := gls.m[goid]; m != nil {
		return m
	}

	m := make(map[interface{}]interface{})
	gls.m[goid] = m
	return m
}

func clearMap(goid int64) {
	gls.Lock()
	defer gls.Unlock()

	delete(gls.m, goid)
}

func GlsMap() map[interface{}]interface{} {
	goid := GetGroutineId()
	return getMap(goid)
}

func GlsClear() {
	goid := GetGroutineId()
	clearMap(goid)
}

func GlsGet(key interface{}) interface{} {
	return GlsMap()[key]
}
func GlsPut(key interface{}, v interface{}) {
	GlsMap()[key] = v
}

func GlsDelete(key interface{}) {
	delete(GlsMap(), key)
}
