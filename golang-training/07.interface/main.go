package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// 接受任意参数
func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	s := square{side: 10}
	c := circle{5}
	info(s)
	fmt.Println(totalArea(s, &c))
	specs(s)
	specs(c)
	specs([]int{1, 2, 3})
	// 接口数组
	critters := []interface{}{c, s, []int{1, 2, 3}}
	specs(critters)
	// 指针
	info(&c)
}
