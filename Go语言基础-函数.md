##### 函数
```go
func 函数名(参数 参数类型)(返回值类型){
    函数体
}
```
##### defer语句
##### 匿名函数
##### 闭包
闭包=函数 + 外层变量的引用

##### 内置函数
`close`：用来关闭channel<br>
`len`：求长度<br>
`new`：分配值类型（int\struct）的内存，返回指针<br>
`make`：分配引用类型（map\slice\chan）的内存<br>
`append`：追加元素到数组和slice中<br>
`panic`和`recover`：用来进行错误处理<br>

##### panic和recover
`panic`可能在任何地方触发<br>
`recover`用于恢复错误现场，只能在defer中使用，defer要在可能触发错误的语句之前就定义<br>

