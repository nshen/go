package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func netTest() {
	newDivider("netTest.go")

	parseURL()
	return
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

func parseURL() {
	s := "postgres://user:pass@host.com:5432/pathA/pathB/pathC?k=v#f123"
	u, err := url.Parse(s) //*url.URL
	if err != nil {
		panic(err)
	}

	p := fmt.Println

	p(u.Scheme) //postgres
	p(u.User)   //user:pass
	if u.User != nil {
		p(u.User.Username()) //user
		pass, b := u.User.Password()
		p(pass, b) //pass true
	}

	p(u.Host) //host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	p(host) //host.com
	p(port) //5432

	p(u.Path)                          //pathA/pathB/pathC
	p(u.Fragment)                      //f123
	p(u.RawQuery)                      //k=v
	m, _ := url.ParseQuery(u.RawQuery) //url.Values
	p(m)                               //map[k:[v]]
	p(m["k"][0])                       //v
}
