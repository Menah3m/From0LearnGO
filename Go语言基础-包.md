##### 导入包时的执行顺序
1. 全局声明（变量等的赋值）
2. init()
3. main()

##### init()函数的执行顺序
后导入的包 先执行init()函数


##### 注意事项
1. 导入的包必须要使用，否则会报错
2. 不能循环导入（即互相导入）