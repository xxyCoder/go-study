package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 模型定义
type Student struct {
	Id   uint   `gorm:"size:10;primaryKey;column:id"`
	Name string `gorm:"size:16;column:name"`
	Age  int    `gorm:"size:3:column:age"`
}

// Hook
func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
	user.Name = "xxx"
	user.Age = 0
	return nil
}

var DB *gorm.DB

func init() {
	fmt.Println("init")
	username := "root"
	password := "?"
	host := "127.0.0.1"
	port := 3306
	dbName := "gorm"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s", username, password, host, port, dbName, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 表名前缀
			SingularTable: true,  // 表名单数
			NoLowerCase:   false, //关闭小写转换
		},
	})
	if err != nil {
		panic("连接数据库失败,err = " + err.Error())
	}
	fmt.Println("连接数据库成功", db)
	DB = db
}

func Conn() {
	// 创建数据表
	DB.AutoMigrate(&Student{})
}
