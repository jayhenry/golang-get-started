package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 1. 声明变量
	var a string = "heihei"
	a2 := "hehe"
	// 2. 指针
	fmt.Println("string", a, &a, "b:", a2, &a2)
	  //println("hello world")

	var b, c int = 1, 2
	fmt.Println("int:", b, c)

	// 3. if
	if b < c {
		fmt.Printf("b < c: %d %d\n", b, c)
	}

	// for
	for i := 0; i < 3; i++ {
		fmt.Println("in for: ", i)
	}

	// 数组
	strings := []string{"google", "runoob"}
	  // numbers := [6]int{1, 2, 3, 5}

	// range遍历
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// 匿名函数，闭包
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	println(getSquareRoot(2))

	// 普通函数
	var x, y = swap("x", "y")
	println(x, y)

	// 字典
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// 接口、struct实现接口
	var phone Phone
	phone = new(NokiaPhone)
	// struct的方法
	phone.call()

	// 函数返回值指定变量名
	// 错误处理的error接口的方法Error()
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	// goroutine轻量级线程
	go say("world")
	// 普通函数调用
	say("hello")

	// channel可用来在两个goroutine之间通信
	sum_with_channel()
}

func swap(x, y string) (string, string) {
	return y, x
}

/* 定义交换值函数*/
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func sum_with_channel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
