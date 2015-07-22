package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*

In Go, a string is in effect a read-only slice of bytes
在Go中,string只是一个高效的只读bytes的slice
string可以存放任意bytes,跟格式无关,不必须为utf8格式
但go的代码使用utf8,所以string literal永远是合法的utf8

用index访问的是每个byte,而不是字符
Go source code is UTF-8, so the source code for the string literal is UTF-8 text.
a regular  will also always contain valid UTF-8.

Those sequences represent Unicode code points, called runes.
rune as an alias for the type int32

只有一种情况go把字符串视为utf8的,就是 for range loop 时候

*/
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" // 16进制表示

func stringTest() {
	newDivider("stringTest.go")

	fmt.Println("Println:")
	fmt.Println(sample) //��=� ⌘

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i]) //bd b2 3d bc 20 e2 8c 98
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample) //bdb23dbc20e28c98

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample) //bd b2 3d bc 20 e2 8c 98

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample) //"\xbd\xb2=\xbc ⌘"

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample) //"\xbd\xb2=\xbc \u2318"

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	//U+65E5 '日' starts at byte position 0
	//U+672C '本' starts at byte position 3
	//U+8A9E '語' starts at byte position 6

	//只有for range会视为utf8 ,其他情况就要用到标准库了,下边用到unicode.utf8库,与上边功能一致
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
	//http://golang.org/pkg/unicode/utf8/

	fmt.Println("------------")
	str1 := "A string"
	str2 := "A " + "string"
	fmt.Println(str1 == str2, &str1 == &str2, &str1, &str2) // true false 0xc082006960 0xc082006970 等号可以判断字符串相等,但地址不同

	str3 := "Étoilé,我是N神"         //utf-8是可变字符编码, 一个字符会有1~4个byte长, 128个ascii码都是单字节
	fmt.Println(len([]byte("包"))) //3
	fmt.Println(len([]rune("包"))) //1

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

	//strconv包,转换类型
	fmt.Println("-----")
	parseTest()
	fmt.Println("-----")
	appendTypes()
	fmt.Println("-----")
	formatTest()
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

//Parse 系列函数把字符串转换为其他类型
func parseTest() {
	b, _ := strconv.ParseBool("false")
	fmt.Println(b)

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) //1.234

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i) //123

	//16进制也支持
	//第2个参数是0,会根据前缀自动选择base,0x开头16进制,0开头8进制,否则都为10进制
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d) //456

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u) //789

	_, e := strconv.Atoi("wat") //Atoi is shorthand for ParseInt
	fmt.Println(e)              //strconv.ParseInt: parsing "wat": invalid syntax

}

//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
func appendTypes() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}

//Format 系列函数把其他类型的转换为字符串
func formatTest() {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}
