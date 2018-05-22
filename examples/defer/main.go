package pkg

//go:noinline
func myprint(x, y int) {
	println(x)
}

//go:noinline
func gofun() {
	defer myprint(123, 4)
	myprint(567, 8)
}
