package main

// http://xargin.com/go-and-plan9-asm/#stacks

func main() {
	println(sum(100))
}

//go:noinline
func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return sum(n-1) + n
}
