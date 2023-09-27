package main

import "fmt"

func Variable() {
	// 变量被声明，默认赋予初始值
	var a string = "Hello GO"
	var b = "Hello xxyCoder"
	var x, y, z int
	c := "Hello World"
	d, e, f := 1, 2.0, "Hi"
	var g *int = &d
	var h int = *g

	fmt.Println(a, b, x, y, z, c, e, f, h)
}
