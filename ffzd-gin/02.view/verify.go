package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type SignInfo struct {
	Name            string   `json:"name" binding:"required"`
	Password        string   `json:"password" binding:"min=4,max=8"`
	ConfirmPassword string   `json:"confirmPassword" binding:"eqfield=Password"`
	Age             int      `json:"age" binding:"lt=30,gt=18"`
	LikeList        []string `json:"like_list" binding:"required,dive,startsWidth=like"`
	IP              string   `json:"ip" binding:"ip"`
	Url             string   `json:"url" binding:"url"`
	Date            string   `json:"date" binding:"datatime=2006-01-02 15:04:05"`
}

/*
required 必填字段
min	最小长度
max	最大长度
len	长度等于
eqfield	等于其他字段
nefield	不等于其他字段
-	忽略字段
startsWith
endsWith
contain
excludes
dive	后面验证针对数组每一个元素
*/

func signValid(fl validator.FieldLevel) bool {
	//...
	var namelist []string = []string{"xxy", "xxxy", "xxx"}
	for _, name := range namelist {
		v := fl.Field().Interface().(string)
		if v == name {
			return false
		}
	}
	return true
}

func verify() {
	router := gin.Default()
	// 自定义
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}
	router.POST("/", func(c *gin.Context) {
		var user SignInfo
		// 类型校验是自带的
		err := c.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"message": err})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})
	router.Run(":8080")
}
