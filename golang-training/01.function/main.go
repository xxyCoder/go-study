package main

import (
	"fmt"
)

func main() {
	res5_1, res5_2 := fv5()
	fmt.Println("function version 5 value", res5_1, res5_2)
	data := []float64{1, 2, 3, 4, 6.6}
	fmt.Println("function version 6 value:", fv6(1, 2, 3, 4, 6.6))
	fmt.Println("function version 7 value:", fv7(data))
	// 函数表达式
	fv8 := func() {
		fmt.Println("function version 8")
	}
	fv8()
	fmt.Println("%T\n", fv8)
	// 调用闭包
	w := wrapper()
	fmt.Println("wrapper", w())
	fmt.Println("wrapper", w())
	// 传入回调函数
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
	// defer 关键字的作用
	defer fv1()
	fmt.Println(fv4())
	// 传递指针
	var s, z = "xxy", 10
	changeValue(&z, &s)
	fmt.Println(s, z)
	// 立即执行函数
	func() {
		fmt.Println("立即执行函数")
	}()
}

func fv1() {
	fmt.Println("无参无返回值")
}

func fv2(a string, b string, c int) int {
	fmt.Println("有参有返回值")
	return c
}

func fv3(x int, a, b string, c int) int {
	fmt.Println("参数类型简写")
	return c
}

func fv4() (s string) {
	s = "提供返回值参数，不需要返回其变量"
	return
}

func fv5() (string, string) {
	return "hello world", "多个返回值"
}

func fv6(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func fv7(sf []float64) float64 {
	return fv6(sf...)
}

func fv9() func(s string) string {
	return func(s string) string {
		return "返回一个函数" + s
	}
}

// 闭包
func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// 传入函数
func visit(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

// todo: filter
func filter(numbers []int, callback func(int) bool) []int {
	var res []int
	for _, v := range numbers {
		if callback(v) {
			res = append(res, v)
		}
	}
	return res
}

// 传递指针
func changeValue(z *int, s *string) {
	*z = *z - 1
	*s = "hello world"
}
