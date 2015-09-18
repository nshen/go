//一些类型相关的东西

package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
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

	//----------------------------

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

func (self *User) String() string {
	return fmt.Sprintf("{%p id:%d name:%s}", self, self.id, self.name)
}

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
	// 数组, 长度是固定的,不可变的
	//	cap(arr) == len(arr)
	// 创建方式:
	//	[length]Type
	//	[N]Type{value1, value2, …, valueN}
	//	[…]Type{value1, value2, …, valueN}	//编译器确定数组长度   //	a2 := [...]int{1, 1, 2, 3, 5}

	var arr [5]int      //这里是长度为5的int数组声明,有长度的是数组,没长度的是slice , 数组是值类型
	var buffer [20]byte //uint8 == byte
	fmt.Printf("%-8s |%s\n", "Type", "Len |Contents")
	//(右对齐,8字符宽度,类型) (2字符宽度,10进制数字,值)
	fmt.Printf("%-8T %4d %v\n", arr, len(arr), arr)          //[5]int    5 [0 0 0 0 0]
	fmt.Printf("%-8T %4d %v\n", buffer, len(buffer), buffer) //[20]uint8 20 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	//遍历数组
	for i, v := range arr {
		fmt.Println("Array element[", i, "]=", v)
	}

	//二维数组
	var grid1 [3][3]int
	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	grid2 := [3][3]int{{4, 3}, {8, 6, 2}}
	fmt.Printf("%-8s |%s\n", "Type", "Len |Contents")
	fmt.Printf("%-8T %4d %v\n", grid1, len(grid1), grid1) //[3][3]int  3 [[0 0 0] [8 6 2] [0 0 0]]
	fmt.Printf("%-8T %4d %v\n", grid2, len(grid2), grid2) //[3][3]int  3 [[4 3 0] [8 6 2] [0 0 0]]

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

	////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////
	//slice是引用类型,多种创建方式
	//slice is a variable-length fixed-capacity sequence of items of the same type.

	//1. slice初始化表达式 []Type{value1, value2, …, valueN} ,[]Type{}
	a := []int{0, 0, 0} // 没有长度的是slice,否则为数组
	a[1] = 10

	//2. 基于数组创建slice
	//s[n] The item at index position n in slice s
	//s[n:m] A slice taken from slice s from index positions n to m - 1
	//s[n:] A slice taken from slice s from index positions n to len(s) - 1
	//s[:m] A slice taken from slice s from index positions 0 to m - 1
	//s[:] A slice taken from slice s from index positions 0 to len(s) - 1
	//cap(s) The capacity of slice s; always ≥ len(s)
	//len(s) The number of items in slice s; always ≤ cap(s)
	//s = s[:cap(s)] Increase slice s’s length to its capacity if they are different

	myarr := [5]int{1, 2, 3, 4, 5} // var myarr []int = []int{1, 2, 3, 4, 5}
	b := myarr[:3]
	fmt.Println(b) //[1 2 3]

	//3. make创建 make([]Type, length, capacity) //省略capacity则默认等于length, 如果指定capacity则length必须小于等于capacity
	c := make([]int, 3, 10) //元素个数为3的数组,并预留10个元素存储空间
	c[1] = 10
	fmt.Println(c)                 // [0 10 0]
	fmt.Println("len(c):", len(c)) // 3 真实长度
	fmt.Println("cap(c):", cap(c)) // 10 分配的存储空间

	//*.不常见
	//	s := new([7]string)[:]
	//	s[0], s[1], s[2], s[3], s[4], s[5], s[6] = "A", "B", "C", "D", "E", "F", "G"

	//遍历
	for i, v := range c {
		fmt.Println("Slice element[", i, "]=", v)
	}

	// 创建len和cap都为0的slice
	// []Type{} == make([]Type, 0)

	//----------
	c = append(c, 9, 8, 7)           //可以append多个值
	c = append(c, []int{6, 5, 4}...) // ...打散一个slice然后append
	fmt.Println(c, len(c), cap(c))   //[0 10 0 9 8 7 6 5 4] 9 10

	//append string
	appendb := []byte{'U', 'V'}
	appendletters := "wxy"
	appendb = append(appendb, appendletters...) // Append a string's bytes to a byte slice
	fmt.Printf("%s\n", appendb)                 //UVwxy

	//-------------------------------
	//slice默认指向一个底层数组,所以,一个修改,所有都会被修改

	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[:5]
	u := s[3 : len(s)-1]
	fmt.Println(s, len(s), cap(s)) //[A B C D E F G] 7 7
	fmt.Println(t, len(t), cap(t)) //[A B C D E] 5 7
	fmt.Println(u, len(u), cap(u)) //[D E F] 3 4 // 貌似指定了start index后cap会比len大1,why?
	u[1] = "x"
	fmt.Println(s, t, u) //[A B C D x F G] [A B C D x] [D x F]

	newSlice := c[0:3]          //基于现有的slice创建
	newSlice2 := make([]int, 3) // 准备copy数组
	copy(newSlice2, c)          // newSlice2是c的拷贝,不会跟着原来的改变
	fmt.Println(c, newSlice, newSlice2)
	c[0] = 999 //修改原来的slice,新的也会跟着改变
	fmt.Println(c, newSlice, newSlice2)

	//slice当作参数传递给函数,传递的是引用
	grades := []int{87, 55, 43, 71, 60, 43, 32, 19, 63}
	inflate(grades, 3)  //inflate函数内部修改了外部的grades
	fmt.Println(grades) //[261 165 129 213 180 129 96 57 189]

	buffer := make([]byte, 20, 60)
	//二维slice
	grid1 := make([][]int, 3)
	for i := range grid1 {
		grid1[i] = make([]int, 3)
	}
	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	grid2 := [][]int{{4, 3, 0}, {8, 6, 2}, {0, 0, 0}}
	cities := []string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
	cities[len(cities)-1] = "Karachi"
	fmt.Println("Type Len Cap Contents")
	fmt.Printf("%-8T %2d %3d %v\n", buffer, len(buffer), cap(buffer), buffer)
	fmt.Printf("%-8T %2d %3d %q\n", cities, len(cities), cap(cities), cities)
	fmt.Printf("%-8T %2d %3d %v\n", grid1, len(grid1), cap(grid1), grid1)
	fmt.Printf("%-8T %2d %3d %v\n", grid2, len(grid2), cap(grid2), grid2)

	//slice长度小于容量的时候,可以这样把长度放到最大
	//slice = slice[:cap(slice)]

	//------------
	//s == s[:i]+s[i:] // s is a string; i is an int; 0 <= i <= len(s)

	users := []*User{{1, "n1"}, {2, "n2"}, {3, "n3"}}
	//等于 slice里等于&User{1, "n1"}

	for _, u := range users {
		u.id += 10
	} //外部users被修改了

	//-------------
	//插入
	//---------------

	s = []string{"M", "N", "O", "P", "Q", "R"}
	x := InsertStringSliceCopy(s, []string{"a", "b", "c"}, 0) // At the front //[a b c M N O P Q R]
	y := InsertStringSliceCopy(s, []string{"x", "y"}, 3)      // In the middle //[M N O x y P Q R]
	z := InsertStringSliceCopy(s, []string{"z"}, len(s))      // At the end //[M N O P Q R z]
	fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)

	//--------------
	// 删除
	//-------------

	//从头部删除

	s = []string{"A", "B", "C", "D", "E", "F", "G"}
	s = s[2:]                // Remove s[:2] from the front
	fmt.Println("从头部删除:", s) //从头部删除: [C D E F G]

	//从尾部删除
	s = []string{"A", "B", "C", "D", "E", "F", "G"}
	s = s[:4]               // Remove s[4:] from the end
	fmt.Println("从尾部删除", s) //从尾部删除 [A B C D]

	//从中间删除,比较困难
	s = []string{"A", "B", "C", "D", "E", "F", "G"} //[A B C D E F G]
	x = RemoveStringSliceCopy(s, 0, 2)              // Remove s[:2] from the front [C D E F G]
	y = RemoveStringSliceCopy(s, 1, 5)              // Remove s[1:5] from the middle [A F G]
	z = RemoveStringSliceCopy(s, 4, len(s))         // Remove s[4:] from the end [A B C D]
	fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)

	//---------------------
	// 排序
	//---------------------

	//排序数字

	//sort.Float64s(fs) Sorts fs of type []float64 into ascending order
	//sort.Float64sAreSorted(fs) Returns true if fs of type []float64 is sorted
	//sort.Ints(is) Sorts is of type []int into ascending order
	//sort.IntsAreSorted(is) Returns true if is of type []int is sorted

	//排序字符串

	files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
	fmt.Printf("Unsorted: %q\n", files)         //Unsorted: ["Test.conf" "util.go" "Makefile" "misc.go" "main.go"]
	sort.Strings(files)                         // 标准库的排序算法,区分大小写
	fmt.Printf("Underlying bytes: %q\n", files) //Underlying bytes: ["Makefile" "Test.conf" "main.go" "misc.go" "util.go"]

	fmt.Println(sort.IsSorted(FoldedStrings(files))) //false
	SortFoldedStrings(files)                         // 自定义的排序算法,忽略大小写
	fmt.Printf("Case insensitive: %q\n", files)      //Case insensitive: ["main.go" "Makefile" "misc.go" "Test.conf" "util.go"]
	fmt.Println(sort.IsSorted(FoldedStrings(files))) //true

	//-------------------
	//查找
	//-------------------
	//线性查找
	files = []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
	target := "Makefile"
	for i, file := range files {
		if file == target {
			fmt.Printf("found \"%s\" at files[%d]\n", file, i)
			break
		}
	}

	// sort.Search() 是内置的二叉查找
	//Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true
	sort.Strings(files)
	fmt.Printf("%q\n", files)
	i := sort.Search(len(files),
		func(i int) bool { return files[i] >= target })
	if i < len(files) && files[i] == target {
		fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
	}

	////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////
	//map
	//
	//创建语法
	//make(map[KeyType]ValueType, initialCapacity)  // m := make(map[string]int, 1000) //事先申请大块内存
	//make(map[KeyType]ValueType)  					// monthdays := make(map[string] int)
	//map[KeyType]ValueType{}
	//map[KeyType]ValueType{key1: value1, key2: value2, ..., keyN: valueN}

	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, //← 逗号是必须的
	}
	fmt.Printf("%v Dec:%d\n", monthdays, monthdays["Dec"]) //map[Jul:31 Sep:30 Dec:31 Feb:28 Apr:30 May:31 Jun:30 Nov:30 Jan:31 Mar:31 Aug:31 Oct:31] Dec:31

	//遍历
	year := 0
	for month, days := range monthdays { //← 键如果没有使用，可以用 _, days
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

	//可以用指针做key
	//因为是指针(内存地址不同),允许同样的两个值做为key
	userMap_pr := make(map[*User]int)
	userMap_pr[&User{1, "A"}] = 1
	userMap_pr[&User{1, "A"}] = 2 //与上一个key值相同
	userMap_pr[&User{2, "B"}] = 3
	fmt.Println(len(userMap_pr), userMap_pr) //3 map[{0xc0820374c0 id:1 name:A}:1 {0xc0820374e0 id:1 name:A}:2 {0xc082037500 id:2 name:B}:3]

	//不允许同样的两个值做key,则用值做key
	userMap_v := make(map[User]int)
	userMap_v[User{1, "A"}] = 1
	userMap_v[User{1, "A"}] = 2 //与上一个key值相同
	userMap_v[User{2, "B"}] = 3
	fmt.Println(len(userMap_v), userMap_v) //2 map[{1 A}:2 {2 B}:3]

	//
	//---------

}

////////////////////////////////////////////
// slice 相关函数
////////////////////////////////////////////

//插入insertion数据到slice,返回一个新的slice
func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion)) //创建一个新的slice
	at := copy(result, slice[:index])                   //copy原slice的前半部分
	at += copy(result[at:], insertion)                  //copy需要插入的slice
	copy(result[at:], slice[index:])                    //copy原slice的后半部分
	return result
}

//插入insertion数据到slice,通过修改原slice(and possibly the inserted slice)
func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}

//删除slice中的数据,返回一个新的slice
func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

//删除slice中的数据,通过修改原slice
func RemoveStringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}

//忽略大小写排序字符串
func SortFoldedStrings(slice []string) {
	sort.Sort(FoldedStrings(slice))
}

//Folded Strings实现了sort.Interface
type FoldedStrings []string

func (slice FoldedStrings) Len() int {
	return len(slice)
}
func (slice FoldedStrings) Less(i, j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}
func (slice FoldedStrings) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

///////////////////////////////////////////////////

func testInterface() {
	fmt.Println("--------------- testInterface! ----------------------")

	var o interface{} = &User{1, "Tom"} //泛类型

	//------------------------------
	//Type Assertions
	//
	//resultOfType, boolean := expression.(Type) // Checked
	//resultOfType := expression.(Type) // Unchecked; panic() on failure
	//------------------------------

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

	classifier(5, -17.9, "ZIP", nil, true, complex(1, 1), o)
	//param #0 is an int
	//param #1 is a float64
	//param #2 is a string
	//param #3 is nil
	//param #4 is a bool
	//param #5 is a complex128
	//default: param #6's type is *main.User

}

//用type switch实现的类型分类
func classifier(items ...interface{}) {
	for i, x := range items {
		switch xtype := x.(type) {

		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is an int\n", i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is an unsigned int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		case complex128:
			fmt.Printf("param #%d is a complex128\n", i)
		default:
			fmt.Printf("default: param #%d's type is %T\n", i, xtype)
		}
	}
}

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}
