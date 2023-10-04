package main

import "fmt"

func mto() {
	// 一对多添加
	// DB.Create(&Author{
	// 	Name: "xxyCoder",
	// 	Articles: []Article{
	// 		{
	// 			Title: "go-study",
	// 		},
	// 		{
	// 			Title: "front_end-study",
	// 		},
	// 	},
	// })

	// a1 := Article{Title: "Git", AId: 1}
	// DB.Create(&a1)

	var a2 Author
	// DB.Take(&a2, 1)
	// DB.Create(&Article{Title: "project", Author: a2})

	// 给已存在作者添加文章
	var a3 Article
	// a3 = Article{Title: "note"}
	// DB.Model(&a2).Association("Articles").Append(&a3)

	// 预加载
	// DB.Preload("Author").Take(&a3)

	// 嵌套预加载
	// DB.Preload("Author.Articles").Take(&a3, 1)

	// 删除
	DB.Preload("Articles").Take(&a2, 1)
	DB.Model(&a2).Association("Articles").Delete(&a2.Articles)

	fmt.Println(a3)
}
