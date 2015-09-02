package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

//打印
func printlnTest() {
	newDivider("printlnTest.go")

	type point struct {
		x int
		y int
	}

	p := point{1, 2}

	//--------------
	//结构体
	//--------------

	fmt.Printf("%v\n", p)  //{1 2}  默认格式
	fmt.Printf("%+v\n", p) //{x:1 y:2} 添加属性
	fmt.Printf("%#v\n", p) //main.point{x:1, y:2}
	fmt.Printf("%T\n", p)  //main.point %T返回类型,可以用在任何类型上

	//--------------
	//bool值
	//--------------

	fmt.Printf("%t %t\n", true, false)                         // %t输出boolean  //true false
	fmt.Printf("%d %d\n", IntForBool(true), IntForBool(false)) //用数字的方式输出//1 0

	//------------
	//整数
	//------------

	fmt.Printf("%d\n", 123) //10进制 //123
	fmt.Printf("%b\n", 123) //2进制  //1111011
	fmt.Printf("%o\n", 123) //8进制  //173
	fmt.Printf("%x\n", 123) //16进制 //7b
	fmt.Printf("%X\n", 123) //16进制 //7B

	fmt.Printf("%c\n", 123) //Unicode code point对应字符 //{
	fmt.Printf("%q\n", 123) //a single-quoted character literal safely escaped with Go syntax //'{'
	fmt.Printf("%U\n", 123) //Unicode format //U+007B

	//2进制格式
	//加6表示宽度6个字符宽度,16是10个character默认右对齐,%-16b 负号表示左对齐,%016是用0占位,% 16b是用空格占位
	//|1111011|         1111011|1111011         |0000000001111011|         1111011|
	fmt.Printf("|%6b|%16b|%-16b|%016b|% 16b|\n", 123, 123, 123, 123, 123)

	//8进制格式
	//#表示用前边加1个0表示, %# 8o表示宽度8个字符空格占位,+表示写出正负号
	//|51|051|     051|    +051|-0000051|
	fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)

	//16进制格式
	//#表示加一个0X在数字前边
	i := 3931
	fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i) //|f5b|F5B|     f5b|00000f5b|0X0F5B|0x0F5B|

	//10进制格式
	//|$569|$000569|$+00569|$***569|
	i = 569
	fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, Pad(i, 6, '*'))

	//----------------
	//浮点数
	//----------------

	fmt.Printf("%f\n", math.Pi) //10进制无指数						//3.141593
	fmt.Printf("%g\n", math.Pi) //for large exponents, %f otherwise //3.141592653589793
	fmt.Printf("%G\n", math.Pi) //3.141592653589793

	fmt.Printf("%e\n", math.Pi) //科学计数法							//3.141593e+00
	fmt.Printf("%E\n", math.Pi) //科学计数法							//3.141593E+00
	fmt.Printf("%b\n", math.Pi) //2进制记数法? strconv.FormatFloat with 'b' format//7074237752028440p-51

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)   //排格式,与整数一样|  1.20|  3.45|
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //左对齐|1.20  |3.45  |

	for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
		fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ',')) //Humanize参数(要显示的小数,显示宽度,小数位数,补位字符,分隔字符)
	}

	//---------------------------
	//character
	//Go中character就是rune(int32) 他们既可以被输出为 整数 也可以按照 unicode字符 输出
	//---------------------------

	//%d 十进制显示
	//%#04x 十六进制,0x开头,以0补齐,占4个位置
	//%U Unicode format
	//%c Unicode character
	fmt.Printf("%d %#04x %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A6') //934 0x03a6 U+03A6 'Φ'

	//---------------
	//字符串
	//字符串可以指定打印最小宽度(太段会被空格补位),最大字符数(太长会被切掉)
	//可以输出为Unicode字符,或一系列code points(runes),或者utf8 bytes
	//----------------

	//%s strings输出
	//%q go字符串格式双引号输出
	//%+q +modifier用来只打印ASCII字符(U+0020 to U+007E),其他输出escapes
	//%#q #modifier用来输出go的raw string,如果失败则输出quoted string
	slogan := "End Óréttlæti♥"
	fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)
	//End Óréttlæti♥
	//"End Óréttlæti♥"
	//"End \u00d3r\u00e9ttl\u00e6ti\u2665"
	//`End Óréttlæti♥`

	s2 := "Dare to be naïve"
	fmt.Printf("|%22s|%-22s|%10s|\n", s2, s2, s2) //|22字符宽度|22字符宽度左对齐|最小10个字符宽度(被撑开所以都打印出来了)
	//|      Dare to be naïve|Dare to be naïve      |Dare to be naïve|
	i2 := strings.Index(s2, "n")
	fmt.Printf("|%.10s|%.*s|%-22.10s|%s|\n", s2, i2, s2, s2, s2)
	//%.10 表示最大10个字符
	//%.*s 需要两个参数最大字符宽度和字符串
	//%-22.10s 表示左对齐,最小22个字符宽度,但最大只显示10个字符(所以没显示完整)
	//|Dare to be|Dare to be |Dare to be            |Dare to be naïve|

	fmt.Printf("%s %s\n", "string", "\"string\"") //string "string" 		//the uninterpreted bytes of the string or slice
	fmt.Printf("%q %q\n", "string", "\"string\"") //"string" "\"string\""  //a double-quoted string safely escaped with Go syntax
	fmt.Printf("%x\n", "hex this")                //16进制跟整数的%x一样

	fmt.Printf("|%6s|%6s|\n", "foo", "b")   //|   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //|foo   |b     |

	//---------------
	//slice
	//---------------

	chars := []rune(slogan)
	fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars) // 16进制 ,加0x的16进制,加0X的16进制
	//[45 6e 64 20 d3 72 e9 74 74 6c e6 74 69 2665]
	//[0x45 0x6e 0x64 0x20 0xd3 0x72 0xe9 0x74 0x74 0x6c 0xe6 0x74 0x69 0x2665]
	//[0X45 0X6E 0X64 0X20 0XD3 0X72 0XE9 0X74 0X74 0X6C 0XE6 0X74 0X69 0X2665]

	//一般的slice都会输出中括号包裹的格式,但[]byte不会输出方括号,除非使用%v
	//每个byte用两个16进制数字表示
	bytes := []byte(slogan)
	fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)
	//End Óréttlæti♥
	//456e6420c39372c3a974746cc3a67469e299a5
	//456E6420C39372C3A974746CC3A67469E299A5
	//45 6E 64 20 C3 93 72 C3 A9 74 74 6C C3 A6 74 69 E2 99 A5
	//[69 110 100 32 195 147 114 195 169 116 116 108 195 166 116 105 226 153 165]

	fmt.Println([]float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%#v\n", []float64{math.E, math.Pi, math.Phi}) //#加了类型
	fmt.Printf("%.5f\n", []float64{math.E, math.Pi, math.Phi})
	//[2.718281828459045 3.141592653589793 1.618033988749895]
	//[2.718281828459045 3.141592653589793 1.618033988749895]
	//[]float64{2.718281828459045, 3.141592653589793, 1.618033988749895}
	//[2.71828 3.14159 1.61803]

	fmt.Printf("%q\n", []string{"Software patents", "kill", "innovation"}) //字符串专用 %q
	fmt.Printf("%v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%#v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%17s\n", []string{"Software patents", "kill", "innovation"})
	//["Software patents" "kill" "innovation"]
	//[Software patents kill innovation]
	//[]string{"Software patents", "kill", "innovation"}
	//[ Software patents              kill        innovation]

	//map
	fmt.Printf("%v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%#v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"}) //用来生成代码
	fmt.Printf("%v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%#v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%04b\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	//map[2:B 3:C 4:D 1:A]
	//map[int]string{4:"D", 1:"A", 2:"B", 3:"C"}
	//map[1:1 2:2 3:4 4:8]
	//map[int]int{3:4, 4:8, 1:1, 2:2}
	//map[0001:0001 0010:0010 0011:0100 0100:1000]

	//-------
	//指针
	//-------
	iPointer := 5
	fPointer := -48.3124
	sPointer := "Tomás Bretón"
	fmt.Printf("|%p → %d|%p → %f|%#p → %s|\n", &iPointer, iPointer, &fPointer, fPointer, &sPointer, sPointer) //%#p 加了＃省略0x
	//|0xc082007a40 → 5|0xc082007a50 → -48.312400|c082007a60 → Tomás Bretón|

	//Sprint 打印成字符串
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	//Fprintf 打印到writer接口,例如可以打印到os.Stdin或http.ResponseWriter
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

func IntForBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s) //utf8.RuneCountInString 返回字符数量
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}
	return s
}

//returns a string representation of the number it is given with grouping separators
//(for languages that use simple three-digit groups) and padding.
func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)        //返回整数与小数部分
	whole := fmt.Sprintf("%+.0f", dollars)[1:] //+表示加正负号,0位小数,这句去掉了整数部分的正负号 Strip "±"
	fraction := ""
	if decimals > 0 { //小数位数修改
		//.*修改器,如果decimals为2,则这句等于%+.2f
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:] //把前边的正负号和0去掉 Strip "±0"
	}
	sep := string(separator) //逗号分隔符
	//整数位数加逗号
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	//把最开始整数去掉的正负号加回来
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 { //如果指定的宽度大于字符串宽度,说明有空位
		return strings.Repeat(string(pad), gap) + number //有几个空位就重复即便pad指定的字符+数字
	}
	return number
}
