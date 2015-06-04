package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func stringTest() {
	newDivider("stringTest.go")

	str1 := "A string"
	str2 := "A " + "string"
	fmt.Println(str1 == str2, &str1 == &str2, &str1, &str2) // true false 0xc082006960 0xc082006970 等号可以判断字符串相等,但地址不同

	//遍历字符串
	str3 := "Étoilé,我是N神" //utf-8是可变字符编码, 一个字符会有1~4个byte长, 128个ascii码都是单字节

	//遍历byte,输出错误~
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c", str3[i]) //Ã.toilÃ©,æ..æ.¯Nç¥.
	}
	fmt.Printf("\n")

	//正确应该用range遍历
	for _, c := range str3 {
		fmt.Printf("%c", c) //Étoilé,我是N神
	}
	fmt.Printf("\n")

	//坏数据(从硬盘读取或从网络读取的不完整字符流)要注意
	bytes := str3[0:7] //0~6 Étoil�
	str4 := string(bytes)

	for i, c := range str4 {
		//RuneError 0xFFFD
		if c == utf8.RuneError {
			str4 = str4[i:]
			fmt.Println("\n bad data", i, c, str4)
			break
		} else {
			fmt.Printf("%c", c)
		}
	}

	aStr := "123"
	bStr := aStr
	runeArr := []rune(aStr) // var runeArr []rune = []rune(aStr)
	runeArr[1] = '9'
	aStr = string(runeArr)
	fmt.Println(aStr, bStr) //193 123

	//utf8
	//---------------------------------
	str := "世"
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]))

	stringsPackage()
}

//strings包 提供字符串常见函数,注意是包方法,不是string对象的方法,所以第一个参数要传进字符串
func stringsPackage() {

	fmt.Println("Contains:", strings.Contains("test", "es")) //是否包含es
	fmt.Println("Count: ", strings.Count("test", "t"),
		strings.Count("cheese", "e"),
		strings.Count("five", "")) //2 3 5
	fmt.Println("HasPrefix", strings.HasPrefix("test", "te"), strings.HasPrefix("test", "es")) //HasPrefix true false
	fmt.Println("HasSuffix", strings.HasSuffix("test", "st"), strings.HasSuffix("test", "te")) //HasSuffix true false
	fmt.Println("Index", strings.Index("test", "e"))                                           //Index 1

	//split与join相反,string转数组,数组转string
	arr_join_split := []string{"a", "b", "c"}
	str_join_split := strings.Join(arr_join_split, "-")
	arr_join_split = strings.Split(str_join_split, "-")
	fmt.Println("Join", str_join_split)  //Join a-b-c
	fmt.Println("Split", arr_join_split) //Split [a b c]

	fmt.Println("Repeat", strings.Repeat("a", 5))                //Repeat aaaaa
	fmt.Println("Replace", strings.Replace("foo", "o", "0", -1)) //Replace f00 最后一个参数-1表示不限制数量,全部替换

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
}
