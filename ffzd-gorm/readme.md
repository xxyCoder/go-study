# orm
- Object Relational Mapping 对象关系映射，解决了对象和关系型数据库之间的数据交互问题

# 模型定义
- 由GO的基本数据类型、实现了Scanner和Valuer接口自定义类型以及其指针或别名组成
- 字段标签
  - type    字段类型
  - size    字段大小
  - column  字段别名
  - primaryKey  将列定义为主键
  - autoIncrement 自动增长
  - unique  将列定义为唯一键
  - default 定义列的默认值
  - not null    不可为空
  - embedde 嵌套字段
  - comment 注释
  - embeded 将结构体嵌入
  - embededPrefix 对嵌入的字段添加前缀
  - <-  设置字段写入权限
  - ->  设置字段读取权限
  - - 省略该字段
  - foreignKey  重写外键名
  - references  重写引用关系
  - constraint  实现外键约束

## 约定
- 默认情况下：使用ID作为主键；结构体名蛇形复数作为表名；字段名蛇形作为列名
- GORM定义了一个gorm.Model结构体
```go
type Model struct {
  ID  uint  `gorm:"primaryKey"`
  CreatedAt time.time
  UpdatedAt time.time
  DeletedAt gorm.DeletedAt  `gorm:"index"`
}
```
- 对于多表，表名+ID被隐含表示两表之间的外键，且需要在定义外键的表添加其外键表名的结构体名