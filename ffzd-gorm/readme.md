# orm
- Object Relational Mapping 对象关系映射，解决了对象和关系型数据库之间的数据交互问题

## 模型定义
- 由GO的基本数据类型、实现了Scanner和Valuer接口自定义类型以及其指针或别名组成
- 字段标签
  - type    字段类型
  - size    字段大小
  - column  字段别名
  - primaryKey  将列定义为主键
  - unique  将列定义为唯一键
  - default 定义列的默认值
  - not null    不可为空
  - embedde 嵌套字段
  - comment 注释