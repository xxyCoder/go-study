package main

import "fmt"

func calc() int {
	return 6
}

func PrintNumber(number int) {
	j := 0
HERE:
	println(j)
	j++
	if j == number {
		return
	}
	// goto语句的标签应该在goto语句出现之前定义
	goto HERE
}

func main() {
	x := 2
	if x == 1 {
		fmt.Println("x = 1")
	} else if x == 2 {
		fmt.Println("x = 2")
	} else {
		fmt.Println("x is unknow")
	}

	if v := 10; v > x {
		fmt.Println("v greater than x")
	} else {
		fmt.Println("v not greater than x")
	}

	// 一旦匹配到某个case或default,在执行完相应代码之后会退出整个switch,即不需要显式指定break
	// 可以使用fallthrough关键字去继续执行后续分支代码
	switch x {
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
		fallthrough
	default:
		fmt.Println("x is unknow")
	}

	switch {
	case x == 1:
		fmt.Println("x = 1")
	case x == 2:
		fmt.Println("x = 2")
	case x > 3:
		fmt.Println("x greater than three")
	default:
		fmt.Println("x = 3")
	}

	switch result := calc(); {
	case result < 0:
		fmt.Println("result less than zero")
	case result == 0:
		fmt.Println("result equal zero")
	default:
		fmt.Println("result greater than zero")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		if i == 4 {
			break
		}
	}

	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Println(i)
	}

	str := "Hello,xxyCoder"
	for pos, char := range str {
		fmt.Println(pos, char)
	}

	PrintNumber(5)
}
