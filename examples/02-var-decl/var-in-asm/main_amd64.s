// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// var debug bool
DATA  ·debug+0(SB)/1,$1
GLOBL ·debug(SB),NOPTR,$1

// utf8: "Hello World!"
// len: 12
DATA  helloworld_data<>+0(SB)/8,$"Hello Wo"
DATA  helloworld_data<>+8(SB)/8,$"rld!"
GLOBL helloworld_data<>(SB),NOPTR,$16

// utf8: "你好, 世界!"
// hex: e4bda0e5a5bd2c20 e4b896e7958c21
// len: 15
DATA  helloworld_zh_data<>+0(SB)/8,$"\xe4\xbd\xa0\xe5\xa5\xbd\x2c\x20"
DATA  helloworld_zh_data<>+8(SB)/8,$"\xe4\xb8\x96\xe7\x95\x8c\x21"
GLOBL helloworld_zh_data<>(SB),NOPTR,$16

// var helloworld string
DATA  ·helloworld+0(SB)/8,$helloworld_data<>(SB)
DATA  ·helloworld+8(SB)/8,$12
GLOBL ·helloworld(SB),NOPTR,$16

// var helloworld_zh string
DATA  ·helloworld_zh+0(SB)/8,$helloworld_zh_data<>(SB)
DATA  ·helloworld_zh+8(SB)/8,$15
GLOBL ·helloworld_zh(SB),NOPTR,$16

// var listen_port int
DATA  ·listen_port+0(SB)/8,$8080
GLOBL ·listen_port(SB),NOPTR,$8

// var point struct{ X, Y int }
DATA  ·point+0(SB)/8,$1234
DATA  ·point+8(SB)/8,$5678
GLOBL ·point(SB),NOPTR,$16

// var point_slice []struct{ X, Y int }

DATA point_slice_data<>+0(SB)/8, $12
DATA point_slice_data<>+8(SB)/8, $34
DATA point_slice_data<>+16(SB)/8, $56
DATA point_slice_data<>+24(SB)/8, $78
GLOBL point_slice_data<>(SB),NOPTR,$32

DATA  ·point_slice+0(SB)/8,$point_slice_data<>(SB)
DATA  ·point_slice+8(SB)/8,$2
GLOBL ·point_slice(SB),NOPTR,$16

// var const_id int
DATA  ·const_id+0(SB)/8,$9527
GLOBL ·const_id(SB),NOPTR|RODATA,$8
