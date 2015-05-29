package main

import (
	"fmt"
	"time"
)

func timeTest() {
	newDivider("timeTest.go")

	//	timeAfter()
	//	stopTimer()
	//	tickerTest() //隔段时间调用一次
	//	ductionTest()
	//todo: parse

	_, month, day := time.Now().Date() // 年月日
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}

func tickerTest() {
	ticker := time.NewTicker(time.Millisecond * 500) //500毫秒
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second * 3) //3秒后关闭
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func stopTimer() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("2秒了")
	timer1.Reset(time.Second) //reset到1秒后
	<-timer1.C
	fmt.Println("1秒了")
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 experied")
	}()
	stopped := timer2.Stop() //让timer2停止
	if stopped {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(time.Second * 3)
}

func ductionTest() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC) //年,月,日,时,分,秒,纳秒,*Location
	fmt.Printf("Go launched at %s\n", t.Local())
	fmt.Println(time.Now().Sub(t.Local()).Hours()) //duction表示一段时间duction = time - time

	t0 := time.Now()
	//-------------------
	//expensiveCall()
	sum := 0
	for i := 0; i < 99999999; i++ {
		sum += i
	}
	//------------------
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
func timeAfter() {
	timeChan := time.After(3 * time.Second)
	//	var timeChan <-chan time.Time = time.After(3 * time.Second)
	<-timeChan
	fmt.Println("3 seconds")
	time.AfterFunc(2*time.Second, func() { fmt.Println("another 2 seconds") })
	fmt.Println("不会阻塞")

	time.Sleep(99999999999999) //没有sleep AfterFunc就没有机会执行
}
