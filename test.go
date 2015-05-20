package main

import (
	"fmt"
	"sync"
)

func goTest() {

	//	testArgs(1, 2, 3, 4, 5)

	//	testInterface()
	expandSlice() //slice内存分配的一个坑
	sliceShrink() //收缩slice
}

//锁
func callLocked(lock *sync.Mutex, f func()) {
	lock.Lock()
	defer lock.Unlock()
	f()
}

func expandSlice() {
	//http://www.zhihu.com/question/27161493

	s0 := make([]int, 2, 10)
	s1 := append(s0, 2)                 //长度购,所以底层与s0指向同一个数组,加个2 [0,0,2]
	s2 := append(s0, 3)                 //长度够,所以底层与s0指向同一个数组,加个3(修改了底层数组,所以s1也变[0,0,3]了) [0,0,3]
	fmt.Println(&s0[0], &s1[0], &s2[0]) // 底层是一个同一个数组
	fmt.Println(s0, s1, s2)             //[0 0] [0 0 3] [0 0 3]

	s0 = []int{0, 0}
	s1 = append(s0, 2)                  //长度不够,所以新创建一个数组,加个2[0,0,2]
	s2 = append(s0, 3)                  //长度不够,所以新创建一个数组,加个3[0,0,3]
	fmt.Println(&s0[0], &s1[0], &s2[0]) //3个不同的数组0xc082006f10 0xc082004e00 0xc082004e20
	fmt.Println(s0, s1, s2)
}

func sliceShrink() {
	fmt.Println("--------------- sliceShrink! ----------------------")
	originSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice := originSlice[4:6]

	fmt.Println(mySlice, len(mySlice), cap(mySlice)) //[5 6] 2 6

	//一个技巧:当底层数组过大,而slice只用到其中一小块,为了垃圾回收,应该copy一个小的出来使用
	var s []int = make([]int, len(mySlice))
	copy(s, mySlice)

	fmt.Println(s, len(s), cap(s))
}

func testArgs(n ...int) {
	fmt.Println("-----testArgs------")
	//n是一个slice
	for i, v := range n {
		fmt.Println(i, v)
	}
}

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue { // 创建对象实例。
	return &Queue{make([]interface{}, 10)}
}
func (*Queue) Push(e interface{}) error { // 省略 receiver 参数名。
	panic("not implemented")
}

// func (Queue) Push(e int) error { // Error: method redeclared: Queue.Push
// panic("not implemented")
// }
func (self *Queue) length() int { // receiver 参数名可以是 self、 this 或其他。
	return len(self.elements)
}
