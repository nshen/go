package main

import (
	"fmt"
	"math/rand"
)

func randomTest() {
	newDivider("randomTest.go")
	//math/rand 提供伪随机数

	// intn 返回 [0,n)
	fmt.Println(rand.Intn(100), rand.Intn(100), rand.Intn(100)) //0 <= n < 100

	fmt.Println(rand.Float64())           //0.0 <= f < 1.0
	fmt.Println((rand.Float64() * 5) + 5) //5.0 <= f' < 10.0.

	//新source,source相通,随机数列就相同
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100), r1.Intn(100), r1.Intn(100)) //0 <= n < 100

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100), r2.Intn(100), r2.Intn(100)) //0 <= n < 100
}
