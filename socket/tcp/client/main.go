package main

// tcp client demo

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"strings"
)

func main()  {
	// 1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil{
		fmt.Printf("Dial failed, err:%v\n", err)
		return 
	}
	// 2。利用该连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q"{
			return 
		}
		// 3.给服务端发消息
		_, err = conn.Write([]byte(s))
		if err != nil{
			fmt.Printf("Send failed, err:%v\n", err)
			return 
		}
	    // 4.从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil{
			fmt.Printf("read failed, err:%v\n", err)
			return 
		}
		fmt.Println("收到服务端回复： ", string(buf[:n]))
	}
}
