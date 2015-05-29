package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutineTest() {

	newDivider("goroutineTest.go")

	//	goOrderTest() //goroutine执行顺序
	//	setTimeoutTest() //5秒后调用函数
	//	dataRace()
	//	channelTest()
	//	mutexTest()

	//	waitGroupTest()
	//	conditionTest()

	//	selectTest() //select语句
	//	selectTimeout() //select实现timeout
	//	doOnceTest() //多goroutine只执行一次

	//	chanDirectionTest()

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
	//		// There was no pending send operation.
	//	}

	time.Sleep(10 * time.Second) //main goroutine等待10秒后结束
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

	fmt.Println(<-ch) // "Hello!"
	fmt.Println(<-ch) // ""  关闭的channel会读出zero value
	fmt.Println(<-ch) // ""
	fmt.Println(<-ch) // ""
	v, ok := <-ch
	fmt.Println(v, ok) // "" false

	//---------------------
	//遍历 channel
	ch2 := Producer()
	for s := range ch2 {
		fmt.Println("aaa", s)
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

func setTimeoutTest() {
	//-----------------------
	//5秒后执行
	setTimeout(func() { fmt.Println("setTimeout") }, time.Second*5)
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

func Producer() <-chan string {
	ch := make(chan string)
	go func() {
		ch <- "我的名字呀,"
		ch <- "叫N神~"
		close(ch)
	}()
	return ch
}

func setTimeout(f func(), delay time.Duration) {
	go func() {
		time.Sleep(delay)
		f()
	}()
}
