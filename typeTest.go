//一些类型相关的东西

package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func typeTest() {
	newDivider("typeTest.go")

	valueAndPointer()
	testTypes()    //值类型
	testRefTypes() //引用类型
	testInterface()
}

// & is sometimes called the address of operator.
// * is sometimes called the contents of operator or the indirection operator or the dereference operator.
func valueAndPointer() {

	var bp *bool = &b            //指针类型,类型前面加* . 指向b
	fmt.Println("指针类型:", b, *bp) //指针类型: false false  ,前面加*解引用dereferencing
	b = true
	fmt.Println("指针类型:", b, *bp) //指针类型: true true

	//------------------------------

	i := 1
	i_pointer := &i //指针

	fmt.Println(i, i_pointer, *i_pointer, reflect.TypeOf(i), reflect.TypeOf(i_pointer)) //1 0xc082006d90 1 int *int
	i = 2
	fmt.Println(i, i_pointer, *i_pointer, reflect.TypeOf(i), reflect.TypeOf(i_pointer)) //2 0xc082006d90 2 int *int

	z := 37                    // z is of type int
	pi := &z                   // pi is of type *int (pointer to int)
	ppi := &pi                 // ppi is of type **int (pointer to pointer to int)
	fmt.Println(z, *pi, **ppi) //37 37 37
	**ppi++                    // Semantically the same as: (*(*ppi))++ and *(*ppi)++
	fmt.Println(z, *pi, **ppi) //38 38 38

	//----------------------------------
	ii := 9
	jj := 5
	product := 0
	swapAndProduct(&ii, &jj, &product)
	fmt.Println(ii, jj, product) //5 9 45

	//----------------------------------
	//3种方式创建User实例
	//new(Type) ≡ &Type{}  //两种语法等价
	myU := User{1, "Nshen"} // user value
	myU1 := &myU            // pointer to user

	myU2 := &User{2, "Nshen2"} // pointer to user

	myU3 := new(User) // pointer to user // var myU3 *User = new(User)
	myU3.id = 3
	myU3.name = "Nshen3"

	myU4 := &User{2, "Nshen2"} //与u2内容一样,注意看比较
	fmt.Println(
		myU2.id == myU4.id,            //true
		myU2.name == myU4.name,        //true
		myU2 == myU4,                  //false
		reflect.DeepEqual(myU2, myU4)) //true

	fmt.Println(myU)                    //{1 Nshen} 只有这个是值
	fmt.Println(myU1, myU2, myU3, myU4) //&{1 Nshen} &{2 Nshen2} &{3 Nshen3} &{2 Nshen2}
	fmt.Println(
		reflect.TypeOf(myU),  //main.User
		reflect.TypeOf(myU1), //*main.User
		reflect.TypeOf(myU2), //*main.User
		reflect.TypeOf(myU3), //*main.User
		reflect.TypeOf(myU4)) //*main.User

	var flag BitFlag = Active | Send
	fmt.Println(flag) //3(Active|Send)
	fmt.Println(BitFlag(0), Active, Send, flag, Receive, flag|Receive)

}

func swapAndProduct(x, y, product *int) {
	*x, *y = *y, *x
	*product = *x * *y
}

////////////////////
//基本类型
///////////////////////
//go推断默认类型:整数为int,浮点数为float64,复数为complex128

//-------------------
//数字类型
//-------------------

//--------------
//整数: 5种有符号,5种无符号,1个指针类型

// 一般都用int,需要的时候再转成其他类型,  byte (uint8)常用来处理 UTF-8 encoded text
var i int   //int8 int16 int32(rune) int64
var ui uint //uint8(byte) uint16 uint32 uint64
//uintptr //An unsigned integer capable of storing a pointer value (advanced)
var abyte byte   // byte是uint8的别名
var r rune = '哈' //21704 (int32) 必须用单引号

//需要高精度必须用标准库里的
//big.Int for integers
//big.Rat for 有理数,可以表示三分之二,但不能表示π等

//从小到大转换是安全的,但从大到小转换会得到非期望的值,所以一般会自定义转换函数
func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}

//----------------
// 浮点数:两钟浮点数,两钟复数

var f float32 //float64

//int(float) 浮点数转整数,浮点数部分直接丢掉
//下边函数可以四舍五入安全转换
func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x) //整数部分,小数部分 (float64)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

//---------------
// 复数: complex64 complex128 ,complex128最常用,应为math/cmplx都是对128操作的
var c complex64 = 5 + 5i //复数 complex128
func testComplex() {

	var c complex64 = 5 + 5i
	fmt.Println(c, real(c), imag(c)) //(5+5i) 5 5
	//real ,image 对complex64操作,会得到float32 ,对complex128操作会得到float64

	f := 3.2e5                       // type: float64
	x := -7.3 - 8.9i                 // type: complex128 (literal)
	y := complex64(-18.3 + 8.9i)     // type: complex64 (conversion)
	z := complex(f, 13.2)            // 用complex函数构造 type: complex128 (construction)
	fmt.Println(x, real(y), imag(z)) // Prints: (-7.3-8.9i) -18.3 13.2

	//	复数其他相关的操作在math/cmplx包里
}

//--------------------------------

var b bool // boolean值默认为false
var s string = "你好世界,Hello World!"

//常量定义
const limit = 512       // constant; type-compatible with any number
const top uint16 = 1421 // constant; type: uint16

//常量表达式在编译时取值
const hlutföllum = 16.0 / 9.0                               // type: float64
const mælikvarða = complex(-2, 3.5) * hlutföllum            // type: complex128
const erGjaldgengur = 0.0 <= hlutföllum && hlutföllum < 2.0 // type: bool

const (
	ca = iota //0 自加常量
	cb = iota //1
	cc        //2
)

//iota运用在自定义类型上
type BitFlag int

func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}

	if len(flags) > 0 { // int(flag) is vital to avoid infinite recursion!
		return fmt.Sprintf("%d-%b(%s)", int(flag), int(flag), strings.Join(flags, "|"))
	}

	return "0()"
}

const (
	Active  BitFlag = 1 << iota // 1 << 0 == 1
	Send                        // Implicitly BitFlag = 1 << iota // 1 << 1 == 2
	Receive                     // Implicitly BitFlag = 1 << iota // 1 << 2 == 4
)

//------------------------
//结构体
type User struct {
	id   int
	name string
}

//func (self *User) String() string {
//	return fmt.Sprintf("%d, %s", self.id, self.name)
//}

func testTypes() {

	fmt.Println("--------------- testTypes! ----------------------")

	fmt.Println(b, s, s[0], i, ui, r, f, c, ca, cb, cc) //false 你好世界,Hello World! 228 0 0 21704 0 (5+5i) 0 1 2
	fmt.Println(0600)                                   //384 //0开头的是8进制

	myint := 123
	var myUint8 uint8 = 123
	//	fmt.Println(myint + myUint) //报错,不是相通类型不能操作,比较
	//必须转型
	//	向上转型
	fmt.Println(myint+int(myUint8), reflect.TypeOf(myint+int(myUint8))) //246 int
	//  向下转型,自定义函数有些麻烦
	myUint8FromInt, _ := Uint8FromInt(myint)
	fmt.Println(myUint8FromInt+myUint8, reflect.TypeOf(myUint8FromInt+myUint8)) //246 uint8

	runeArr := []rune(s)                             //Rune 是 int32 的别名。用 UTF-8 进行编码。
	runeArr[1] = '坏'                                 //原始字符串标识的值在引号内的字符是不转义的                        //必须用单引号
	fmt.Println(string(runeArr), string(runeArr[0])) //32位数组
	fmt.Printf("Value is: %v \n", c)                 //%v 按原始打印

	abyte = '3'          //go的字符字面量用单引号,并且适应上下文,这里视为一个byte
	digit := abyte - '0' //
	//In UTF-8 (and 7-bit ASCII) 字符'0'的code point (character)十进制为 48, '3' 是code point 51, '3' - '0' 就是51− 48 结果为整数3
	fmt.Println("digit:", digit)

	var anumber uint = '4' //一个字符字面量,就是一个数字,适配任意数字类型
	fmt.Println(anumber)   // 52 //ASCII

	//--------
	// 数组
	// 创建方式:
	//	[length]Type
	//	[N]Type{value1, value2, …, valueN}
	//	[…]Type{value1, value2, …, valueN}	//编译器确定数组长度   //	a2 := [...]int{1, 1, 2, 3, 5}

	var arr [5]int //这里是长度为5的int数组声明,有长度的是数组,没长度的是slice , 数组是值类型

	//遍历数组
	for i, v := range arr {
		fmt.Println("Array element[", i, "]=", v)
	}

	//二维数组
	//	var matrix [4][4]float64

	//二维数组
	//	var bigDigits = [][]string{
	//		{"  000  ",
	//			" 0   0 ",
	//			"0     0",
	//			"0     0",
	//			"0     0",
	//			" 0   0 ",
	//			"  000  "},

	//		{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},

	//		{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	//		{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	//		{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
	//			"   4  "},
	//		{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	//		{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	//		{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	//		{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	//		{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
	//	}

	//数组是值类型,默认会复制
	va1 := [...]int{1, 2}
	va2 := va1
	va2[0] = 3
	fmt.Printf("%d %d\n", va1[0], va2[0]) //1 ,3

	testComplex()

}

//引用类型 slice map channel ,全部用make()创建
//还有functions, and methods.
func testRefTypes() {
	// maps, slices, channels, functions, and methods.
	fmt.Println("--------------- testRefTypes! ----------------------")

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

	//slice当作参数传递给函数,传递的是引用
	grades := []int{87, 55, 43, 71, 60, 43, 32, 19, 63}
	inflate(grades, 3)  //inflate函数内部修改了外部的grades
	fmt.Println(grades) //[261 165 129 213 180 129 96 57 189]

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

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}
