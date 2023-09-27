package main

import "fmt"

const PI = 3.1415926              // 隐式类型定义
const S string = "Hello Constant" // 显示类型定义
// 数字型常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出
const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009

// 作枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

// iota作为枚举值
// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
const (
	a = iota      // a = 0
	b             // b = 1
	c             // c = 2
	d = 5         // d = 5
	e             // e = 5
	x             // x = 5
	f = iota + 10 // f = 16
	g             // 每当 iota 在新的一行被使用时，它的值都会自动加 1，并且没有赋值的常量默认会应用上一行的赋值表达式
)

func f1() int {
	return 1
}

func Constant() {
	// 常量的值必须是能够在编译时期就能够确定的
	// const n = f1()
	const n = 2 / 3.0
	fmt.Println(x, f, g)
}
