package main

import (
	"fmt"
	"io"
	"strings"
)

func ioTest() {
	newDivider("ioTest.go")

	ioReadTest()

}

func ioReadTest() {
	data, err := ReadFrom(strings.NewReader("from stringgg"), 12)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data, len(data))
	for _, str := range data {
		fmt.Print(string(str))
	}
	fmt.Print("\n")
}

//ReadFrom可以把实现io.Reader接口任何东东(标准输入、文件、字符串等)读到byte数组里
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
