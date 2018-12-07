/* 类似于python的爬虫  利用的是curl工具 linux里的*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] { //获取命令行输入 url

		resp, err := http.Get(url) //得到返回对象
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :%v \n", err)
			os.Exit(1)

		}
		// b, err := ioutil.ReadAll(resp.Body)          //iotout为读取响应流数据  具体是拷贝结构体到os.Stdout 需要申请一个缓冲区 b   方式1
		nbytes, err := io.Copy(os.Stdout, resp.Body) //直接读取内容  方式2   返回内容 大小 和err 直接写入到标准输出中 io.Stdout  (打印到屏幕) 不需要申请缓冲区间
		resp.Body.Close()                            //读取完毕需要关闭 防止资源泄露
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s : %v \n", url, err)
			os.Exit(1)

		}
		fmt.Printf("%d  ", nbytes) //打印标准输出流

	}

}
