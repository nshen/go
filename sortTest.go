package main

import (
	"fmt"
	"sort"
)

func sortTest() {
	newDivider("sortTest.go")

	//内置排序
	//-------
	strs := []string{"cbc", "dd3", "b33", "aaa"}
	sort.Strings(strs)
	fmt.Println(strs) //[aaa b33 cbc dd3]

	ints := []int{3, 5, 61, 1, 5, 2, 0}
	s := sort.IntsAreSorted(ints) //是否已经排序
	fmt.Println("Sorted: ", s)    // false
	sort.Ints(ints)
	fmt.Println(ints)                                 //[0 1 2 3 5 5 61]
	fmt.Println("Sorted: ", sort.IntsAreSorted(ints)) //true

	//自定义类型排序
	//-------------

	//实现sort.Interface就可以被sort.Sort排序

	/*
			type Interface interface {
		        // Len is the number of elements in the collection.
		        Len() int
		        // Less reports whether the element with
		        // index i should sort before the element with index j.
		        Less(i, j int) bool
		        // Swap swaps the elements with indexes i and j.
		        Swap(i, j int)
			}
	*/

	strArr := []string{"sfasf", "Ddfsdf", "aad333", "66", "bnnn"}
	sort.Sort(ByLength(strArr)) //按字符串长度排序
	fmt.Println(strArr)

	//还有其他的example
	//http://golang.org/pkg/sort
}

type ByLength []string

func (a ByLength) Len() int {
	return len(a)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
