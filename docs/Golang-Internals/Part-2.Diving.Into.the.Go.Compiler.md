# 深入浅出 Golang，第二部分：深入虎穴，一探编译器 [Part 2: Diving Into the Go Compiler][1]

你是否知道我们是怎么通过Golang的运行时，让[interface][2]去引用到一个变量的？这其实是一个非常有深度的问题，因为在Golang里面一个type实现了某一个interface，但type本身没有存储任何信息关联到这个interface，当然我们可以尝试用我们从[Part 1]里面了解到的信息，从go的编译器实现角度来回答这个问题。

为了更精准的回答上面的类似问题，我们接下来更加深入的分析go的编译器：我们写一个非常小的golang版本的"hello world"，然后通过分析这个"hello world"，了解内部的类型转换等相关机制，通过例子也进一步对node-tree的生成过程做详细解析。当然了解node-tree的生成过程不是我们的最终目的，我们的目的是以此为基础去横向涉猎go编译器的其他特性。

- - -

## 前戏
为了准备了解编译器，我们准备我们的实验环境，我们从golang的编译器直接入手，而不是通过golang的集成工具。可以通过如下的方式使用编译器(注意安装go1.4版本)：

```
go tool 6g test.go
```

如果使用的是go1.5 或者更高的版本可以用如下的命令：

```
go tool compile test.go
```

上面的命令会编译源文件test.go，然后在当前目录生成object文件。在笔者机器上6g是一个amd64架构的编译器，要生成其他架构的代码必须要用相应架构的编译器。
我们直接操作编译器的时候，我们可以手动在编译器上加上一些命令行参数，让编译器给我们生成相关辅助信息，更详细的编译器参数可以参考这个[地方][3]，我们这里会给编译器上架_-W_参数，加上这个参数，编译器会把node-tree打印出来。

## 编写"hello world"程序
首先我们编写一个用于分析的简单go程序，我的版本是这样的：

```
  1  package main
  2
  3  type I interface {
  4          DoSomeWork()
  5  }
  6
  7  type T struct {
  8          a int
  9  }
 10
 11  func (t *T) DoSomeWork() {
 12  }
 13
 14  func main() {
 15          t := &T{}
 16          i := I(t)
 17          print(i)
 18  }
```

上面的示例代码非常简洁，严格意义上也就只有第17行看起来是可以不用的，但在golang里面第17行却也是必须的，因为没有17行就会是unused variable，这在golang里面是编译器错误。把上面的代码保存为test.go, 接下来我们就通过编译器编译这个源文件。
```
go tool 6g -W test.go
```
执行上面的命令会看到编译器打印了源文件里面每一个函数的node-tree，我们这里主要是打印出main函数和init函数，其中init函数编译器自动为每一个包文件加上的内部函数，我们这里直接无视init函数就好。
编译器打印node-tree的时候，会为每一个函数打印两个版本的node-tree，第一个版本是语法此法解析后的原始版本，第二个版本是执行了类型检查，并填充了相关类型信息并做一些必要的修正后的版本。

```
before I.DoSomeWork
.   CALLINTER l(20) tc(1)
.   .   DOTINTER l(20) x(0+0) tc(1) FUNC-method(*struct {}) func()
.   .   .   NAME-main..this u(1) a(true) g(1) l(20) x(0+0) class(PPARAM) f(1) esc(s) tc(1) used(true) main.I
.   .   .   NAME-main.DoSomeWork u(1) a(true) l(20) x(0+0)
after walk I.DoSomeWork
.   CALLINTER u(100) l(20) tc(1)
.   .   DOTINTER u(2) l(20) x(0+0) tc(1) FUNC-method(*struct {}) func()
.   .   .   NAME-main..this u(1) a(true) g(1) l(20) x(0+0) class(PPARAM) f(1) esc(s) tc(1) used(true) main.I
.   .   .   NAME-main.DoSomeWork u(1) a(true) l(20) x(0+0)
```

## 理解main方法的node-tree
我们开始先看一下在编译器修订之前的main方法的node-tree:
```
DCL l(15)
.   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T

AS l(15) colas(1) tc(1)
.   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T
.   PTRLIT l(15) esc(no) ld(1) tc(1) PTR64-*main.T
.   .   STRUCTLIT l(15) tc(1) main.T
.   .   .   TYPE <S> l(15) tc(1) implicit(1) type=PTR64-*main.T PTR64-*main.T

DCL l(16)
.   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I

AS l(16) tc(1)
.   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T
.   NAME-main.t u(1) a(1) g(1) l(15) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) PTR64-*main.T

AS l(16) colas(1) tc(1)
.   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I
.   CONVIFACE l(16) tc(1) main.I
.   .   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T

VARKILL l(16) tc(1)
.   NAME-main.autotmp_0000 u(1) a(1) l(16) x(0+0) class(PAUTO) esc(N) tc(1) used(1) PTR64-*main.T

PRINT l(17) tc(1)
PRINT-list
.   NAME-main.i u(1) a(1) g(2) l(16) x(0+0) class(PAUTO) f(1) ld(1) tc(1) used(1) main.I
```
尝试理解一下上面node-tree。接下来的解释过程，我们会对node-tree做适当的删减，去掉一些意义不大的部分，让他看起来更加简洁精炼。
第一个node非常简单：
```
DCL l(15)
.   NAME-main.t l(15) PTR64-*main.T
```
第一行`DCL l(15)` 声明node，其中`l(15)`的意思是node来自于源文件的第15行。下面的一行`NAME-main.t l(15) PTR64-*main.T`是节点的名字是`main.t`对应到源文件的15行变量`t`，变量t是一个64位的指针，指向类型为`main.T`的变量

接下来的会稍微复杂一点，理解起来也会棘手一些：
```
AS l(15)
.   NAME-main.t l(15) PTR64-*main.T
.   PTRLIT l(15) PTR64-*main.T
.   .   STRUCTLIT l(15) main.T
.   .   .   TYPE l(15) type=PTR64-*main.T PTR64-*main.T
```
第一行`AS l(15)` 说明是一个用于赋值的node, node的一个孩子`NAME-main.t l(15) PTR64-*main.T`是一个具名的节点，代表`main.t`这个变量；第二个孩子`PTRLIT l(15) PTR64-*main.T`是我们用来赋值给`main.t`的节点，这个用来赋值的节点字面`PRTLIT`上是一个取地址操作，相当于`&`取地址符，这个取地址操作的节点也有一个孩子节点`STRUCTLIT l(15) main.T`，这个节点是指向具体的类型`main.T`；
接下来又是一个声明节点：
```
DCL l(16)
.   NAME-main.i l(16) main.I
```
声明了一个类型为`main.I`的变量`main.i`

然后编译器创建了一个临时变量`autotmp_0000`，并且把`main.t` 赋值给它。

```
AS l(16) tc(1)
.   NAME-main.autotmp_0000 l(16) PTR64-*main.T
.   NAME-main.t l(15) PTR64-*main.T
```
好接下来的几个节点是我们真正感兴趣的部分：
```
AS l(16)
.   NAME-main.i l(16)main.I
.   CONVIFACE l(16) main.I
.   .   NAME-main.autotmp_0000 PTR64-*main.T
```
我们看到编译器把一个特殊节点`CONVIFACE`赋值给了`main.i`，这一步有一些意思，但是还是裹了一层薄纱，没有真正看到内涵。要搞清楚内涵信息我们得继续往下看编译器做了修订以后的版本(也就是标注为`after walk main`以后的片段)。

- - -

## 编译器是怎么修订赋值node的
下面我们可以看到编译器修订过程具体做了哪些猫腻：
```
AS-init
.   AS l(16)
.   .   NAME-main.autotmp_0003 l(16) PTR64-*uint8
.   .   NAME-go.itab.*"".T."".I l(16) PTR64-*uint8

.   IF l(16)
.   IF-test
.   .   EQ l(16) bool
.   .   .   NAME-main.autotmp_0003 l(16) PTR64-*uint8
.   .   .   LITERAL-nil I(16) PTR64-*uint8
.   IF-body
.   .   AS l(16)
.   .   .   NAME-main.autotmp_0003 l(16) PTR64-*uint8
.   .   .   CALLFUNC l(16) PTR64-*byte
.   .   .   .   NAME-runtime.typ2Itab l(2) FUNC-funcSTRUCT-(FIELD-
.   .   .   .   .   NAME-runtime.typ·2 l(2) PTR64-*byte, FIELD-
.   .   .   .   .   NAME-runtime.typ2·3 l(2) PTR64-*byte PTR64-*byte, FIELD-
.   .   .   .   .   NAME-runtime.cache·4 l(2) PTR64-*PTR64-*byte PTR64-*PTR64-*byte) PTR64-*byte
.   .   .   CALLFUNC-list
.   .   .   .   AS l(16)
.   .   .   .   .   INDREG-SP l(16) runtime.typ·2 G0 PTR64-*byte
.   .   .   .   .   ADDR l(16) PTR64-*uint8
.   .   .   .   .   .   NAME-type.*"".T l(11) uint8

.   .   .   .   AS l(16)
.   .   .   .   .   INDREG-SP l(16) runtime.typ2·3 G0 PTR64-*byte
.   .   .   .   .   ADDR l(16) PTR64-*uint8
.   .   .   .   .   .   NAME-type."".I l(16) uint8

.   .   .   .   AS l(16)
.   .   .   .   .   INDREG-SP l(16) runtime.cache·4 G0 PTR64-*PTR64-*byte
.   .   .   .   .   ADDR l(16) PTR64-*PTR64-*uint8
.   .   .   .   .   .   NAME-go.itab.*"".T."".I l(16) PTR64-*uint8
AS l(16)
.   NAME-main.i l(16) main.I
.   EFACE l(16) main.I
.   .   NAME-main.autotmp_0003 l(16) PTR64-*uint8
.   .   NAME-main.autotmp_0000 l(16) PTR64-*main.T
```
从截取的输出看到，编译器首先给赋值节点加另一个初始化节点`AS-init`，在初始化节点`AS-init`里面，创建一个新的自动变量`autotmp_0003`，并且赋值为`go.itab.*"".T."".I`，这一步后，就检查这个变量是否为nil`LITERAL-nil`，如果变量不为nil，则调用函数`runtime.typ2Itab`，并传递如下参数：

`NAME-type.*"".T l(11)` 一个指向类型`main.T`的指针
`NAME-type."".I l(16)` 一个指向类型`main.I`的指针
以及`NAME-go.itab.*"".T."".I l(16)` 一个指向`go.itab.*"".T."".I`的变量

从上面不难发现编译器创建了一个临时变量存储`main.T`转到`main.I`的类型转换结果。

## 深入老巢，观察gititab函数
我们先把the-fucking-code:[typ2Itab][4]列出来：
```
func typ2Itab(t *_type, inter *interfacetype, cache **itab) *itab {
	tab := getitab(inter, t, false)
	atomicstorep(unsafe.Pointer(cache), unsafe.Pointer(tab))
	return tab
}
```
擦。。。，上面的代码太简单了，所有事情其实都是给[getitab][5]干了，自己只是把结构存储到了cache里面。好接下来看getitab:
```
m =
    (*itab)(persistentalloc(unsafe.Sizeof(itab{})+uintptr(len(inter.mhdr)-1)*ptrSize, 0,
    &memstats.other_sys))
    m.inter = interm._type = typ

ni := len(inter.mhdr)
nt := len(x.mhdr)
j := 0
for k := 0; k < ni; k++ {
	i := &inter.mhdr[k]
	iname := i.name
	ipkgpath := i.pkgpath
	itype := i._type
	for ; j < nt; j++ {
		t := &x.mhdr[j]
		if t.mtyp == itype && t.name == iname && t.pkgpath == ipkgpath {
			if m != nil {
				*(*unsafe.Pointer)(add(unsafe.Pointer(&m.fun[0]), uintptr(k)*ptrSize)) = t.ifn
			}
		}
	}
}
```
getitab函数非常大，这里只截取了笔者认为最有价值的部分。首先会申请内存存储返回的结果：
```
m =
    (*itab)(persistentalloc(unsafe.Sizeof(itab{})+uintptr(len(inter.mhdr)-1)*ptrSize, 0,
    &memstats.other_sys))
    m.inter = interm._type = typ
```

这里调用的是一个很奇怪的函数申请的内存，为什么会调用这么一个东东，我们得看下`itab`结构体的定义：
```
type itab struct {
	inter  *interfacetype
	_type  *_type
	link   *itab
	bad    int32
	unused int32
	fun    [1]uintptr // variable sized
}
```
结构体的最后一个变量`fun`是一个只有一个元素的数组，这里其实是一个可变长度的函数指针数组，存储了对应到interface定义的所有函数，Go的设计者这里是通过unsafe包提供的能力自己手动管理内存的，所以需要申请的内存大小原来的大小在加上interface定义的方法数减1乘以指针大小`unsafe.Sizeof(itab{})+uintptr(len(inter.mhdr)-1)*ptrSize`。

接下来会看到两个嵌套loop, 第一个loop我们遍历interface的方法，为每个方法尝试找到在type里面对应的方法，用如下的代码判断方法是否相等：
`if t.mtyp == itype && t.name == iname && t.pkgpath == ipkgpath`
如果找到了我们就把函数指针存储到`fun`里面：
`*(*unsafe.Pointer)(add(unsafe.Pointer(&m.fun[0]), uintptr(k)*ptrSize)) = t.ifn`

笔者发现这里的一个优化点：在interfae和type的方法都是按照字母序排序好的情况下，这里的循环其可以O(n+m)的情况下完成，而不需要O(n*m){译者注：这里的n 和m对应到interface和type的方法数，也就意味着n和m都是很小的数，所以这里没太大必要做过度优化去牺牲可阅读性}。

好时光倒转一下，我们继续看前面关于node-tree的解析的最后一部分：
```
AS l(16)
.   NAME-main.i l(16) main.I
.   EFACE l(16) main.I
.   .   NAME-main.autotmp_0003 l(16) PTR64-*uint8
.   .   NAME-main.autotmp_0000 l(16) PTR64-*main.T
```
这里把一个`EFACE l(16)`的节点赋值给了`NAME-main.i`变量，而`EFACE l(16)`包含了一个`autotmp_0003`的引用，而且前面的分析我们也知道`autotmp_0003`是一个指向`itab`结构体的指针，存储了`untime.typ2Itab`返回值；同时`EFACE l(16)`还包含了`autotmp_0000`的引用，而这个变量存储的值就是`main.t`,所以`main.i`就已经关联了一个`itab`和`main.T`，就能访问相关的方法和变量了。也就是`main.i`其实是runtime包里面的`iface`的实例，`iface`结构体定义如下：
```
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```

{译者注：好吧，看到这里，对于学过C++等多态语言的，不难类比到C++的多态机制，只是go的这里是动态多态，他的虚表itab是动态生成的, 同时把虚表单独拿出来和一个对象结合到一起变成一个iface，这样就提供了非常大的灵活性。}


## 接下来路在何方
通过前面的缀叙，我们还仅仅只是覆盖了go编译器和go运行时非常小的一部分，还有非常多的东西可以探讨。在本系列接下来的文章里面我们还会继续探讨：object文件，链接过程，重定向等。



[1]: http://blog.altoros.com/golang-internals-part-2-diving-into-the-go-compiler.html "Part 2: Diving Into the Go Compiler"
[2]: http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go "interface in go"
[3]: https://golang.org/cmd/compile/ "Compiler"
[4]: https://golang.org/src/cmd/compile/internal/gc/builtin/runtime.go?h=typ2Itab#L63 "typ2Itab"
[5]: https://golang.org/src/runtime/iface.go?h=getitab#L22 "getitab"
