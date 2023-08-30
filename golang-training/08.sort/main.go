package main

import (
	"fmt"
	"sort"
)

type people []string

// func (p people) Less(i, j int) bool {
// 	return p[i] < p[j]
// }
// func (p people) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }
// func (p people) Len() int {
// 	return len(p)
// }

func main() {
	studyGroup := people{"job", "work", "study"}
	numbers := []int{4, 2, 1, 3, 5, 6}
	fmt.Println(studyGroup)
	// sort.Sort(studyGroup)
	sort.Strings(studyGroup)
	sort.Ints(numbers)
	fmt.Println(studyGroup)
	fmt.Println(numbers)
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)
}
