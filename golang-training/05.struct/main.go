package main

import (
	"encoding/json"
	"fmt"
)

// 别名
type Person struct {
	first       string `json:"-"`
	last        string
	age         int `json:"wisdom score"`
	notExported int
}

// 结构体定义方法
func (p Person) fullName() string {
	return p.first + p.last
}

// 嵌套结构体
type doubleZero struct {
	person        Person
	LicenseToKill bool
}

func main() {
	p1 := Person{"xxy", "Coder", 21, 007}
	fmt.Println(p1.fullName())
	p2 := doubleZero{
		person: Person{
			first:       "666",
			last:        "777",
			age:         67,
			notExported: 0,
		},
		LicenseToKill: true,
	}
	fmt.Println(p2)
	// 使用指针
	p3 := &Person{
		first:       "888",
		last:        "999",
		age:         89,
		notExported: 0,
	}
	fmt.Println(p3.first)
	bs, w := json.Marshal(p1)
	fmt.Println(bs, w, string(bs))
}
