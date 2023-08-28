package main

import "fmt"

func main() {
	// 创建切片
	var a = []int{1, 2, 3}
	fmt.Println(a[0:2])
	// make创建切片
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zǎo'ān")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
	// 合并切片
	var b = []int{5, 6, 7}
	b = append(b, a...)
	fmt.Println(b)
	// 删除某部分
	var c = append(b[:2], a[2:]...)
	fmt.Println(c)
	// 多维创建
	students_1 := [][]string{}
	students_2 := make([][]string, 35)
	fmt.Println(students_1, students_2)
}
