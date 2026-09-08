package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// 1. 完整声明：var 变量名 类型 = 值
	var name string = "gopher"

	var age = 18

	country := "zhongguo"

	fmt.Println("姓名：", name)
	fmt.Println("年龄：", age)
	fmt.Println("国家：", country)
}
