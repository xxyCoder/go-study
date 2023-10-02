package main

import (
	"fmt"
)

func main() {
	// s1 := Student{Name: "xxyCoder", Age: 21}
	// create接收指针
	// err := DB.Create(&s1).Error
	// if err != nil {
	// 	fmt.Println("插入失败", err)
	// }
	// 批量插入
	var studentList []Student
	for i := 0; i < 10; i++ {
		studentList = append(studentList, Student{Name: fmt.Sprintf("xxy%d", i+1), Age: 22 + i})
	}
	// err = DB.Create(&studentList).Error
	// if err != nil {
	// 	fmt.Println("批量插入失败", err)
	// }
	// 查询单条记录
	var s2 Student
	DB.Take(&s2)
	fmt.Println(s2)
	s2 = Student{}
	DB.First(&s2)
	fmt.Println(s2)
	s2 = Student{}
	DB.Last(&s2)
	fmt.Println(s2)
	// 根据主键查询
	s2 = Student{}
	DB.Take(&s2, 2)
	fmt.Println(s2)
	// 根据其他字段查询
	s2 = Student{}
	DB.Take(&s2, "name=?", "xxy6")
	fmt.Println(s2)
	// 根据结构体查，只能根据主键查
	s2 = Student{Id: 6}
	DB.Take(&s2)
	fmt.Println(s2)
	// 查询多条记录、获取查询记录数
	count := DB.Find(&studentList).RowsAffected
	fmt.Println(count)
	// 根据多个主键查
	studentList = []Student{}
	DB.Find(&studentList, []int{2, 4, 6, 8, 10})
	fmt.Println(studentList)
	// 根据其他条件查
	studentList = []Student{}
	DB.Find(&studentList, "name in ?", []string{"xxyCoder", "xxy1"})
	fmt.Println(studentList)
	// 字段更新之save
	s2 = Student{}
	DB.Take(&s2, 8)
	s2.Name = "xxyCoder666"
	s2.Age = 150
	// 改变的全部更新
	// DB.Save(&s2)
	// 选择字段更新，即使有多个字段也只更新选中的
	DB.Select("name").Save(&s2)
	// 批量更新
	studentList = []Student{}
	DB.Find(&studentList, []int{1, 3, 5, 7}).Update("age", 0)
	// 批量更新多列
	studentList = []Student{}
	DB.Find(&studentList, []int{3, 5, 7, 9}).Updates(Student{Age: 50, Name: "xxx"})
	// 删除操作
	s2 = Student{}
	DB.Delete(&s2, 10)
	fmt.Println(s2)
	// 批量删除
	studentList = []Student{}
	DB.Delete(&studentList, []int{4, 6, 8})
	fmt.Println(studentList)
}
