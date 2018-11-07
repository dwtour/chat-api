package main

import (
	"fmt"
	"net"
)

func main(){
	conn, _ := net.Dial("tcp", ":8080")
	buf := make([]byte, 1024)
	defer conn.Close()
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
	conn.Write([]byte("2: how are you\n"))
	n, _ = conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
