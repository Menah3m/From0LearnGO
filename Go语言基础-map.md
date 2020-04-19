##### map定义
```go
map[keyType]valueType
```
map类型的变量默认初始值为nil,需要使用make()函数来分配内存。语法为：
```go
make(map[keyType]valueType, [cap])
```
其中cap表示map的容量，该参数不是必须的，但是在初始化map时就应该指定好

##### map添加k-v对
map中的数据都是成对出现的