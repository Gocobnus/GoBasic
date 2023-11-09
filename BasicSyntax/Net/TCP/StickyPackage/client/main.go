package main

import (
	"GoBasic/BasicSyntax/Net/TCP/StickyPackage/proto"
	"fmt"
	"net"
)


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Printf("dial failed, err : %v", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "my name is gocobnus"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Printf("encode failed, err : %v", err)
			return 
		}
		conn.Write(data)
	}
}