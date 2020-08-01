#### channel类型
channel是一种特殊的类型，一种引用类型，声明之后需要初始化才能使用

##### channel的声明
```Go
var 变量名 chan 元素类型
```

##### channel的初始化
使用`make`来初始化
(Go语言中需要使用make函数进行初始化的数据类型：slice, map, chan)
```Go
make(chan 元素类型, [缓冲大小])
```
缓冲大小选项是可选的

举几个例子：
```Go
ch1 := make(chan int)
```

##### channel的操作
通道有`发送send`、`接收receive`、`关闭close`三种操作

+ 发送
  ```Go
  ch <- 10
  ```
+ 接收
  ```Go
  x := <- ch  将channel中的数据赋值给x
  <- ch      从channel中接收数据，忽略结果
  ```
+ 关闭
  ```Go
  close(ch)
  ```
