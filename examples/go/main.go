package pkg

//go:noinline
func myprint(x, y int) {
	println(x)
}

//go:noinline
func gofun() {
	myprint(123, 4)
	go myprint(567, 8)
}
