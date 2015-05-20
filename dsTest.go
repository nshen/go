package main

import (
	"container/list"
	"fmt"
)

func dsTest() {
	newDivider("数据结构")
	//标准库中的双向链表
	l := list.New()
	l.PushBack(123)
	l.PushBack("abc")
	l.PushBack("456")
	l.PushFront("xyz")

	//遍历
	fmt.Println("正向遍历")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}

	fmt.Println("逆向遍历")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v\n", e.Value)
	}

	//container包里还有一个环形队列,和一个heap

}
