// +build ignore

package main

func main() {}

func fn_local() {
	a := 1
	b := new(int)
	c := 3

	println(a, b, c)
	fn_noinline(b)

	func(q int, aa *int) *int {
		println(a, *aa)
		return &a
	}(a, b)
}

//go:noinline
func fn_noinline(a *int) {}
