// main包是可独立运行的go程序，会产生可执行文件
// 其他包名则会生成.a文件
// 规则:大写字母开头的方法是public 小写字母开头的方法是private的
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU核数: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	randomTest()
	//	printlnTest()
	//	sortTest() //排序
	//	ioTest()
	//	fileTest()
	//	timeTest() //时间相关
	//	goroutineTest()
	//	panicTest() //错误处理
	//  typeTest() //数据类型
	//	goTest() //未整理
	//	stringTest() //字符串
	//	netTest()//网络相关
	//	regexpTest() //正则
	//	dsTest() //数据结构
	//	jsonTest()
	//	goTest()
}

func newDivider(str string) {
	fmt.Println("--------------- ", str, " ---------------")
}
