package main

import (
	"bufio"
	"fmt"
	"net"
)


func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err : %v\n", err)
		return 
	}
	// 这个for循环一直在监听20000端口
	for {
		fmt.Println("waiting connect")
		conn, err := listen.Accept()
		// 查看服务端端口号
		fmt.Println(conn.LocalAddr().String())
		if err != nil {
			fmt.Printf("accept failed, err : %v\n", err)
			continue
		}
		fmt.Println("begin process")
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	// 这个for循环一直在处理client的请求
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n ,err := reader.Read(buf[:])
		fmt.Println(n)
		if err != nil {
			fmt.Printf("read from client failed, err : %v\n", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Printf("receive from client is : %v\n", recvStr)
		conn.Write([]byte(recvStr))
	}
	fmt.Println("close conn")
}