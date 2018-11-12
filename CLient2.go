package main

import (
	"fmt"
	"net"
)

func main(){
	conn, _ := net.Dial("tcp", ":8080")
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		fmt.Println("another user: ", string(buf[:n]))
	}
}
