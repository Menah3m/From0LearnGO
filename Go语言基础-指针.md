##### 指针地址和指针类型
每个变量都拥有一个地址，这个地址代表该变量位置<br>
指针是一种特殊的变量--保存了其他变量的地址的变量<br>
```go
ptr := &v    //&操作符用于取地址，v的类型为T
```


##### new 和 make
用来分配内存使用<br>
new用于基本数据类型的分配，返回指针<br>
make只用于slice、map和chan类型，并返回该类型本身<br>