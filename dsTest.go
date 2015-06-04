package main

import (
	"container/list"
	"fmt"
	"strings"
)

func dsTest() {
	newDivider("数据结构")

	//slice
	//-----------------

	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(slice_indexOf(strs, "pear"))                                //2
	fmt.Println(slice_include(strs, "grape"), slice_include(strs, "apple")) //false true

	fmt.Println(slice_any(strs, func(v string) bool { //true
		return strings.HasPrefix(v, "p") //以p开头的
	}))
	fmt.Println(slice_all(strs, func(v string) bool { //false
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(slice_filter(strs, func(v string) bool { //[peach apple pear]
		return strings.Contains(v, "e")
	}))

	fmt.Println(slice_map(strs, func(v string) string { //[PEACH APPLE PEAR PLUM]
		return strings.ToUpper(v)
	}))

	//标准库中的双向链表
	//--------------
	l := list.New()
	l.PushBack(123)
	l.PushBack("abc")
	l.PushBack("456")
	l.PushFront("xyz")

	//遍历链表
	fmt.Println("链表正向遍历")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}

	fmt.Println("链表逆向遍历")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v\n", e.Value)
	}

	//TODO:container包里还有一个环形队列,和一个heap

}

func slice_indexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func slice_include(vs []string, t string) bool {
	return slice_indexOf(vs, t) >= 0
}

//返回vs里是否有满足f的元素
func slice_any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//如果元素全符合f,返回true,否则false
func slice_all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !(f(v)) {
			return false
		}
	}
	return true
}

//符合f的元素装进新的slice里返回
func slice_filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//返回一个经f处理过的新slice
func slice_map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))

	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
