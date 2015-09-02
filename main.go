// main包是可独立运行的go程序，会产生可执行文件
// 其他包名则会生成.a文件
// 多个文件用一个package名,与都写在一个文件里没什么区别
// 规则:大写字母开头的方法是public 小写字母开头的方法是private的
// go build 会编译到当前目录,go install会编译到bin目录
package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func init() {
	//	fmt.Println("main init()")
}

func env() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows")
	default:
		// freebsd, openbsd,
		// plan9...
		fmt.Printf("%s.", os)
	}

	fmt.Println("CPU核数: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	//环境变量
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
	os.Setenv("FOO", "aa1")               //不会真的改系统环境变量,临时设置?
	fmt.Println("FOO:", os.Getenv("FOO")) //aa1
	fmt.Println("BAR:", os.Getenv("BAR")) //
}

func main() {
	who := "World!"
	if len(os.Args) == 1 {

		fmt.Println("这是我的golang学习实验项目:)")
		//		fmt.Println(filepath.Base(os.Args[0])) //go-experiments.exe 最后一个节点
		//		fmt.Println(os.Args[0])                //F:\mygo\src\go-experiments\go-experiments.exe

	} else if len(os.Args) > 1 { /* os.Args[0] 是 "go-experiments.exe" */
		who = strings.Join(os.Args[1:], " ") //os.Args是个string的slice,把参数拼成字符串
	}

	fmt.Println("Hello", who)

	//	mem()
	//	env()

	//	fmt.Println("aaaaaa")
	//	mu := make(chan bool)
	//	go func() {
	//		<-mu
	//		fmt.Println("cccccc")
	//	}()
	//	fmt.Println("bbbbb")
	//	mu <- true

	//	<-mu

	//	cryptoTest()
	//	randomTest()
	//	printlnTest()
	//	sortTest() //排序
	//	ioTest()
	//	fileTest()
	//	timeTest() //时间相关
	//	goroutineTest()
	//	panicTest() //错误处理
	typeTest() //数据类型
	//	goTest() //未整理
	//	stringTest() //字符串
	//	netTest() //网络相关
	//	regexpTest() //正则
	//	dsTest() //数据结构
	//	jsonTest()
	//	xmlTest()
	//	commandLineTest()
	//	goTest()
	//	logTest()
}

func newDivider(str string) {
	fmt.Println("--------------- ", str, " ---------------")
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
