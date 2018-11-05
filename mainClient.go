package main

import (
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")
	for i := 1; i < 10; i++ {
		conn.Write([]byte("hello guys\n"))
		duration := time.Second
		time.Sleep(duration)
	}
	conn.Close()
}
