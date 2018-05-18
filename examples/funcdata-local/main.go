package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	test := "{\"endpoint\":\"b\"}"
	data := map[string]interface{}{"a": test}
	v := jsonutilToString(data)
	data2 := map[string]interface{}{"a": v}
	v2 := jsonutilToString(data2)

	fmt.Println(v2)
}

// gclocals·568470801006e5c0dc3947ea998fe279 SRODATA dupok size=10
//	0x0000 02 00 00 00 02 00 00 00 00 02
func foo() {
	m := make(map[int]string)
	fmt.Println(m)
}

// gclocals·80b0d9ad63c242a848af3702a9ec7107 SRODATA dupok size=11
// 0x0000 03 00 00 00 05 00 00 00 00 01 14
func foo1() {
	m := make(map[int]string)
	p := new(int)
	fmt.Println(m, p)
}

// gclocals·568470801006e5c0dc3947ea998fe279 SRODATA dupok size=10
// 0x0000 02 00 00 00 02 00 00 00 00 02
func foo2() {
	p := new(int)
	fmt.Println(p)
}

func foo3() {
	p := new(int)
	println(p)
}

func Encode(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

func Decode(y []byte, o interface{}) error {
	return json.Unmarshal(y, o)
}

func jsonutilToString(o interface{}) string {
	b, err := Encode(o)
	if err != nil {
		return ""
	}
	return string(b)
}
