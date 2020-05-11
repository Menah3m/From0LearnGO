#### 结构体定义
```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    ...
}

```
+ 类型名： 自定义结构体的名称，在同一个包内不可重复
+ 字段名： 结构体的字段名，结构体内的字段名必须唯一
+ 字段类型： 结构体字段的具体类型

例如：
```go
type person struct {
    name string
    age int
    city string
}
```