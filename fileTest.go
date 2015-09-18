package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

//TODO ioTest与fileTest合并?

func fileTest() {
	newDivider("fileTest.go")

	//	dirTest()

	//	readFile()
	//	writeFile()
	pathTest()

}

func dirTest() {
	os.Mkdir("nshen", 0666)
	os.MkdirAll("nshen/test1/test2/", 0666)
	err := os.Remove("nshen")
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll("nshen")
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile() {
	///ioutil.WriteFile 一次写入
	err := ioutil.WriteFile("./tmp/test.txt", []byte("写一个文件123abc哈哈!!!"), 0644) //0644?
	checkErr(err)

	//保守先创建再写入
	f, err := os.Create("tmp/test1.txt")
	checkErr(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10} //some
	n2, err := f.Write(d2)               // 写入 byte slice
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n") //直接写字符串
	checkErr(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() //写入硬盘

	//bufio包实现了带buffer的reader和writer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()

}

func readFile() {
	//----------------------
	//ioutil.ReadFile 一次读出整个文件
	//----------------------

	dat, err := ioutil.ReadFile("main.go") //dat是[]byte
	checkErr(err)
	fmt.Println(string(dat))

	//----------------------
	//os.File的Read方法可以读部分文件
	//--------------------------
	f, err := os.Open("main.go") //*os.File
	checkErr(err)
	defer f.Close() //有开就要有关

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkErr(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1)) //5 bytes: // ma

	//File的Seek方法从中间读
	o2, err := f.Seek(6, 0)
	checkErr(err)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2)) //4 bytes @ 6: n包

	//io.ReadAtLeast方法实现上边相同功能
	o3, err := f.Seek(6, 0)
	checkErr(err)
	b3 := make([]byte, 4)
	n3, err := io.ReadAtLeast(f, b3, 4)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//回到头部
	_, err = f.Seek(0, 0)
	checkErr(err)

	//bufio包实现了带buffer的reader和writer
	r4 := bufio.NewReader(f) //*Reader
	b4, err := r4.Peek(5)    //[]byte next 5 bytes
	checkErr(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
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

	// 取当前目录
	currentPath, currentFile := filepath.Split(os.Args[0]) //当前目录,当前文件名
	fmt.Println(currentPath, currentFile)
	// 读取当前目录的其他文件
	inFilename := filepath.Join(currentPath, "README.md")
	inFile, err := os.Open(inFilename)
	if err != nil {
		fmt.Println(err)
	}
	defer inFile.Close()

	inFileBytes, _ := ioutil.ReadAll(inFile)
	fmt.Println(string(inFileBytes))

	components := []string{"a", "path", "with", "..", "relative", "elements"} //".."取消前一个目录
	myPath := path.Join(components...)                                        //   a/path/relative/elements
	fmt.Println("myPath:", myPath)
	var decomposed []string = filepath.SplitList(myPath) //???这个??
	for _, dir := range decomposed {
		fmt.Println(dir, string(filepath.Separator))
	}

	fmt.Println(os.Args[0], filepath.Base(os.Args[0])) //F:\mygo\src\go-experiments\go-experiments.exe go-experiments.exe

	//这样测试 go-experiments.exe *.go // 输出所有.go文件
	//还有这样子目录 go-experiments.exe */*.png
	if len(os.Args) > 1 {
		//test glob
		fmt.Println(os.Args[1:]) // *.go
		fmt.Println(commandLineFiles(os.Args[1:]))
		//[commandLineTest.go cryptoTest.go dsTest.go fileTest.go goroutineTest.go ioTest.
		//go jsonTest.go logTest.go main.go mem.go netTest.go panicTest.go printlnTest.go
		//randomTest.go regexpTest.go sortTest.go stringTest.go test.go timeTest.go typeTe
		//st.go xmlTest.go]
	}
}

// file globbing
//例如将 *.txt, 替换成匹配的文件  README.txt, INSTALL.txt 等等..
func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" { //unix平台会自动glob,windows需要自己处理
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // Invalid pattern
			} else if matches != nil { // At least one match
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}
