package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var staticRegexp = regexp.MustCompile("p([a-z]+)ch") //编译正则很慢,常用正则应该一次编译

func regexpTest() {
	newDivider("regexpTest.go")
	//http://golang.org/pkg/regexp/

	//是否匹配模式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	//编译后的正则可以重复使用
	r, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		fmt.Println("正则编译错误")
	}
	fmt.Println(r.MatchString("peach"))                                // true
	fmt.Println(r.FindString("peach punch"))                           //peach  返回第一个匹配
	fmt.Println(r.FindStringIndex("peach punch"))                      //[0 5]  返回第一个匹配的 开始与结尾的索引
	fmt.Println(r.FindStringSubmatch("peach punch"))                   //[peach ea] 同时返回陪陪p([a-z]+)ch 与 ([a-z]+).
	fmt.Println(r.FindAllString("peach punch pinch", 2))               //[peach punch] 第2个参数指定匹配几个
	fmt.Println(r.FindAllString("peach punch pinch", -1))              //[peach punch pinch]  -1表示全部匹配
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //上边方法的组合[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	fmt.Println("Replaced: ", r.ReplaceAllString("peach123", "punch")) //Replaced:  punch123

	//去掉方法名中的String则为匹配byte数组
	fmt.Println(r.Match([]byte("peach"))) //true

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) //匹配的文字,交给func处理
	fmt.Println(string(out))                   //a PEACH

	//MustCompile is like Compile but panics if the expression cannot be parsed. It simplifies safe initialization of global variables holding compiled regular expressions.
	fmt.Println(staticRegexp) //p([a-z]+)ch
	fmt.Println(staticRegexp.ReplaceAllString("a peach", "<fruit>"))
}
