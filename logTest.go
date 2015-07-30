package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func logTest() {
	newDivider("logTest.go")

	//--------------------------
	//log包自带一个可配置logger
	//-------------------------

	//logger可以配置输出的Writer, 前缀,和flag(附加信息)
	log.SetFlags(0)            //不显示时间等信息
	log.SetOutput(os.Stderr)   //Stderr通道
	log.SetPrefix("Buildin: ") //前缀

	//------------------------
	// new 一个logger
	//------------------------

	//log第一个参数可支持任意io.Writer接口
	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	//passing these messages to os.Stderr allows other applications running your program to know an error has occurred.
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	//log可以写到txt文件里
	file, err := os.OpenFile("./tmp/logTest.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//第二个参数flag
	//	O_RDONLY // open the file read-only.
	//	O_WRONLY // open the file write-only.
	//	O_RDWR   // open the file read-write.
	//	O_APPEND // append data to the file when writing.
	//	O_CREATE // create a new file if none exists.
	//	O_EXCL   // used with O_CREATE, file must not exist
	//	O_SYNC   // open for synchronous I/O.
	//	O_TRUNC  // if possible, truncate file when opened.

	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}

	//log写到多处
	multi := io.MultiWriter(file, os.Stdout)

	StdoutAndFile := log.New(multi, "Multi:", log.Ldate|log.Ltime|log.Lshortfile)

	//开始输出!

	//不同的logger一起输出不能保证顺序!
	log.Println("Hello logger")
	log.Println("log print line")
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")

	StdoutAndFile.Println("同时写在txt文件和stdout中")

	log.Println("log print line")
}
