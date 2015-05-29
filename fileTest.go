package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func fileTest() {
	newDivider("fileTest.go")

	pathTest()

	//	readFile()
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
