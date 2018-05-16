package main

// http://xargin.com/go-and-plan9-asm/#stacks

func main() {
	println("0:", sum(0))
	println("1:", sum(1))
	println("2:", sum(2))
	println("3:", sum(3))

	println(sum0(100))
	println(sum1(100))
}

func sum(n int) int

//go:noinline
func sum0(n int) int {
	if n <= 0 {
		return 0
	}
	return sum0(n-1) + n
}

func sum1(n int) (result int) {
	var temp1 int
	var temp2 int

	if n <= 0 {
		goto L_END
	}
	goto L_STEP_TO_END

L_STEP_TO_END:
	temp1 = n - 1
	temp2 = sum1(temp1)
	result = n + temp2
	return result

L_END:
	return 0
}
