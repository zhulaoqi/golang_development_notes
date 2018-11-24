package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	//结构也是一种类型
	Name string //定义struct的属性
	Age  int
}

func main() {
	//fmt.Printf("hello world!")
	//var s []string
	//s := []string{"dd", "fsdf", "f"}
	//	////s[3] = "dsf"
	//	//slice1 := s[1:3]
	//	////slice1 := slice1[:1]
	//	//slice1 = append(slice1, "aa", "bb")
	//	//fmt.Println(slice1)
	//s = append(s, "fsd")

	//s := []string{"123", "123"} //切片
	//var slice1 []string = make([]string, 2)
	//slice1 = append(slice1, "a", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd")
	//fmt.Println(slice1)
	//p := new([2]int)
	//p[0] = 22
	//b := make([]int, 0, 50)
	//fmt.Println(p, b)
	//var m map[string]int
	//m := map[string]int{}
	//m := map[string]Person{}
	//p := Person{Name: "jerry", Age: 12}
	//m["ONE"] = p
	//fmt.Println(m)
	//p := Person{Name: "jerry", Age: 12}
	//p.list()
	//const name = "ok"
	//const name1 string = "ok"
	//fmt.Println(name)
	//fmt.Println(name1)

	var data map[interface{}]interface{}
	jsonStr := "{\"Name\":\"test\",\"Age\":19,\"1\":{\"info\":\"hello\"}}"
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
func (p *Person) list() {
	fmt.Println(p.Name)
}
