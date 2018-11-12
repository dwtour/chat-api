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
		//time.Sleep(time.Second)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error - %s", err)
			break
		}
		fmt.Println("another user: ", string(buf[:n]))

	}
}
