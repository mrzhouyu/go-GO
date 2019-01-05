package main

import "fmt"
import "strconv"

func main() {
	//转换为字符串后追加到字节数组后面 这里的数组类型必须是自己类型
	//定义一个 byte类型数组 长度为0  容量为1000
	sli := make([]byte, 0, 1000)
	//将布尔值添转换为ascii 添加到数组
	sli = strconv.AppendBool(sli, true)
	//第二个数为需要添加的数字 第三个指定十进制格式追加 将10进制 转换为ascii添加到数组
	sli = strconv.AppendInt(sli, 1024, 10)
	//转为ascii添加到数组
	sli = strconv.AppendQuote(sli, "abcd")
	fmt.Println(sli) //打印的时候可以转换为string类型

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(true) //bool类型转为字符串
	//浮点型转为字符串  'f'表示以小数方式 -1指的是小数点位置(这里是紧缩模式)  64表示以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)

	//整形转字符串
	str = strconv.Itoa(1024)
	//字符串转为其他类型
	var value bool
	var err error
	value, err = strconv.ParseBool("flase") //字符串转为bool型  err可以用_取代
	//字符型转为整形
	i, _ := strconv.Atoi("1234")
}
