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

	//	selectTest() //select语句

	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	wg := new(sync.WaitGroup)
	wg.Add(len(people))

	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s’s message.\n", name)
	default:
		// There was no pending send operation.
	}

	time.Sleep(10 * time.Second) //main goroutine等待10秒后结束
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

// RandomBits returns a channel that produces a random sequence of bits.
//--------------------------------------
// select 语句
//
//	检查每个case语句
//	如果有任意一个chan是send or recv read，那么就执行该block
//	如果多个case是ready的，那么随机找1个并执行该block
//	如果都没有ready，那么就block and wait
//	如果有default block，而且其他的case都没有ready，就执行该default block
//---------------------------------------
//  用select实现超时
//
//select {
//case news := <-NewsAgency:
//    fmt.Println(news)
//case <-time.After(time.Minute):
//    fmt.Println("Time out: no news in one minute.")
//}
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
	wg.Wait()
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
