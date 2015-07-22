package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func xmlTest() {
	newDivider("xmlTest.go")

	//	readXML()
	writeXML()
}

type Servers struct {
	// struct tag 用来辅助反射
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"` //.attr解析属性
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"` //string 并且innerxml,则原始xml累加到此字段上
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func writeXML() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{ServerName: "SH_VPN", ServerIP: "127.0.0.1"})
	v.Svs = append(v.Svs, server{ServerName: "BJ_VPN", ServerIP: "127.0.0.2"})

	output, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	fmt.Println("------")
	os.Stdout.Write(output)
}

func readXML() {
	file, err := os.Open("static/servers.xml") //for read access
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := &Servers{}
	err = xml.Unmarshal(data, v) //
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v.XMLName)
	fmt.Println(v.Version)
	fmt.Println(v.Description)
	fmt.Println(v.Svs)
}
