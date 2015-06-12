package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func commandLineTest() {
	newDivider("commandLineTest.go")

	//		scan() //输入的字母转大写输出
	commandArgs()
	flagTest()
}

//go-experiments.exe -numb=123 -fork=true -name="nshen" -svar=a b c d
//输出
//numb: 123
//fork: true
//name: nshen
//svar: a
//tail: [b c d]
func flagTest() {
	numbPtr := flag.Int("numb", 42, "an int") //名字,默认值,描述
	boolPtr := flag.Bool("fork", false, "a bool")
	namePtr := flag.String("name", "value", "usage")
	var stringVar string
	flag.StringVar(&stringVar, "svar", "bar", "a string var") //跟上边一样,但可以绑定变量

	flag.Parse()

	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("name:", *namePtr)
	fmt.Println("svar:", stringVar)
	fmt.Println("tail:", flag.Args())
}

func commandArgs() {
	argsWithProgram := os.Args //参数数组

	if len(argsWithProgram) <= 1 {

		return
	}

	argsWithoutProgram := os.Args[1:]

	fmt.Println(argsWithProgram)    //[go-experiments.exe a b c d]
	fmt.Println(argsWithoutProgram) //[a b c d]
	arg := os.Args[3]
	fmt.Println(arg) //c

	//运行 go-experiments.exe a b c d
}

func scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
		if err := scanner.Err(); err != nil {
			fmt.Println(os.Stderr, "error", err)
			os.Exit(1)
		}
	}

}
