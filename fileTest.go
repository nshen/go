package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//TODO ioTest与fileTest合并?

func fileTest() {
	newDivider("fileTest.go")

	//	pathTest()

	readFile()

	//读部分文件
	f, err := os.Open("main.go") //*os.File
	checkErr(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkErr(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1)) //5 bytes: // ma

	o2, err := f.Seek(6, 0)
	checkErr(err)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2)) //4 bytes @ 6: n包

	fmt.Println(len([]byte("包"))) //3
	fmt.Println(len([]rune("包"))) //1

}

//一次读出整个文件
func readFile() {
	dat, err := ioutil.ReadFile("main.go") //dat是[]byte
	checkErr(err)
	fmt.Println(string(dat))
}

//func readFile() {
//	file, err := os.Open("test.go")
//	if err != nil {
//		fmt.Println("Error: ", err.Error())
//		return
//	}
//	buffer := make([]byte, 100)
//	for n,e := file.Read(buffer);e==nil;
//}
func pathTest() {
	components := []string{"a", "path", "with", "..", "relative", "elements"} //".."取消前一个目录
	myPath := path.Join(components...)                                        //   a/path/relative/elements
	fmt.Println("myPath:", myPath)
	var decomposed []string = filepath.SplitList(myPath) //???这个??
	for _, dir := range decomposed {
		fmt.Println(dir, string(filepath.Separator))
	}
}
