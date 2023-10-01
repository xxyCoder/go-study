package main

import (
	"fmt"
	"reflect"
)

type T struct{ a, b int }

type struct1 struct {
	i1  int
	f1  float64
	str string
}

type Interval struct {
	start int
	end   int
}

// 带标签
type TagType struct {
	filed1 bool   "An important answer"
	filed2 string "The name of the string"
}

type innerS struct {
	int1 int
	int2 int
	b    int
}

// 匿名字段，此时类型就是字段名称，故每一种类型只能有一个匿名字段
type outterS struct {
	b int
	c int
	int
	// 内嵌结构体，可以模拟继承行为
	// 内嵌即将内层结构体简单插入外层结构体
	innerS
}

func (is innerS) addThem() int {
	return is.int1 + is.int2
}

func (is *innerS) addToParam(param int) int {
	return is.int1 + is.int2 + param
}

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring"
}

// 多重继承
type CameraPhone struct {
	Camera
	Phone
}

func main() {
	// 赋值方式
	var s1 T
	s1.a = 1
	s1.b = 1
	var s2 *T
	s2 = new(T)
	s2.a = 2
	s2.b = 2

	ms := new(struct1)
	ms.str = "xxyCoder learn go"
	ms = &struct1{10, 3.14, "xxy"}
	ms2 := struct1{17, 7.1, "17gaming"}
	fmt.Println(ms2)
	// 初始化方式
	intr1 := Interval{0, 3}
	intr2 := Interval{end: 5, start: 1}
	intr3 := Interval{end: 666}
	fmt.Println(intr1, intr2, intr3)
	// 带标签的结构体
	tt := TagType{true, "xxy"}
	for i := 0; i < 2; i++ {
		ttType := reflect.TypeOf(tt)
		ixField := ttType.Field(i)
		fmt.Printf("%v\n", ixField.Tag)
	}
	// 匿名字段
	outer := new(outterS)
	outer.b = 123
	outer.int = 666
	outer.innerS = innerS{1, 2, 3}
	// 命名冲突，外层覆盖内层，如果在同一个级别出现，则报错
	fmt.Println(outer, outer.b)
	// 调用方法
	fmt.Println(outer.addThem(), outer.innerS.addToParam(123))
	// 多重继承
	cp := new(CameraPhone)
	fmt.Println(cp.Call(), cp.TakeAPicture())
}
