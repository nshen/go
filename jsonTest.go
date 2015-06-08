package main

import (
	"encoding/json"
	"fmt"
)

func jsonTest() {
	newDivider("jsonTest.go")

	bolB, _ := json.Marshal(true) //[]byte
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//自定义数据类型
	type Response1 struct {
		Page   int
		Fruits []string
	}

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	type Response2 struct {
		Page   int      `json:"page"` //自定义json key name
		Fruits []string `json:"fruits"`
	}
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//decode
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)                    //map[num:6.13 strs:[a b]]
	strs := dat["strs"].([]interface{}) //类型转换?
	str1 := strs[0].(string)            //类型转换?
	fmt.Println(strs, str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//也可以直接encode到os.Writers里,例如os.Stdout或http
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
