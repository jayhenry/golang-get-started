package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {

	x1 := 0
	x2 := 1
	func1 := func() int {
		last := x1
		x1, x2 = x2, x1+x2
		return last
	}

	return func1
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
