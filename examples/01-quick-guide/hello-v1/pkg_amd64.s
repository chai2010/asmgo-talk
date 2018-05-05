// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// var helloworld string = "Hello World!"
GLOBL ·helloworld(SB),NOPTR,$32               // var helloworld string
DATA  ·helloworld+0(SB)/8,$·helloworld+16(SB) // reflect.StringHeader.Data
DATA  ·helloworld+8(SB)/8,$12                 // reflect.StringHeader.Len
DATA  ·helloworld+16(SB)/8,$"Hello Wo"        // ...string data...
DATA  ·helloworld+24(SB)/8,$"rld!"            // ...string data...

// func HelloWorld()
TEXT ·HelloWorld(SB), $16-0
	MOVQ ·helloworld+0(SB), AX; MOVQ AX, 0(SP)
	MOVQ ·helloworld+8(SB), BX; MOVQ BX, 8(SP)
	CALL runtime·printstring(SB)
	CALL runtime·printnl(SB)
	RET
