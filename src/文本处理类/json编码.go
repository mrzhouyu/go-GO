package main

import (
	"encoding/json"
	"fmt"
)

//Create根据结构体生成 json 成员变量名首字母必须大写
type Create struct {
	Name   string   `json:"name"` //若json该字段想要用小写 可以如此操作  属于二次编码
	Id     int      `json:"-"`    //表示此字段不会输出到屏幕
	Lang   []string `json:"_"`    //表示将Lang换成_
	Gpa    float32
	Status bool `json:",string"` //表示转换成字符串类型
}

func main() {
	//定义结构体并且初始化
	S := Create{"Yuchou", 1, []string{"Golang", "C", "Python"}, 5.0, true}

	//编码 根据内容	生成json文本

	//不格式化输出
	buf1, err1 := json.Marshal(S)
	if err1 != nil { //文件出错时候结束运行
		fmt.Println(err1)
		return
	}
	fmt.Println("buf1=", string(buf1)) //返回的是字节 需要string转为字符串

	//格式化输出
	buf2, err2 := json.MarshalIndent(S, "", " ")
	if err2 != nil { //文件出错时候结束运行
		return
	}
	fmt.Println("buf2=", string(buf2))

	//通过map生成json
	m := make(map[string]interface{}, 4) //初始化 由于value有很多类型 所以这里用万能指针类型 长度为4
	m["id"] = 110
	m["name"] = "chen"
	m["sli"] = []string{"C", "C++"}
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println("出错")
		return
	}
	fmt.Println(string(res))

}
