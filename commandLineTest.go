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

	//	scan() //输入的字母转大写输出
	commandArgs()

	flagTest()
}

func flagTest() {
	wordPtr := flag.String("name", "value", "usage")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

func commandArgs() {
	argsWithProgram := os.Args

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
