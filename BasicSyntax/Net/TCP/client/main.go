package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("cllient err : %v\n", err)
		return 
	}
	// 查看客户端端口号
	fmt.Println(conn.LocalAddr().String())
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return 
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("receive from server failed. err : %v\n", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}