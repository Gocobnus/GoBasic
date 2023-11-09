package main

import (
	"GoBasic/BasicSyntax/Net/TCP/StickyPackage/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)


func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return 
		}
		if err != nil {
			fmt.Printf("decode message failed, err : %v", err)
			return
		}
		fmt.Printf("receive client message : %v\n", msg)
	}
}


func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("listen failed, err : %v", err)
		return 
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err : %v", err)
			continue
		}
		go process(conn)
	}
}