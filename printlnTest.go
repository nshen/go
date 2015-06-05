package main

import (
	"fmt"
	"os"
)

func printlnTest() {
	newDivider("printlnTest.go")

	type point struct {
		x int
		y int
	}

	p := point{1, 2}

	//打印
	//---------------------

	//结构体
	fmt.Printf("%v\n", p)  //{1 2}
	fmt.Printf("%+v\n", p) //{x:1 y:2}
	fmt.Printf("%#v\n", p) //main.point{x:1, y:2}
	fmt.Printf("%T\n", p)  //类型 main.point

	//bool值
	fmt.Printf("%t\n", true)

	//整数
	fmt.Printf("%d\n", 123)             //10进制整数 123
	fmt.Printf("|%6d|%16d|\n", 12, 345) //加数字按格式输出 |    12|             345|
	fmt.Printf("%b\n", 123)             //2进制 1111011
	fmt.Printf("%c\n", 123)             //ASCII字符 {
	fmt.Printf("%x\n", 123)             //16进制 7b
	//浮点数
	fmt.Printf("%f\n", 78.9)                   //78.900000
	fmt.Printf("%e\n", 123.4)                  //1.234000e+02 //科学记数法
	fmt.Printf("%E\n", 123400000.0)            //1.234000E+08 //科学记数法
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)   //排格式,与证书一样|  1.20|  3.45|
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //左对齐|1.20  |3.45  |

	//字符串
	fmt.Printf("%s\n", "\"string\"")        //"string"
	fmt.Printf("%q\n", "\"string\"")        //"\"string\""
	fmt.Printf("%x\n", "hex this")          //16进制跟整数的%x一样
	fmt.Printf("|%6s|%6s|\n", "foo", "b")   //|   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //|foo   |b     |
	s := fmt.Sprintf("a %s", "string")      //a string
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
