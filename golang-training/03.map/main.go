package main

import "fmt"

func main() {
	// 创建映射方法一
	var myGreeting = make(map[string]string)
	fmt.Println(myGreeting)
	myGreeting["Tom"] = "Cat"
	myGreeting["Jenny"] = "Mouse"
	fmt.Println(myGreeting)
	// 创建映射方法二
	myGreeting_2 := map[string]string{}
	myGreeting_2["s"] = "stop"
	// 创建并初始化
	myGreeting_3 := map[string]string{
		"Tom":   "Cat",
		"Jenny": "Mouse",
	}
	fmt.Println(len(myGreeting_3))
	for key, val := range myGreeting {
		fmt.Println(key, "-", val)
	}
	// rune
	word := "abcdefghijklmnopqrstuvwxyz"
	for i, j := range word {
		fmt.Println(string(j), rune(word[i]))
	}
}
