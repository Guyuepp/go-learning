// main 包是一个无任何用处的包, 想到什么写什么, 主要用于自己学习
package main

import (
	"fmt"

	"github.com/Guyuepp/go-learning/hello/morestrings"

	// "io"

	"github.com/Guyuepp/go-learning/hello/interfaces"
)

func main() {
	interfaces.Test1()
	fmt.Println(morestrings.ReverseRunes("!dlroW ,olleH"))
	// var r io.Reader = interfaces.InfiniteAReader{}
	// io.Copy(os.Stdout, r)
	mp := make(map[string]int)
	mp["one"] = 1
	if _, ok := mp["one"]; ok {
		fmt.Println("mp[\"one\"] =", mp["one"])
	}
}
