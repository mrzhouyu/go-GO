package main

import "encoding/json"
import "fmt"

func main() {
	//		原始的json文本
	j := `
	{
 	"name": "Yuchou",
 	"id": 1,
 	"lang": [
  		"Golang",
  		"C",
 		 "Python"
 	],
	 "gpa": 5.12,
	 "status": true
	}`
	//定义一个map
	m := make(map[string]interface{}, 4)
	//这里只能传入切片类型 byte 所以需要将其转换
	err := json.Unmarshal([]byte(j), &m) //将json文本转为结构体  因为是更改uncode内容这里需要取地址 给uncode赋值
	if err != nil {
		fmt.Println("err=", err)
		return //不为nil说明出错   结束进程
	}
	fmt.Println(m)
	fmt.Printf("uncode=%+v\n", m)

	//需要注意的是 因为定义的map是interface{}类型 所有
	name := m["name"] //这样是可以的
	fmt.Println(name)
	var str string
	//str = m["name"] //这样不行 虽然看起来name对应的value是string类型但是无法这样转换
	//fmt.Println(str)
	for _, value := range m { //这里的value必须是接口类型的变量
		switch data := value.(type) {
		case string:
			str = data
			fmt.Println("str=", str, "   这样就可以反推类型了")
		case []interface{}: //这里其实不是[]string类型切片 而是interface类型
			fmt.Println(value)

		}
	}

}
