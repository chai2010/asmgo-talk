// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// var Id int = 9527
DATA  ·Id+0(SB)/8,$9527
GLOBL ·Id(SB),NOPTR,$8

// func GetId() int
TEXT ·GetId(SB), $0-8
	MOVQ ·Id(SB), AX   // read id
	MOVQ AX, ret+0(FP) // return id
	RET
