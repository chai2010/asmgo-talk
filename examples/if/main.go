package main

// http://xargin.com/go-and-plan9-asm/#stacks

func main() {
	println("If0:", If0(true, 1, 2))
	println("If0:", If0(false, 1, 2))

	println("If1:", If1(true, 1, 2))
	println("If1:", If1(false, 1, 2))

	println("If:", If(true, 1, 2))
	println("If:", If(false, 1, 2))
}

func If(ok bool, a, b int) int

func If0(ok bool, a, b int) int {
	if ok {
		return a
	} else {
		return b
	}
}

func If1(ok bool, a, b int) int {
	if !ok {
		goto L
	}
	return a
L:
	return b
}
