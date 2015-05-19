package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/websocket"
)

func netTest() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows")
	default:
		// freebsd, openbsd,
		// plan9...
		fmt.Printf("%s.", os)
	}

	log.Println("netTest.go")
	//golangTest.exe -addr ":8080"
	addr := flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() //调用parse后才有值

	//fmt.Println("服务器开启", *addr) // addr是string指针

	log.Println("服务器开启", *addr)

	//首页未定义
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// 输出用户自定义错误
		http.Error(res, "404 找不到页面", 404)
		//		io.WriteString(res, "hello, world!\n")
	})
	//静态服务器,将localhost/assets/ 的访问映射到 static目录
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static"))))

	//websocket
	http.HandleFunc("/ws", upgradeHandler)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("listenAndServer err")
	}
}

//:= 只能在方法内使用
var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
}

func upgradeHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", 405)
		return
	}
	//	conn, err := upgrader.Upgrade(res, req, nil)
	//	if err != nil {
	//		log.Fatalln("upgrader error")
	//	}
}
