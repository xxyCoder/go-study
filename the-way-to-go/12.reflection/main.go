package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	m := 1
	fmt.Println("Type", reflect.TypeOf(m))
	v := reflect.ValueOf(m)
	fmt.Println("Value", v)
	fmt.Println("Type", v.Kind())
	fmt.Println("Value", v.Int())
	fmt.Println("Interface", v.Interface())
	y := v.Interface().(int)
	fmt.Println(y)

	x := 3.14
	z := reflect.ValueOf(x)
	fmt.Println("settability of z", z.CanSet())
	z = reflect.ValueOf(&x)
	fmt.Println("settability of z", z.CanSet())
	z = z.Elem()
	fmt.Println("The Elm of z", z)
	fmt.Println("settability of z", z.CanSet())
	if z.CanSet() {
		z.SetFloat(3.13)
		fmt.Println("After set z", z)
		fmt.Println(z.Interface())
	}
	// 结构中只有被导出字段（首字母大写）才是可设置的
	t := T{23, "xxyCoder"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
