package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func jsonTest() {
	newDivider("jsonTest.go")

	//-----------------------------------------------
	//	encode
	//-----------------------------------------------

	//基本类型
	bolB, _ := json.Marshal(true)                   //bolB是[]byte
	fmt.Println(reflect.TypeOf(bolB), string(bolB)) //[]uint8 true

	intB, _ := json.Marshal(1)
	fmt.Println(reflect.TypeOf(intB), string(intB)) //[]uint8 1

	fltB, _ := json.Marshal(2.34)
	fmt.Println(reflect.TypeOf(fltB), string(fltB)) //[]uint8 2.34

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) //"gopher"

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) //["apple","peach","pear"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) //{"apple":5,"lettuce":7}

	//自定义数据类型
	type Response1 struct {
		Page   int
		Fruits []string
	}

	res1D := &Response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) //{"Page":1,"Fruits":["apple","peach","pear"]}

	type Response2 struct {
		Page   int      `json:"page"` //json key tag
		Fruits []string `json:"fruits"`
	}
	res2D := &Response2{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) //{"page":1,"fruits":["apple","peach","pear"]}

	//-----------------------------------------------
	//	decode
	//-----------------------------------------------

	//---------------------------------------
	//解析成map ,key为string, 值为interface{}

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat) //map[num:6.13 strs:[a b]]

	//解析后,通过断言的方式访问

	strs := dat["strs"].([]interface{}) //断言
	str1 := strs[0].(string)            //断言
	fmt.Println(strs, str1)             //[a b] a

	for k, v := range dat {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	//---------------------------
	//解析成结构体
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)

	fmt.Println(res, *res)               //&{1 [apple peach]} {1 [apple peach]}
	fmt.Println(res.Fruits[0], res.Page) //apple 1

	//todo: https://github.com/bitly/go-simplejson

	//也可以直接encode到os.Writers里,例如os.Stdout或http
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) //{"apple":5,"lettuce":7}

}
