package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")
	buf := make([]byte, 1024)
	duration := 5*time.Second
	time.Sleep(duration)
	defer conn.Close()
	conn.Write([]byte("1: hello guys\n"))
	//for i := 1; i < 10; i++ {

		//duration := time.Second
		//time.Sleep(duration)
	//}
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
	n, _ = conn.Read(buf)
	fmt.Println(string(buf[:n]))
	//time.Sleep(time.Second)
	conn.Write([]byte("1: ok\n"))
	n, _ = conn.Read(buf)
	fmt.Println(string(buf[:n]))

}
