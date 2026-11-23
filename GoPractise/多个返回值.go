package main

import "fmt"

/*
多个返回值
函数可以返回任意数量的返回值。

swap 函数返回两个字符串。
*/
func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
