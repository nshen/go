package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

func stringTest() {
	log.Println("stringTest.go-----------------------")
	fmt.Println("stringTest.go-----------------------")

	a := "Golang!"
	b := "Hello,"

	a, b = swap(a, b) //多返回值
	fmt.Println(a, b)

	str := "世"
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]))

	str1 := "A string"
	str2 := "A " + "string"

	fmt.Println(str1 == str2, &str1 == &str2, &str1, &str2) // 字符串相等,但地址不同 true false
	//遍历字符串

	str3 := "Étoilé,我是N神" //utf-8是可变字符编码, 一个字符会有1~4个byte长, 128个ascii码都是单字节
	// Don’t do this! Ã.toilÃ©,æ..æ.¯Nç¥.
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c", str3[i])
	}
	fmt.Printf("\n")
	// Do this instead Étoilé,我是N神
	for _, c := range str3 {
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")

	//坏数据(从硬盘读取或从网络读取的不完整字符流)要注意
	bytes := str3[0:7]
	str4 := string(bytes)

	for i, c := range str4 {
		//RuneError 0xFFFD
		if c == utf8.RuneError {
			str4 = str4[i:]
			fmt.Println("\n bad data", c)
			break
		} else {
			fmt.Printf("%c", c)
		}
	}

	//trim
	trimStr := "\tThis is a string \n"
	trimStr = strings.Trim(trimStr, " \t\n\r")
	fmt.Println(trimStr) //This is a string
	//split
	words := strings.Split(trimStr, " ") // a slice of strings
	for idx, word := range words {
		fmt.Println(idx, word)
	}

	trimStr = "\tThis is a string \n"
	trimStr = strings.TrimFunc(trimStr, unicode.IsSpace)

	aStr := "123"
	bStr := aStr
	runeArr := []rune(aStr) // var runeArr []rune = []rune(aStr)
	runeArr[1] = '9'
	aStr = string(runeArr)
	fmt.Println(aStr, bStr) //193 123
}

func swap(a string, b string) (string, string) {
	return b, a
}
