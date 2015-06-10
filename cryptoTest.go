package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func cryptoTest() {
	newDivider("cryptoTest.go")

	s := "sha1 this string"

	h := sha1.New()    //hash.Hash
	h.Write([]byte(s)) //字符串转byte slice
	bs := h.Sum(nil)   // []byte 最终结果,可以append到一个已经存在的byte slice上

	fmt.Println(bs)
	fmt.Printf("%x\n", bs)

	//md5用法与sha1一模一样
	m := md5.New()
	m.Write([]byte(s))
	ms := m.Sum(nil)
	fmt.Println(ms)
	fmt.Printf("%x\n", ms)

	//base64
	data := "abc123!?$*&()'-=@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data)) //[]byte > string
	fmt.Println(sEnc)
	sDec, err := base64.StdEncoding.DecodeString(sEnc) //string > []byte
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sDec))
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
