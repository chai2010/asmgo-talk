# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

default:
	-go build -o a.out && ./a.out
	-@rm a.out

asm:
	go tool compile -S main.go > main.s.txt
	-@rm *.o

clean:
	-@rm a.out
