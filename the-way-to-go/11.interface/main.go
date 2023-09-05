package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (cr *Circle) Area() float32 {
	return math.Pi * cr.radius * cr.radius
}

type Circle struct {
	radius float32
}

func (sq *Square) add() float32 {
	return sq.side + sq.side
}

type ReadWrite interface {
	Read(b int) bool
	Write(b int) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock()
	Close()
}

type Stringer interface {
	String() string
}

// 空接口
type Any interface{}

func classifier(item ...interface{}) {
	for i, x := range item {
		switch x.(type) {
		case bool:
			fmt.Println("param #%d is bool\n", i)
		case float64:
			fmt.Println("param #%d is a float64\n", i)
		case int, int64:
			fmt.Println("param #%d is int\n", i)
		case nil:
			fmt.Println("param #%d is nil\n", i)
		case string:
			fmt.Println("param #%d is string\n", i)
		default:
			fmt.Println("param #%d is unkown\n", i)
		}
	}
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Println(areaIntf.Area())
	// 类型断言
	if t, ok := areaIntf.(*Square); ok {
		fmt.Println(t)
	} else {
		fmt.Println("NO")
	}
	// type-switch
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Println("Type Square %T", t)
	case *Circle:
		fmt.Println("Type Circle %T", t)
	case nil:
		fmt.Println("nil %T", t)
	default:
		fmt.Println("unexpected %T", t)
	}
	classifier(13, -14.3, true, "xxyCoder", nil, false)
	// 测试一个值是否实现了某个接口
	var v Stringer
	if sv, ok := v.(Stringer); ok {
		fmt.Println(sv.String())
	}
}
