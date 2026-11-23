package main

import "fmt"

/*
当两个或多个连续的具名函数参数具有相同类型时，可以省略除最后一个参数之外其他参数的类型。
在这个例子中，我们将
x int, y int
简写为
x, y int
*/
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add2(45, 44))

}
