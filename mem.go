package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

/**

GO 内存模型


为优化的考虑 compilers 和 processors 会重排对变量读写的执行顺序 , 导致代码的书写顺序和执行顺序不一致.
所以
要研究内存模型
Go内存模型规范了在什么条件下,一个Goroutine对某个变量的修改一定对其它Goroutine可见。

Go的并发模型是基于CSP（Communicating Sequential Process）


在同一个goroutine里,变量读写顺序与代码书写顺序一致, 因为只有在不影响逻辑顺序的情况下,才会对变量读写进行重排,
但正由于这种重排,导致在其他goroutine里观察到的执行顺序会不一致,例如在一个goroutine里执行a = 1; b = 2;在另一个goroutine里观察到的可能是b先于a执行

为了明确表达局部内存执行顺序,定义了happens before

happens before

If event e1 happens before event e2, then we say that e2 happens after e1. Also, if e1 does not happen before e2 and does not happen after e2, then we say that e1 and e2 happen concurrently.

v = 共享变量

什么时候对v的读操作能观察到对v的写操作?

1 读 not happens before 写 (读在写之后,或者与写并行)
2 写和读之间不包含其他对v的写操作

什么时候对v的读操作一定能观察到对v的写操作?

1. 写 happens before r
2. 任何其他对v的写操作,都只在当前的写之前或读之后发生


但在多个Goroutine里如果要访问一个共享变量，我们就必须使用同步工具(synchronization events )来建立happens-before条件，来保证对该变量的读操作能读到期望的修改值。

要保证并行执行体对共享变量的顺序访问方法就是用锁


Synchronization

	初始化
	If a package p imports package q, the completion of q's init functions happens before the start of any of p's.
	The start of the function main.main happens after all init functions have finished.

	Goroutine的创建

		var a string
		func f() {
			print(a)
		}
		func hello() {
			a = "hello, world"
			go f() //在a设置后创建,保证输出hello, world
		}

	Goroutine的销毁

		var a string
		func hello() {
			go func() { a = "hello" }()
			print(a) //并发,不能保证看到a被赋值, 可能输出为空字符串
		}

	为了使对a的修改可见,应该使用锁或者go channel进行同步

Channel同步

	channel是不同goroutine之间同步的主要手段,对channel的send操作,一般都对应一个在其他goroutine中的receive操作

规则:
  *在有缓冲的channel中,发送在接收之前发生
	A send on a channel happens before the corresponding receive from that channel completes.

  *在无缓冲的channel中,接收在发送之前发生
	A receive from an unbuffered channel happens before the send on that channel completes.

下边这段代码,channel有没有缓冲得到的结果完全不同

	var a string = "null"
	m := make(chan int) //无缓冲,receive先执行
	//m := make(chan int,10)//有缓冲,send先执行

	go func() {
		a = "hello, world"
		<-m
	}()

	m <- 1
	print(a) //如果channel有缓冲,则send先执行,所以先print输出null,
			 //如果channel无缓冲(send不进去,要有receive配对才行),则receive先执行,所以a被赋值,最后输出hello,world


下边的技巧,限制最多3个工作同时进行

var limit = make(chan int, 3)
func main() {
	for _, w := range work {
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
	select{}
}

两种锁

sync.Mutex
sync.RWMutex

sync.Mutex -> java.util.concurrent.ReentrantLock
sync.RWMutex -> java.util.concurrent.locks.ReadWriteLock

规则 For any sync.Mutex or sync.RWMutex variable l and n < m, call n of l.Unlock() happens before call m of l.Lock() returns.

var l sync.Mutex
func lockFunc() {
	a = "hello, world"
	l.Unlock()
}
func lockTest() {
	l.Lock()
	go lockFunc() // unlock一定会在下一句lock之前执行
	l.Lock()
	print(a) //保证输出 hello, world
}


RWMutex...

For any call to l.RLock on a sync.RWMutex variable l, there is an n such that the l.RLock happens (returns) after call n to l.Unlock and the matching l.RUnlock happens before call n+1 to l.Lock.


Once

	var a string
	var once sync.Once

	func setup() {
		a = "hello, world"
	}

	func doprint() {
		once.Do(setup)
		print(a)
	}

	func twoprint() {
		go doprint()
		go doprint()
	}

calling twoprint causes "hello, world" to be printed twice. The first call to doprint runs setup once.




*/

/*
参考
http://golang.org/ref/mem
http://hugozhu.myalert.info/2013/04/20/31-golang-memory-model.html
*/

func printfunc() {
	fmt.Println(a)
}

func creation() {
	a = "hello, world"
	go printfunc()
}

var a string = "null"

func destruction() {

	m := make(chan int)
	//m := make(chan int,10)

	go func() {
		a = "hello, world"
		<-m
	}()

	m <- 1
	print(a)
}

var l sync.Mutex

func lockFunc() {
	a = "hello, world"
	l.Unlock()
}
func lockTest() {
	l.Lock()
	go lockFunc() // unlock一定会在下一句lock之前执行
	l.Lock()
	print(a)
}

var aa, bb int

func abfunc() {
	bb = 2
	aa = 1
}

func abg() {
	print(bb)
	print(aa)
}

func abTest() {
	go abfunc()
	abg()
}

func mem() {
	//	destruction()
	//	lockTest()
	abTest()
	bufio.NewScanner(os.Stdin).Scan()
}
