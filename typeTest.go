//一些类型相关的东西

package main

import (
	"fmt"
	"reflect"
)

func typeTest() {
	newDivider("typeTest.go")
	testTypes()    //值类型
	testRefTypes() //引用类型
	testInterface()
}

//基本类型
var b bool // boolean值默认为false
var s string = "你好世界,Hello World!"
var i int   //int8 int16 int32 int64
var ui uint //uint8 uint16 uint32 uint64
//uintptr
//byte(uint8)
var r rune = '哈'         //21704 (int32) 必须用单引号
var f float32            //float64
var c complex64 = 5 + 5i //复数 complex128

//常量定义
const (
	ca = iota //0 自加常量
	cb = iota //1
	cc        //2
)

//结构体
type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func testTypes() {
	fmt.Println("--------------- testTypes! ----------------------")

	fmt.Println(b, s, s[0], i, ui, r, f, c, ca, cb, cc) //false 你好世界,Hello World! 228 0 0 21704 0 (5+5i) 0 1 2
	fmt.Println(0600)                                   //384 //0开头的是8进制
	runeArr := []rune(s)                                //Rune 是 int32 的别名。用 UTF-8 进行编码。
	runeArr[1] = '坏'                                    //原始字符串标识的值在引号内的字符是不转义的                        //必须用单引号
	fmt.Println(string(runeArr), string(runeArr[0]))    //32位数组
	fmt.Printf("Value is: %v \n", c)                    //%v 按原始打印

	var arr [5]int //这里是长度为5的int数组声明,有长度的是数组,没长度的是slice , 数组是值类型
	//遍历数组
	for i, v := range arr {
		fmt.Println("Array element[", i, "]=", v)
	}

	//二维数组
	//	var matrix [4][4]float64

	//编译器确定数组长度
	//	a2 := [...]int{1, 1, 2, 3, 5}

	//数组是值类型,默认会复制
	va1 := [...]int{1, 2}
	va2 := va1
	va2[0] = 3
	fmt.Printf("%d %d\n", va1[0], va2[0]) //1 ,3

}

//引用类型 slice map channel
func testRefTypes() {

	fmt.Println("--------------- testRefTypes! ----------------------")

	//	值类型也可以做为引用
	i := 1
	i_ref := &i

	fmt.Println(i, i_ref, *i_ref, reflect.TypeOf(i), reflect.TypeOf(i_ref)) //1 0xc082006aa0 1 int *int
	i = 2
	fmt.Println(i, i_ref, *i_ref, reflect.TypeOf(i), reflect.TypeOf(i_ref)) //2 0xc082006aa0 2 int *int

	//-------------------------
	//slice是引用类型,多种创建方式
	//1. slice初始化表达式
	a := []int{0, 0, 0} // 没有长度的是slice,否则为数组
	a[1] = 10

	//2. 基于数组创建slice
	myarr := [5]int{1, 2, 3, 4, 5} // var myarr []int = []int{1, 2, 3, 4, 5}
	b := myarr[:3]
	fmt.Println(b) //[1 2 3]

	//3. make创建
	c := make([]int, 3, 10) //元素个数为3的数组,并预留10个元素存储空间
	c[1] = 10
	fmt.Println(c)                 // [0 10 0]
	fmt.Println("len(c):", len(c)) // 3 真实长度
	fmt.Println("cap(c):", cap(c)) // 10 分配的存储空间

	//遍历
	for i, v := range c {
		fmt.Println("Slice element[", i, "]=", v)
	}

	c = append(c, 9, 8, 7)           //可以append多个值
	c = append(c, []int{6, 5, 4}...) // ...打散一个slice然后append
	fmt.Println(c)                   //[0 10 0 9 8 7 6 5 4]

	newSlice := c[0:3]          //基于现有的slice创建
	newSlice2 := make([]int, 3) // 准备copy数组
	copy(newSlice2, c)          // newSlice2是c的拷贝,不会跟着原来的改变
	fmt.Println(c, newSlice, newSlice2)
	c[0] = 999 //修改原来的slice,新的也会跟着改变
	fmt.Println(c, newSlice, newSlice2)

	//--------------------
	//map
	// 用make声明, monthdays := make(map[string] int)
	// m := make(map[string]int, 1000) //事先申请大块内存

	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, //← 逗号是必须的
	}
	fmt.Printf("%d\n", monthdays["Dec"])

	//遍历
	year := 0
	for month, days := range monthdays { //← 键没有使用，因此用 _, days
		fmt.Println(month, days)
		year += days
	}
	fmt.Printf("Numbers of days in a year: %d\n", year)
	//查询是否有值
	_, ok := monthdays["Jan"]
	fmt.Println("查询map: ", ok)
	//删除
	delete(monthdays, "Mar")

	type Any interface{} // 任意类型的key
	anyKeyMap := make(map[Any]int)
	anyKeyMap[12] = 12
	anyKeyMap["12"] = 13
	anyKeyMap[12.0] = 14
	fmt.Println(anyKeyMap[12], anyKeyMap["12"], anyKeyMap[12.0])
	//---------

}

func testInterface() {
	fmt.Println("--------------- testInterface! ----------------------")

	//3种方式创建User实例
	myU := User{1, "Nshen"}
	myU1 := &myU // 指针

	myU2 := &User{2, "Nshen2"}

	myU3 := new(User) // var myU3 *User = new(User)
	myU3.id = 3
	myU3.name = "Nshen3"

	fmt.Println(myU, myU1, myU2, myU3)                                                                 //{1 Nshen} 1, Nshen 2, Nshen2 3, Nshen3
	fmt.Println(reflect.TypeOf(myU), reflect.TypeOf(myU1), reflect.TypeOf(myU2), reflect.TypeOf(myU3)) //main.User *main.User *main.User *main.User

	var o interface{} = &User{1, "Tom"} //泛类型

	//是否是User类型的实例
	u, ok := o.(*User)
	// u := o.(User) // panic: interface is *main.User, not main.User
	fmt.Println(u, ok) //1, Tom true

	//判断实例是否实现了fmt.Stringer接口
	if i, ok := o.(fmt.Stringer); ok { // ok-idiom
		fmt.Println("ok", i) //ok 1, Tom
	}

	//type switch
	switch v := o.(type) {
	case nil: // o == nil
		fmt.Println("nil")
	case fmt.Stringer: // interface

		fmt.Println("fmt.Stringer", v)
	case *User: // *struct
		fmt.Printf("%d, %s\n", v.id, v.name)
	case func() string: // func
		fmt.Println("func()", v())

	default:
		fmt.Println("unknown")
	}
}
