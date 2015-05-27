package main

import (
	"fmt"
	"time"
)

func timeTest() {
	newDivider("timeTest.go")

	//	timeAfter()
	//	ductionTest()
	//todo: parse

	_, month, day := time.Now().Date() // 年月日
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
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
