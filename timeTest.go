package main

import (
	"fmt"
	"time"
)

func timeTest() {
	newDivider("timeTest.go")

	p := fmt.Println

	now := time.Now()
	p(now)                      //2015-06-09 09:53:30.2621972 +0800 CST
	_, month, day := now.Date() // 年月日
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC) //年,月,日,时,分,秒,纳秒,*Location
	p(then)                                                          //2009-11-17 20:34:58.651387237 +0000 UTC
	p(then.Year())
	p(then.Month(), then.Month() == time.November) //November , true
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location()) //UTC
	p(then.Local())    //2009-11-18 04:34:58.651387237 +0800 CST
	//Local returns t with the location set to local time.

	p("---------------")
	p(then.Weekday())   //Tuesday
	p(then.Before(now)) //true
	p(then.After(now))  //false
	p(then.Equal(now))  //false

	p("---------------")
	thenCopy := then                      //值copy
	p(thenCopy)                           //2009-11-17 20:34:58.651387237 +0000 UTC
	thenCopy = thenCopy.AddDate(-1, 0, 0) //年月日
	p(thenCopy)                           //2008-11-17 20:34:58.651387237 +0000 UTC
	p(then)                               //2009-11-17 20:34:58.651387237 +0000 UTC

	p("---------------")
	//加减得到两个时间点的间隙
	var diff time.Duration = now.Sub(then)
	p("diff =", diff)     //48702h11m4.834164063s
	p(diff.Hours())       //48702.2113387928
	p(diff.Minutes())     //2.9221326803275677e+06
	p(diff.Seconds())     //1.7532796081965408e+08
	p(diff.Nanoseconds()) //175327960819654063
	p("then = ", then)
	p("now = ", now)
	p(then.Add(diff)) //不准确?
	p(then.Add(-diff))

	//

	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000

	fmt.Println(secs)   //1433820891
	fmt.Println(millis) //1433820891536
	fmt.Println(nanos)  //1433820891536026200
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

	expensiveCall()

	// format与parse是相对的,要提供layout常量

	//使用time内置的layout
	p(now.Format(time.RFC3339)) //2015-06-09T11:50:45+08:00
	t1, _ := time.Parse(time.RFC3339, "2015-06-09T11:50:45+08:00")
	p(t1)
	//使用自定义layout,格式必须参考时间 Mon Jan 2 15:04:05 MST 2006
	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2) //0000-01-01 20:00:41 +0000 UTC

	//也可以用标准string格式输出
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()) //2015-06-09T14:09:19-00:00

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM") //错误的parse也会报错
	p(e)

	//	timeAfter() //timeAfter可以创造channel,阻塞等待指定时间执行
	//	stopTimer() //timer可以被取消
	tickerTest() //隔段时间调用一次

}

func expensiveCall() {
	t0 := time.Now()
	//-------------------
	sum := 0
	for i := 0; i < 99999999; i++ {
		sum += i
	}
	//------------------
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
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

func timeAfter() {
	timeChan := time.After(3 * time.Second)
	//	var timeChan <-chan time.Time = time.After(3 * time.Second)
	<-timeChan
	fmt.Println("3 seconds")

	time.AfterFunc(2*time.Second, func() { fmt.Println("another 2 seconds") })
	fmt.Println("AfterFunc不会阻塞")

	time.Sleep(99999999999999) //没有sleep AfterFunc就没有机会执行
}
