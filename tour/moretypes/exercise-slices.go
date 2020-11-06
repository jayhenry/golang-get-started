package main

import (
	"math"

	"golang.org/x/tour/pic"
	//"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	var n [][]uint8 = make([][]uint8, dx)

	for x := 0; x < dx; x++ {
		n[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			n[x][y] = uint8(float64(x) * math.Sin(8.0*float64(y)/float64(dy)))
		}
	}
	//fmt.Println(x)
	return n
}

func main() {
	//fmt.Println(Pic(10,10))
	pic.Show(Pic)
}
