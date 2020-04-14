##### Go项目结构

在进行Go语言开发时，我们的代码保存在`$GOPATH/src`目录下。

在工程经过`go build`、`go install`或`go get`等命令后，会将下载的第三方源文件存放在`$GOPATH/src`目录下，产生的二进制可执行文件放在`$GOPATH/bin`目录下，生成的中间缓存文件放在`$GOPATH/pkg`目录下。



##### 版本控制

`$GOPATH/src/github.com/个人用户名/代码`