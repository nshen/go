package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
	Unicode

	Every Unicode character has a unique identifying number called a code point.
	There are more than 100000 Unicode characters defined, with code points ranging in value from 0x0 to 0x10FFFF (unicode.MaxRune)
	每个code point通常写成u+4个16进制数字, 例如U+21D4
	每个code point在go中表示为一个rune (int32的别名)
	Unicode text whether in files or in memory—must be represented using an encoding. go使用utf-8

	一般编程语言string都是定长字符(fixed-width characters),Go使用变长字符(variable-width characters)一般使用UTF8 encoding

	UTF-8 encoding
	使用1~4个 bytes来表示每个code point , 128个 ascii码都用1个字节保存,对照表http://www.asciima.com

	在Go中,string实际上只是一个只读的[]byte,string可以存放任意bytes,跟格式无关,不必须为utf8,
	但go的代码使用utf8,所以string literal永远是合法的utf8

	用index访问的是每个byte,而不是字符,所以只访问ascii内的字符时比较有用
 	如果需要index访问字符,可以随时把string转成[]rune(Unicode code points数组) ,再索引访问

	[]rune与string互相转换
	chars := []rune(s) //[]rune or []int32
	s := string(chars)

	string其实就是[]byte, 所以转换至[]byte只需要O1
	[]byte(string)

	遍历的话, for range loop 遍历unicode code points

	用+=连接string,也可以用bytes.Buffer更有效率的做这件事
	var buffer bytes.Buffer
	for {
		if piece, ok := getNextValidString(); ok {
		buffer.WriteString(piece)
	} else {
		break
		}
	}
	fmt.Print(buffer.String(), "\n")
*/

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" // 16进制表示

func stringTest() {
	newDivider("stringTest.go")

	æs := ""
	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00E6'} {
		æs += string(char)
	}
	fmt.Println(æs) //ææææææ

	text1 := "\"what's that?\", he said" // Interpreted string literal
	text2 := `"what's that?", he said`   // 重音符包裹的是Raw string literal,不可转义
	radicals := "√ \u221A \U0000221a"    // .go文件使用utf8编码,所以,可以直接写utf8编码 radicals == "√ √ √"

	fmt.Println(text1, text2, radicals)

	fmt.Println("打印16进制表示的字符串:")
	fmt.Println(sample) //��=� ⌘

	//len(s)函数得到bytes长度,字符长度用len([]rune(s)),或者用更快一点的utf8.RuneCountInString()
	fmt.Println(utf8.RuneCountInString("abcd123"))      //7
	fmt.Println(utf8.RuneCountInString("abc_N神d123哈哈")) //12

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
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:]) //第一个rune,和占用的byte宽度
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

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

	//----------------------
	// String相关的包
	//----------------------

	stringsPackage() //strings包 http://golang.org/pkg/strings/
	strconvPackage() //strconv包 http://golang.org/pkg/strconv/
	utf8Package()    //utf8包 	 http://golang.org/pkg/unicode/utf8/
	unicodePackage() //unicode包 http://golang.org/pkg/unicode/
	regexpPackage()  //regexp包  http://golang.org/pkg/regexp/

}

func regexpPackage() {
	//忘光了,重读 Mastering Regular Expressions
}

//test unicode code points的属性
func unicodePackage() {
	fmt.Println(unicode.IsSpace(' '), unicode.IsSpace('a'), //true false
		unicode.IsLetter('1'), unicode.IsLetter('a'), //false true
		unicode.IsDigit('a'), unicode.IsDigit('5')) //false true

	//是否为16进制
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit, 'b')) //true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit, '5')) //true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit, 'B')) //true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit, 'G')) //false

}

//utf8包实现utf-8 byte序列与runes之间的转换
func utf8Package() {

	//	包内常量
	//	RuneError = '\uFFFD'   // the "error" Rune or "Unicode replacement character"
	//	RuneSelf = 0x80        // characters below Runeself are represented as themselves in a single byte.
	//	MaxRune = '\U0010FFFF' // Maximum valid Unicode code point.
	//	UTFMax = 4             // maximum number of bytes of a UTF-8 encoded Unicode character.

	//Decode系列函数解析字符串(或[]byte),返回解析出的rune和大小
	//-------------------------------
	fmt.Println("\n---------- DecodeLastRune / DecodeLastRuneInString  -------------")
	//	b := []byte("Hello, 世界")
	b := "Hello,世界"
	for len(b) > 0 { //从后向前遍历
		//		r, size := utf8.DecodeLastRune(b)
		r, size := utf8.DecodeLastRuneInString(b)
		fmt.Printf("%c %v, ", r, size) //界 3, 世 3, , 1, o 1, l 1, l 1, e 1, H 1, 97 1
		b = b[:len(b)-size]
	}

	fmt.Println("\n--------- DecodeRune / DecodeRuneInString  -------------")
	b = "Hello,世界"
	for len(b) > 0 { //从前向后遍历
		//		r, size := utf8.DecodeRune(b)
		r, size := utf8.DecodeRuneInString(b)
		fmt.Printf("%c %v ", r, size) //H 1 e 1 l 1 l 1 o 1 , 1 世 3 界 3 97 1
		b = b[size:]
	}

	fmt.Println("\n--------- EncodeRune  -------------")
	buf := make([]byte, 3)
	w := utf8.EncodeRune(buf, '世')
	fmt.Printf("写入%v个bytes到%#v中\n", w, buf)

	fmt.Println("--------- RuneCount / RuneCountInString  -------------")
	//获得字数
	fmt.Println(utf8.RuneCount([]byte("abc_N神d123哈哈"))) //12
	fmt.Println(utf8.RuneCountInString("abc_N神d123哈哈")) //12

	fmt.Println("----- RuneLen-----")

	fmt.Println(utf8.RuneLen('a')) //1 //需要多少字节encode这个rune
	fmt.Println(utf8.RuneLen('界')) //3

	fmt.Println("---- RuneStart -----")
	// RuneStart 报告是否为编码后rune的第一个byte
	buf = []byte("世界")
	fmt.Println(utf8.RuneStart(buf[0])) //true
	fmt.Println(utf8.RuneStart(buf[1])) //false
	fmt.Println(utf8.RuneStart(buf[2])) //false
	fmt.Println(utf8.RuneStart(buf[3])) //true

	fmt.Println("\n--------- FullRune / FullRuneInString / Valid / ValidString / ValieRune  -------------")

	// FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.
	// An invalid encoding is considered a full Rune since it will convert as a width-1 error rune.

	str := "世"
	//true if str begins with a UTF-8-encoded rune
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]))

	buf = []byte{228, 184, 150} //世
	//true if buf begins with a UTF-8-encoded rune
	fmt.Println(utf8.FullRune(buf))
	fmt.Println(utf8.FullRune(buf[:2]))

	//Valid
	fmt.Println(utf8.Valid(buf), utf8.Valid(buf[:2]), utf8.Valid(buf[1:])) //true false false
	//ValidString
	fmt.Println(utf8.ValidString(str), utf8.ValidString(str[:2]), utf8.ValidString(str[1:])) //true false false
	//ValidRune
	fmt.Println(utf8.ValidRune('神'), utf8.ValidRune(utf8.MaxRune), utf8.ValidRune(utf8.MaxRune+1)) //true true false

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

	trimStr = "\tThis is a string \n"
	trimStr = strings.TrimFunc(trimStr, unicode.IsSpace)

	//split---------------
	words := strings.Split(trimStr, " ") // a slice of strings
	for idx, word := range words {
		fmt.Println(idx, word)
	}

	names := "Niccolò•Noël•Geoffrey•Amélie••Turlough•José"
	fmt.Print("|")
	for _, name := range strings.Split(names, "•") {
		fmt.Printf("%s|", name)
	}
	fmt.Println() //|Niccolò|Noël|Geoffrey|Amélie||Turlough|José|

	for _, record := range []string{"László Lajtha*1892*1963", "Édouard Lalo\t1823\t1892", "José Ángel Lamas|1775|1814"} {
		fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
			//FieldsFunc第一个参数是字符串,第二个参数是一个func(c rune)bool函数,字符串中每个字符都会执行第二个函数
			switch char {
			case '\t', '*', '|':
				return true
			}
			return false
		}))
	}

	//replace
	names = " Antônio\tAndré\tFriedrich\t\t\tJean\t\tÉlisabeth\tIsabella \t"
	//把/t换成空格
	names = strings.Replace(names, "\t", " ", -1) //-1 meaning as many as possible
	fmt.Printf("|%s|\n", names)                   //| Antônio André Friedrich   Jean  Élisabeth Isabella  |
	//多余空格删除
	fmt.Println(strings.Join(strings.Fields(strings.TrimSpace(names)), " ")) //Antônio André Friedrich Jean Élisabeth Isabella

	//Map()方法用来替换某个字符
	asciiOnly := func(char rune) rune {
		if char > 127 {
			return '?' //非ascii字符用?代替,return -1则删掉这个字符
		}
		return char
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich")) //第一个参数是一个函数用来修改字符
	//J?r?me ?sterreich
}

func strconvPackage() {
	parseFuncs()
	formatFuncs()
	appendFuncs()
}

//Parse 系列函数把字符串转换为其他类型
func parseFuncs() {
	////字符串转boolean
	b, _ := strconv.ParseBool("false")
	fmt.Println(b) //false

	for _, truth := range []string{"1", "t", "TRUE", "false", "F", "0", "5"} {
		if b, err := strconv.ParseBool(truth); err != nil {
			fmt.Printf("\n{%v}", err)
		} else {
			fmt.Print(b, " ")
		}
	}
	fmt.Println() //true true true false false false {strconv.ParseBool: parsing "5": invalid syntax}

	////字符串转数字
	x, err := strconv.ParseFloat("-99.7", 64)
	fmt.Printf("%8T %6v %v\n", x, x, err) // float64  -99.7 <nil>

	y, err := strconv.ParseInt("71309", 10, 0) //返回int64
	fmt.Printf("%8T %6v %v\n", y, y, err)      //   int64  71309 <nil>

	//第2个参数指定进制数,如果是0则会根据前缀自动选择base,0x开头16进制,0开头8进制,否则都为10进制
	d, _ := strconv.ParseInt("0x1c8", 0, 64)         //16进制
	fmt.Println(d)                                   //456
	j, err := strconv.ParseInt("0707", 0, 32)        // 8进制
	fmt.Println(j, err)                              //455 <nil>
	k, err := strconv.ParseInt("10111010001", 2, 32) //2进制
	fmt.Println(k, err)                              //1489 <nil>

	u, _ := strconv.ParseUint("789", 0, 64) //有负号会失败
	fmt.Println(u)                          //789

	z, err := strconv.Atoi("71309")       //(ASCII to int) 相当于strconv.ParseInt(s, 10, 0)但会返回int
	fmt.Printf("%8T %6v %v\n", z, z, err) // int  71309 <nil>
	_, e := strconv.Atoi("wat")
	fmt.Println(e) //strconv.ParseInt: parsing "wat": invalid syntax
}

//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
func appendFuncs() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendFloat(str, 3.1415926, 'E', -1, 64) //64位
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str)) //4567false"abcdefg"3.1415926E+00'单'
}

//Format 系列函数把其他类型的转换为字符串
func formatFuncs() {
	fmt.Println(strconv.FormatBool(false))                //false
	fmt.Println(strconv.Itoa(1234))                       //Integer to ASCII 1234
	fmt.Println(strconv.FormatInt(1234, 10))              //1234
	fmt.Println(strconv.FormatInt(1234, 2))               //10011010010
	fmt.Println(strconv.FormatInt(1234, 16))              //4d2
	fmt.Println(strconv.FormatUint(1234, 10))             //1234
	fmt.Println(strconv.FormatFloat(123.23, 'g', 12, 64)) //123.23

	s := "Alle ønsker å være fri."
	quoted := strconv.Quote(s)
	fmt.Println(quoted)                  //"Alle ønsker å være fri."
	fmt.Println(strconv.Unquote(quoted)) //Alle ønsker å være fri. <nil>
}
