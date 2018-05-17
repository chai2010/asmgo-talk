package main

func main() {}

// "".func_0x.args_stackmap SRODATA size=8
// 0x0000 01 00 00 00 00 00 00 00
func func_0x()

// "".func_1b.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 02 00 00 00 00
// "".func_1i.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 02 00 00 00 00
// "".func_1p.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 02 00 00 00 01
// "".func_ret_1p.args_stackmap SRODATA size=10
// 0x0000 02 00 00 00 02 00 00 00 00 01
func func_1b(a bool)
func func_1i(a int)
func func_1p(a *int)
func func_ret_1p() (a *int)

// "".func_2b.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 02 00 00 00 00
// "".func_2i.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 04 00 00 00 00
// "".func_1i_1p.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 04 00 00 00 02
func func_2b(a, b bool)
func func_2i(a, b int)
func func_1i_1p(a int, b *int)

// "".func_3b.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 02 00 00 00 00
// "".func_3i.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 06 00 00 00 00
func func_3b(a, b, c bool)
func func_3i(a, b, c int)

// "".func_4b.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 02 00 00 00 00
// "".func_4i.args_stackmap SRODATA size=9
// 0x0000 01 00 00 00 08 00 00 00 00
func func_4b(a, b, c, d bool)
func func_4i(a, b, c, d int)

// "".func_4i_1p.args_stackmap SRODATA size=10
// 0x0000 01 00 00 00 0a 00 00 00 10 00                    .
func func_4i_1p(a, b, c, d int, e *int)

/*
// $(GOROOT)/src/cmd/compile/internal/gc/pgen.go

func emitptrargsmap(fn *Node) {
	if fn.funcname() == "_" {
		return
	}
	sym := lookup(fmt.Sprintf("%s.args_stackmap", fn.funcname()))
	lsym := sym.Linksym()

	nptr := int(fn.Type.ArgWidth() / int64(Widthptr))
	bv := bvalloc(int32(nptr) * 2)
	nbitmap := 1
	if fn.Type.NumResults() > 0 {
		nbitmap = 2
	}
	off := duint32(lsym, 0, uint32(nbitmap))
	off = duint32(lsym, off, uint32(bv.n))

	if fn.IsMethod() {
		onebitwalktype1(fn.Type.Recvs(), 0, bv)
	}
	if fn.Type.NumParams() > 0 {
		onebitwalktype1(fn.Type.Params(), 0, bv)
	}
	off = dbvec(lsym, off, bv)

	if fn.Type.NumResults() > 0 {
		onebitwalktype1(fn.Type.Results(), 0, bv)
		off = dbvec(lsym, off, bv)
	}

	ggloblsym(lsym, int32(off), obj.RODATA|obj.LOCAL)
}
*/
