package main

import "fmt"

func main() {
	var arr1 [5]int
	// 数组常量
	var arr2 = [5]string{"xxyCoder", "xxy", "666"}
	var arr3 = [...]int{1, 2, 3, 4, 5}
	var arr4 = []int{1, 2, 3, 4}
	var arr5 = []int{3: 1, 5: 4}
	fmt.Println(len(arr1), len(arr2), len(arr3), len(arr4), len(arr5))
	fmt.Println(arr1, arr2, arr3, arr4, arr5)
	// 多维数组
	var arr6 [3][3]int
	for i := 0; i < len(arr6); i++ {
		for j := 0; j < len(arr6[i]); j++ {
			arr6[i][j] = i * j
			fmt.Print(arr6[i][j], " ")
		}
		fmt.Println()
	}
	// 切片 [star:end)
	var slice1 []int = arr1[2:4]
	arr1[2] = 10
	fmt.Println(len(slice1), cap(slice1), slice1[0])
	// make创建切片 make([]type,len,cap)
	var slice2 []int = make([]int, 10, 20)
	fmt.Println(len(slice2), cap(slice2))
	// new方式
	var slice3 []int = new([100]int)[0:50]
	fmt.Println(len(slice3), cap(slice3))
	// 多维切片
	var slice4 [][]int = make([][]int, 10, 20)
	for i := 0; i < len(slice4); i++ {
		slice4[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			slice4[i][j] = i + j
		}
	}
	fmt.Println(slice4)
	// 切片重组
	for i := 0; i < cap(slice2); i++ {
		slice2 = slice2[0 : i+1]
		slice2[i] = i
		fmt.Println("reslice:", len(slice2))
	}
	// 切片的复制与追加
	slFrom := []int{1, 2, 3}
	slFrom = append(slFrom, 4, 5, 6)
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo, "copy len:", n)
}
