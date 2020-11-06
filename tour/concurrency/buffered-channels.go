package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	// for i := 0; i < 100; i++ {
	// ch <- i
	// }
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
