package main

import (
	"regexp"
	"runtime"
)

var (
	reInit    = regexp.MustCompile(`init·\d+$`) // main.init·1
	reClosure = regexp.MustCompile(`func·\d+$`) // main.func·001
)

// caller types:
// runtime.goexit
// runtime.main
// main.init
// main.main
// main.init·1 -> main.init
// main.func·001 -> main.func
// code.google.com/p/gettext-go/gettext.TestCallerName
// ...
func PrintCallerName(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		println("", file, line)
	}
	name := runtime.FuncForPC(pc).Name()
	if reInit.MatchString(name) {
		println(reInit.ReplaceAllString(name, "init"), file, line)
	}
	if reClosure.MatchString(name) {
		println(reClosure.ReplaceAllString(name, "func"), file, line)
	}
	println(name, file, line)
}
