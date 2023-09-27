package main

import (
	"fmt"
	"strconv"
	"strings"
)

func String() {
	// 非解释类型字符串
	str1 := `This is a raw string\n`
	fmt.Println(str1)
	// 获取字符串某个字节地址行为是非法的
	// &str1[0]
	// 前缀和后缀
	str2 := "Hello xxyCoder"
	fmt.Println(strings.HasPrefix(str2, "Hello"), strings.HasSuffix(str2, "Coder"))
	// 字符串的包含关系
	fmt.Println(strings.Contains(str2, "xxy"))
	// 判断子串或字符在父字符串中出现的位置
	fmt.Println(strings.Index(str2, "xxy"))
	// 字符串替换
	fmt.Println(strings.Replace(str2, "Hello", "Hi", -1), str2)
	// 统计出现次数
	fmt.Println(strings.Count(str2, "x"))
	// 改变大小写
	fmt.Println(strings.ToLower(str2), strings.ToUpper(str2))
	// 分割字符串
	fmt.Println(strings.Split(str2, " "))
	// 拼接字符串
	fmt.Println(strings.Join([]string{"Hi", "xxyCoder"}, " "))
	// 类型转换
	var str3 = "666"
	fmt.Println(strconv.Atoi(str3))
}
