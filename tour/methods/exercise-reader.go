package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func (MyReader) Read(b []byte) (n int, err error) {
	nb := len(b)
	n = nb
	for i := 0; i < n; i++ {
		b[i] = 'A'
	}
	err = nil
	return
}

func main() {
	reader.Validate(MyReader{})
}
