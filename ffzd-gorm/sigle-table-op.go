package main

import (
	"fmt"

	"gorm.io/gorm"
)

func AgeGt23(db *gorm.DB) *gorm.DB {
	return db.Where("age > 23")
}

func sto() {
	s1 := Student{Name: "xxyCoder", Age: 21}
	// create接收指针
	// err := DB.Create(&s1).Error
	// if err != nil {
	// 	fmt.Println("插入失败", err)
	// }
	// 批量插入
	var studentList []Student
	// for i := 0; i < 10; i++ {
	// 	studentList = append(studentList, Student{Name: fmt.Sprintf("xxy%d", i+1), Age: 22 + i})
	// }
	// err = DB.Create(&studentList).Error
	// if err != nil {
	// 	fmt.Println("批量插入失败", err)
	// }

	// 用指定字段创建记录
	// DB.Select("Name", "Age").Create(&s1)

	// 查询单条记录
	var s2 Student = Student{}

	// DB.Take(&s2)
	// DB.First(&s2)
	// DB.Last(&s2)

	// 根据主键查询
	// DB.Take(&s2, 2)

	// 根据其他字段查询
	// DB.Take(&s2, "name=?", "xxy6")

	// 根据结构体查，只能根据主键查
	// s2 = Student{Id: 6}
	// DB.Take(&s2)

	// 查询多条记录、获取查询记录数
	// count := DB.Find(&studentList).RowsAffected
	// fmt.Println(count)

	// 根据多个主键查
	studentList = []Student{}
	// DB.Find(&studentList, []int{2, 4, 6, 8, 10})

	// 根据其他条件查
	// DB.Find(&studentList, "name in ?", []string{"xxyCoder", "xxy1"})

	// 字段更新之save
	// DB.Take(&s2, 8)
	// s2.Name = "xxyCoder666"
	// s2.Age = 150

	// 改变的全部更新
	// DB.Save(&s2)

	// 选择字段更新，即使有多个字段也只更新选中的
	// DB.Select("name").Save(&s2)

	// 批量更新
	// DB.Find(&studentList, []int{1, 3, 5, 7}).Update("age", 0)

	// 批量更新多列
	// DB.Find(&studentList, []int{3, 5, 7, 9}).Updates(Student{Age: 50, Name: "xxx"})

	// 删除操作
	// DB.Delete(&s2, 10)

	// 批量删除
	// DB.Delete(&studentList, []int{4, 6, 8})

	// where查询
	// DB.Where("name = ?", "xxyCoder").Find(&studentList)
	// DB.Not("name = ?", "xxyCoder").Find(&studentList)
	// DB.Where("name in ?", []string{"xxyCoder", "xxx"}).Find(&studentList)

	// 模糊查询
	// DB.Where("name like ?", "x%").Find(&studentList)

	// 与查询
	// DB.Where("age > ?", "10").Where("name like ?", "x__").Find(&studentList)

	// 或查询
	// DB.Where("age > ?", "10").Or("name like ?", "x__").Find(&studentList)

	// select查询
	// DB.Select("name", "age").Find(&studentList)
	// type APIStudent struct {
	// 	Name string
	// }
	// select name from students
	// DB.Model(&Student{}).Find(&APIStudent{})

	// Between
	// DB.Where("age between ? AND ?", 10, 20).Find(&studentList)

	// 排序
	// DB.Order("age desc").Find(&studentList)

	// 分页
	// DB.Limit(2).Offset(0).Find(&studentList)
	// DB.Limit(2).Offset(2).Find(&studentList)
	// DB.Limit(2).Offset(4).Find(&studentList)
	// DB.Limit(2).Offset(6).Find(&studentList)

	// 去重
	// DB.Distinct("age").Select("name", "age", "id").Find(&studentList)

	// 分组
	var countList []int
	// DB.Model(Student{}).Select("count(id)").Group("age").Scan(&countList)

	// 原生sql
	// DB.Raw("select count(id) from student group by age").Scan(&countList)

	// 子查询
	// DB.Where("age > (?)", DB.Model(Student{}).Select("avg(age)")).Find(&studentList)

	// Scope
	var res []map[string]any

	DB.Model(Student{}).Scopes(AgeGt23).Find(&res)

	fmt.Println(s1)
	fmt.Println(res)
	fmt.Println(s2)
	fmt.Println(countList)
	fmt.Println(studentList)
}
