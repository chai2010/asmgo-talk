package main

/*
"".main0.args_stackmap SRODATA size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
"".main1.args_stackmap SRODATA size=9
        0x0000 01 00 00 00 02 00 00 00 00                       .........
"".main2.args_stackmap SRODATA size=9
        0x0000 01 00 00 00 04 00 00 00 00                       .........
"".main3.args_stackmap SRODATA size=9
        0x0000 01 00 00 00 06 00 00 00 00                       .........
"".main4.args_stackmap SRODATA size=9
        0x0000 01 00 00 00 08 00 00 00 00                       .........
"".main5.args_stackmap SRODATA size=10
        0x0000 01 00 00 00 0a 00 00 00 10 00                    ..........
*/

func main0()
func main1(a int)
func main2(a, b int)
func main3(a, b, c int)
func main4(a, b, c, d int)
func main5(a, b, c, d int, e *int)

func main() {
	main2(1, 2)
}
