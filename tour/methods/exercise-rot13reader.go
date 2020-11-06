package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	fmt.Println(n, err)

	if err == nil {
		return
	}
	for i := 0; i < n; i++ {
		//b[i]
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
