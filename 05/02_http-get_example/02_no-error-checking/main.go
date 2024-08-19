package main

import (
	"fmt"       //用于格式化输入和输出。例如，打印到控制台。
	"io/ioutil" // 为了读取响应体（response body）
	"net/http"  //用于发送HTTP请求和接收HTTP响应
)

func main() {
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	//使用http.Get函数发送一个GET请求到指定URL (http://www.geekwiseacademy.com/)。返回的结果存储在res中
	page, _ := ioutil.ReadAll(res.Body)
	//使用ioutil.ReadAll函数从响应体（res.Body）中读取所有数据，并将其存储在page变量中
	res.Body.Close()
	//关闭响应体，释放连接
	fmt.Printf("%s", page) //打印
}
