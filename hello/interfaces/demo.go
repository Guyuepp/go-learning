// interfaces 是关于接口的一些尝试
package interfaces

import (
	"fmt"
)

var a struct {
	x int
}

func Test1() {
	if a.x == 0 {
		fmt.Println("b is nil")
	}
}

type InfiniteAReader struct{}

func (InfiniteAReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
