package main

// see https://godoc.org/golang.org/x/tour/tree#Tree
import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walkin(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walkin(t.Left, ch)
	ch <- t.Value
	Walkin(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	Walkin(t, ch)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		n1, ok1 := <-c1
		n2, ok2 := <-c2
		if ok1 != ok2 {
			return false
		} else if ok1 == false {
			return true
		} else if n1 != n2 {
			return false
		}
	}
}

func main() {
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	// for i := range ch {
	// 	fmt.Println(i)
	// }

	res := Same(tree.New(1), tree.New(1))
	fmt.Println(res)
}
