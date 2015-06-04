package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func goroutineTest() {

	newDivider("goroutineTest.go")

	//	goOrderTest() //goroutine执行顺序
	//	dataRace()
	//	atomicTest() //原子操作
	//	mutexTest() //原子锁
	testme()
	//	setTimeoutTest() //用goroutine实现n秒后调用函数

	//	channelTest()
	//	chanDirectionTest() //只进只出的channel

	//	waitGroupTest() //group中全部执行完,再继续
	//	conditionTest()

	//	selectTest() //select语句
	//	selectTimeout() //select实现timeout
	//	doOnceTest() //多goroutine只执行一次

	//	workerPool()

	//	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	//	match := make(chan string, 1) // Make room for one unmatched send.
	//	wg := new(sync.WaitGroup)
	//	wg.Add(len(people))

	//	for _, name := range people {
	//		go Seek(name, match, wg)
	//	}
	//	wg.Wait()
	//	select {
	//	case name := <-match:
	//		fmt.Printf("No one received %s’s message.\n", name)
	//	default:
	//		// There  was no pending send operation.
	//	}

	time.Sleep(10 * time.Second) //main goroutine等待10秒后结束
}

func testme() {
	//顺序执行

	myfunc := func(obj interface{}) {
		fmt.Println("顺序执行 ", obj)
	}
	myfuncChan := make(chan interface{})
	//	outchan := make(chan interface{})

	go func() {
		for {
			myi := <-myfuncChan
			myfunc(myi)
		}
	}()

	for i := 0; i < 100; i++ {
		go func(myi int) {
			myfuncChan <- myi
			runtime.Gosched()
		}(i)
	}

	for j := 200; j < 300; j++ {
		go func(myj int) {
			myfuncChan <- myj
			runtime.Gosched()
		}(j)
	}

}

//多个goroutine执行顺序是不确定的
func goOrderTest() {
	i := 1 //i被若干goroutine同时访问

	go fmt.Println("from g1: i = ", i)
	go func() {
		fmt.Println("from g2: i = ", i)
	}()

	fmt.Println("from main goroutine: i = ", i)
	i++
}

/*
   data race

   多个goroutine读写共享内存数据的时候会出现这种情况,应用锁或channel解决

   名言:Don’t communicate by sharing memory; share memory by communicating
   http://tip.golang.org/doc/articles/race_detector.html
*/
func dataRace() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // one access:read ,increment,write
		close(wait)
	}()
	n++    //another conflicting access
	<-wait //等待goroutine执行完毕

	fmt.Println("now n:", n)
	//此时由于两个goroutine都对n进行了读写,造成data race,所以此时不能确定n的值是几
}

func atomicTest() {

	var ops uint64 = 0

	for i := 0; i < 50; i++ { // 50个goroutine同时for 1秒钟可以加多少个1
		go func() {
			for { //如果不加for,一定是50次
				atomic.AddUint64(&ops, 1)

				runtime.Gosched() // 出让cpu
			}

		}()

	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)

}

func workerPool() {

	/*
			用3个gorutine去执行9个任务

		 		----jobs-------    g1
		main 	 				   g2
				----results----	   g3

	*/

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for g := 0; g < 3; g++ { //开3个goroutine
		go func(id int, jobChan <-chan int, resultChan chan<- int) {
			for j := range jobChan { //用for range遍历channel,channel必须close
				time.Sleep(time.Second) //模拟复杂任务,执行1秒
				resultChan <- j * 2
			}
		}(g, jobs, results)
	}

	t1 := time.Now() //计时开始

	for i := 0; i < 9; i++ {
		jobs <- i
	}
	close(jobs) //不关闭会死锁

	for j := 0; j < 9; j++ {
		fmt.Println(j, <-results)
	}

	fmt.Println("jobs done ", time.Now().Sub(t1)) // 9个1秒任务,用3个worker同时执行,共用3秒
}

func chanDirectionTest() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func ping(pings chan<- string, msg string) {
	//pings是只进channel
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) { //pings是只出channel
	msg := <-pings
	pongs <- msg
}

//TODO: 寻找更好的例子
func conditionTest() {
	m := make(map[int]string)
	m[2] = "First Value"
	var mutex sync.Mutex
	cv := sync.NewCond(&mutex)

	updateCompleted := false
	go func() {
		cv.L.Lock()
		m[2] = "Second Value"
		updateCompleted = true
		cv.Signal()
		cv.L.Unlock()
	}()
	cv.L.Lock()
	for !updateCompleted {
		fmt.Println("for")
		cv.Wait()
		fmt.Println("for2")
	}
	v := m[2]
	cv.L.Unlock()
	fmt.Printf(v)
}

func doOnceTest() {
	var once sync.Once
	done := make(chan bool)

	someFunc := func() {
		fmt.Println("do some func")
	}
	for i := 0; i < 10; i++ {
		go func() {
			//someFunc() //直接调用
			once.Do(someFunc)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		fmt.Println("###", name)
		// Wait for someone to receive my message.
	}
	wg.Done()
}
func selectTest() {
	ch := RandomBits()
	for i := 100; i > 0; i-- {
		fmt.Print(<-ch)
	}
}

func selectTimeout() {

	ch1 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "result 1"
	}()

	select {
	case <-ch1:
		fmt.Println("result1")
	case <-time.After(2 * time.Second):
		fmt.Println("result1 timeout")
	}
	fmt.Println("---")
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "result 2"
	}()
	select {
	case <-ch2:
		fmt.Println("result 2")
	case <-time.After(2 * time.Second):
		fmt.Println("result 2 timeout")
		//	default:
		//		fmt.Println("default可以实现非阻塞立即执行")
	}

}

// RandomBits returns a channel that produces a random sequence of bits.
//--------------------------------------
// select 语句
//
//	检查每个case语句
//	如果有任意一个chan是send or recv read，那么就执行该block
//	如果多个case是ready的，那么随机找1个并执行该block
//	如果都没有ready，那么就block and wait
//	如果有default block，而且其他的case都没有ready，就执行该default block
//----------------------------------------------
func RandomBits() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0: // note: no statement
			case ch <- 1:
			}
		}
	}()
	return ch
}
func waitGroupTest() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Print(n)
			wg.Done()
		}(i)
	}
	wg.Wait() //只有wg所有都done了,才往下执行
	fmt.Println("\nwg all done")
}

func mutexTest() {
	m := make(map[int]string) //共享内存

	m[2] = "First Value"

	var lock sync.Mutex //zero initialization

	go func() {
		lock.Lock()
		m[2] = "Second Value"
		lock.Unlock()
	}()

	lock.Lock()
	v := m[2]
	lock.Unlock()
	fmt.Println(v)
}
func channelTest() {
	//--------------------------
	//channel

	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()

	fmt.Println(<-ch)  // "Hello!"
	fmt.Println(<-ch)  // ""  关闭的channel会读出zero value
	fmt.Println(<-ch)  // ""
	fmt.Println(<-ch)  // ""
	v, ok := <-ch      //ok会指示是否还有值
	fmt.Println(v, ok) // "" false

	//---------------------
	//遍历 channel
	ch2 := make(chan string)
	go func() {
		ch2 <- "我的名字呀,"
		ch2 <- "叫N神~"
		close(ch2)
	}()
	for s := range ch2 {
		fmt.Println("aaa", s) //由于关闭了ch2,所以这里会执行2次,不然会在第3次执行时阻塞
	}
}

func setTimeoutTest() {
	//5秒后执行
	setTimeout(func() { fmt.Println("setTimeout") }, time.Second*5)
}

//更好的实现方式要看 timeTest.go
func setTimeout(f func(), delay time.Duration) {
	go func() {
		time.Sleep(delay)
		f()
	}()
}
