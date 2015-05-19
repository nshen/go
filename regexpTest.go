package main

import (
	"fmt"
	"regexp"
)

var static = regexp.MustCompile(", *") //编译正则很慢,常用正则应该一次编译

func regexpTest() {
	fmt.Println("#### RegExp -- regexpTest.go ####")
	r, err := regexp.Compile("abcd*")
	if err != nil {
		fmt.Println("正则编译错误")
	}
	str := "abcddd fish, wibble abcd, abc, foo"
	fmt.Printf("Replaced: %v\n", r.ReplaceAllString(str, "bar"))

	fmt.Printf("Replaced: %v\n", static.ReplaceAllString(str, ". "))
}
