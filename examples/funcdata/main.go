package main

func main() {}

type FuncData struct {
	nbitmap    uint32
	bitvec_len uint32
	bitvec     [1]byte // [nbitmap][len]bit
}

// funcdata 结构:
// nbitmap    uint32; bitmap个数, 只有参数为1, 有返回值的为2
// bitvec.len uint32; 每个bitmap位向量长度, argssize/sizeof(prt) * 2; 如果 nbitmap 为 2, 则后面有 2个位向量
//
// 位向量排列:
// 第一个 bitvec 对应参数部分, 如果是指针则bit为1
// 第二个 bitvec 对应参数+返回值部分, 如果是指针则bit为1
//
// 因此, 第二个 bitvec 可能包含了第一 bitvec 的信息

// 局部变量的 bitvec 可能是无法手工创建的(或者动态创建?)
// 唯一要处理的是stack分裂时, 保证栈上的指针是正确的(gc不会挪动指针)

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
// "".func_ret_2i.args_stackmap SRODATA size=10
// 0x0000 02 00 00 00 04 00 00 00 00 00
// "".func_ret_1i_1p.args_stackmap SRODATA size=10
// 0x0000 02 00 00 00 04 00 00 00 00 02
// "".func_1p_ret_1p.args_stackmap SRODATA size=10
// 0x0000 02 00 00 00 04 00 00 00 01 03
// "".func_1p_ret_1i_1p.args_stackmap SRODATA size=10
// 0x0000 02 00 00 00 06 00 00 00 01 05
func func_2b(a, b bool)
func func_2i(a, b int)
func func_1i_1p(a int, b *int)
func func_ret_2i() (a, b int)
func func_ret_1i_1p() (a int, b *int)
func func_1p_ret_1p(a *int) (b *int)
func func_1p_ret_1i_1p(_ *int) (a int, b *int)

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
// 0x0000 01 00 00 00 0a 00 00 00 10 00
// "".func_4i_1p_4i_1p.args_stackmap SRODATA size=11
// 0x0000 01 00 00 00 14 00 00 00 10 02 00
// "".func_4i_1p_ret_4i_1p.args_stackmap SRODATA size=14
// 0x0000 02 00 00 00 14 00 00 00 10 00 00 10 02 00
// "".func_ret_4i_1p_4i_1p.args_stackmap SRODATA size=14
// 0x0000 02 00 00 00 14 00 00 00 00 00 00 10 02 00
// "".func_1i_ret_3i_1p_4i_1p.args_stackmap SRODATA size=14
// 0x0000 02 00 00 00 14 00 00 00 00 00 00 10 02 00
// "".func_1i_1p_ret_2i_1p_4i_1p.args_stackmap SRODATA size=14
// 0x0000 02 00 00 00 14 00 00 00 02 00 00 12 02 00
func func_4i_1p(a, b, c, d int, e *int)
func func_4i_1p_4i_1p(a, b, c, d int, e *int, g int, h int, i int, j int, k *int)
func func_4i_1p_ret_4i_1p(a, b, c, d int, e *int) (int, int, int, int, *int)
func func_ret_4i_1p_4i_1p() (a, b, c, d int, e *int, g int, h int, i int, j int, k *int)
func func_1i_ret_3i_1p_4i_1p(a int) (b, c, d int, e *int, g int, h int, i int, j int, k *int)
func func_1i_1p_ret_2i_1p_4i_1p(a int, b *int) (c, d int, e *int, g int, h int, i int, j int, k *int)

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
