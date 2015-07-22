package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go-experiments/session"

	"github.com/gorilla/websocket"
)

func netTest() {
	newDivider("netTest.go")
	sessionTest()
	//	urlValuesTest()
	//	parseURL()

	//	return
	//golangTest.exe -addr ":8080"
	addr := flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() //调用parse后才有值

	fmt.Println("服务器开启", *addr) // addr是string指针

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

	//http://localhost:8080/hello?url_long=111&url_long=222
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/upload", uploadFunc)
	http.HandleFunc("/cookie", cookieTest)

	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("listenAndServer err")
	}
}

func urlValuesTest() {

	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	fmt.Println(v.Encode()) // "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

	fmt.Println("--------------------")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数,默认不解析
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginFunc!!!!")
	sess := globalSessions.SessionStart(w, r)
	fmt.Println(sess)
}

func loginFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("loginFunc!!!!")
	//	fmt.Println("method:", r.Method) // 获取请求方法
	sess := globalSessions.SessionStart(w, r)
	fmt.Println(sess)

	if r.Method == "GET" {

		crutime := time.Now().UnixNano()
		//		h := md5.New()
		//		io.WriteString(h, strconv.FormatInt(crutime, 10))
		//		token := fmt.Sprintf("%x", h.Sum(nil))
		token := strconv.FormatInt(crutime, 10) //token用来防止重复提交
		t, _ := template.ParseFiles("static/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		fmt.Println(crutime, token)
		//		t.Execute(w, token)
		t.Execute(w, sess.Get("username"))
	} else {

		r.ParseForm() //必须先Parse才能访问, r.Form 是 url.Values
		token := r.Form.Get("token")
		if token != "" {
			//验证token合法性

		} else {
			//token不存在报错
		}

		if len(r.Form["username"][0]) == 0 {
			fmt.Println("username为空")
		} else {
			sess.Set("username", r.Form["username"][0])
		}

		ageint, err := strconv.Atoi(r.Form["age"][0])

		if err != nil {
			fmt.Println("age不是数字")
		} else {
			if ageint > 100 {
				fmt.Println("age太大")
			}

		}
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
			fmt.Println("username是汉字")
		}

		//template.HTMLEscapeString用来过滤html,防止跨站脚本攻击
		fmt.Println("username:", template.HTMLEscapeString(r.Form["username"][0]))
		fmt.Println("password:", template.HTMLEscapeString(r.Form["password"][0]))
		fmt.Println("age:", template.HTMLEscapeString(r.Form["age"][0]))

	}
}

func uploadFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	if r.Method == "GET" {

		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("static/upload.gtpl")
		t.Execute(w, token)
	} else {
		//接收enctype="multipart/form-data"
		r.ParseMultipartForm(32 << 20)                 // maxMemory 32M 不需要调用ParseForm,会自动调用
		file, handler, err := r.FormFile("uploadfile") //copy这个file
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header) //客户端打印结构体 FileHeader
		f, err := os.OpenFile("./static/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}

func cookieTest(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
	}

	fmt.Println(r.Cookies()) //[]*Cookie
	fmt.Printf("%v", r.Cookies())
	fmt.Println("---")

	//遍历cookie
	for _, cookie := range r.Cookies() {
		fmt.Fprintln(w, cookie.Name, cookie.Value)
	}

	//判断是否存在某cookie
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		fmt.Println("没有 auth")
		//		w.Header().Set("Location", "/login")
		//		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error()) //some other error
	} else {
		fmt.Println("已经auth")
	}

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0) //1nian
	cookie := http.Cookie{Name: "username", Value: "nshen", Expires: expiration}
	http.SetCookie(w, &cookie)

	//跳转
	//	w.Header().Set("Location", "/chat")
	//	w.Header()["Location"] = []string{"/chat"} //跟上边一句一样
	//	w.WriteHeader(http.StatusTemporaryRedirect)
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

	p(u.Path) //pathA/pathB/pathC
	//segs := strings.Split(u.Path, "/")
	p(u.Fragment)                      //f123
	p(u.RawQuery)                      //k=v
	m, _ := url.ParseQuery(u.RawQuery) //url.Values
	p(m)                               //map[k:[v]]
	p(m["k"][0])                       //v
}

//-----------------------------
// session

var globalSessions *session.Manager

func sessionTest() {

	var err error
	globalSessions, err = session.NewManager("memory", "gosessionid", 3600)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("globalSessions init", globalSessions)
}
