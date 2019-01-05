package main

import "encoding/json"
import "fmt"

//由于被解码的jsonkey字段都是首字母都是小写  所以这里要二次编码对应过去
type Create struct {
	Name   string   `json:"name"`
	ID     int      `json:"id"`
	Lang   []string `json:"lang"`
	Gpa    float32  `json:"gpa"`
	Status bool     `json:"status"`
}

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
	//定义一个结构体变量
	var uncode Create
	//这里只能传入切片类型 byte 所以需要将其转换
	err := json.Unmarshal([]byte(j), &uncode) //将json文本转为结构体  因为是更改uncode内容这里需要取地址 给uncode赋值
	if err != nil {
		fmt.Println("err=", err)
		return //不为nil说明出错   结束进程
	}
	fmt.Println(uncode)
	fmt.Printf("uncode=%+v\n", uncode)

}
