package main

import (
	"fmt"
	"time"
)

// 声明在外部定义的函数
// func flushICache(begin, end uintptr)

// 返回多个返回值
func getX2AndX3(x int) (int, int) {
	return 2 * x, 3 * x
}
func getX2AndX3_2(x int) (x2 int, x3 int) {
	x2 = 2 * x
	x3 = 3 * x
	return
}

// 传递变长参数
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	m := s[0]
	for _, v := range s {
		if v < m {
			m = v
		}
	}
	fmt.Println(m)
	return m
}

// defer允许我们推迟函数返回之前一刻执行
func tryDefer() {
	fmt.Println("try defer start")
	defer min(1, 2, 3, 4, 5)
	fmt.Println("try defer end")
}

// 将函数作为参数
func callback(y int, f func(int) (int, int)) {
	a, b := f(y)
	fmt.Println(a, b)
}

// 闭包
func add(x int) func() {
	return func() {
		x++
		fmt.Println("add", x)
	}
}

// 计算函数执行时间
func calcTime() {
	start := time.Now()
	getX2AndX3_2(1)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("getX2AndX3_2 use ", delta)
}

func main() {
	// 空白符_表示丢弃该值
	x2, _ := getX2AndX3(1)
	fmt.Println(x2)
	defer fmt.Println("call func before")
	tryDefer()
	fmt.Println("main return")
	// 多个defer,会逆序执行,类似栈结构
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	callback(1, getX2AndX3)
	a := add(1)
	a()
	calcTime()
}
