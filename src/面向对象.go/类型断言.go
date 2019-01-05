package main

import "fmt"

type Student struct {
	code int
	name string
}

func main() {
	var i = make([]interface{}, 3)
	i[0] = 100
	i[1] = "hello go"
	i[2] = Student{1, "i am student"}

	for index, data := range i {
		if _, ok := data.(int); ok == true {
			fmt.Printf("i[%d]类型是int型\n", index)
		} else if _, ok := data.(string); ok == true {
			fmt.Printf("i[%d]类型是string型\n", index)
		} else if _, ok := data.(Student); ok == true {
			fmt.Printf("i[%d]类型是Student\n", index)
		}
	}

	for index, data1 := range i {
		switch t := data1.(type) { //t其实是值data1必须是接口类型的变量才能调用.(type)
		case int:
			fmt.Printf("i[%d]类型是%v型\n", index, t)
		case string:
			fmt.Printf("i[%d]类型%v型\n", index, t)
		case Student:
			fmt.Printf("i[%d]类型是%v型\n", index, t)
		}

	}
}
