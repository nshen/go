package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func funcTest() {
	newDivider("funcTest.go")

	//---------------------------
	//多返回值直接作为参数
	for i := 1; i <= 4; i++ {
		a, b, c := PythagoreanTriple(i, i+1)
		Δ1 := Heron(a, b, c)
		Δ2 := Heron(PythagoreanTriple(i, i+1))
		fmt.Printf("Δ1 == %10f == Δ2 == %10f\n", Δ1, Δ2)
	}

	//----------------------------
	//参数不定数量
	sliceArgs(1, 2, 3, 4, 5)
	sliceArgs(9, 8, 7)

	//----------------------------
	//factory function
	addZip := MakeAddSuffix(".zip")
	addTgz := MakeAddSuffix(".tar.gz")
	fmt.Println(addTgz("filename"), addZip("filename"), addZip("gobook.zip")) //filename.tar.gz filename.zip gobook.zip

}

//最后的类型前加...,会让前边的参数变成改类型的slice
func sliceArgs(n ...int) {
	fmt.Println(reflect.TypeOf(n)) //[]int //n是一个slice
	for i, v := range n {
		fmt.Println(i, v)
	}
}

//-------------------
// 辅助函数

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//海伦公式,三边求面积
func Heron(a, b, c int) float64 {
	α, β, γ := float64(a), float64(b), float64(c)
	s := (α + β + γ) / 2
	return math.Sqrt(s * (s - α) * (s - β) * (s - γ))
}

//这个需要研究一下
//http://mathforum.org/dr.math/faq/faq.pythag.triples.html
//
//a = r2 - s2,
//b = 2mn,
//c = r2 + s2,
//m > n > 0 are whole numbers,
//m - n is odd, and the greatest common divisor of m and n is 1.
func PythagoreanTriple(m, n int) (a, b, c int) {
	if m < n {
		m, n = n, m
	}
	return (m * m) - (n * n), (2 * m * n), (m * m) + (n * n)
}
