package main

import (
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")
	//buf := make([]byte, 1024)
	defer conn.Close()
	conn.Write([]byte("1: hello guys\n"))

}
