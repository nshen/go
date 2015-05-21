package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutineTest() {

	newDivider("goroutineTest.go")

	//goroutine里执行顺序不确定
	//	i := 1 //i被若干goroutine同时访问
	//	go fmt.Println("from goroutine 1: i = ", i)
	//	go func() {
	//		fmt.Println("from goroutine 2: i = ", i)
	//	}()
	//	fmt.Println("from main goroutine: i = ", i)
	//	i++
	//	time.Sleep(time.Second)

	//-----------------------
	//5秒后执行
	//	setTimeout(func() { fmt.Println("setTimeout") }, time.Second*5)

	//--------------------------
	//channel

	//	ch := make(chan string)
	//	go func() {
	//		ch <- "Hello!"
	//		close(ch)
	//	}()

	//	fmt.Println(<-ch) // "Hello!"
	//	fmt.Println(<-ch) // ""
	//	fmt.Println(<-ch) // ""
	//	fmt.Println(<-ch) // ""
	//	v, ok := <-ch
	//	fmt.Println(v, ok) // "" false

	//---------------------
	//遍历 channel
	ch := Producer()
	for s := range ch {
		fmt.Println("aaa", s)
	}

	time.Sleep(10 * time.Second) //main goroutine等待10秒后结束
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

func mutexTest() {
	m := make(map[int]string) //共享内存

	m[2] = "First Value"

	var lock sync.Mutex //zero initialization
	lock.Lock()
	go func() {
		m[2] = "Second Value"
	}()
	lock.Unlock()
	lock.Lock()
	v := m[2]
	lock.Unlock()
	fmt.Println(v)
}

func setTimeout(f func(), delay time.Duration) {

	go func() {
		time.Sleep(delay)
		f()
	}()
}
