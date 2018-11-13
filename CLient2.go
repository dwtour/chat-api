package main

import (
	"encoding/json"
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
		var dat[]interface{}
		json.Unmarshal(buf[:n], &dat)
		fmt.Println("another user: ", dat[0])
		fmt.Println("another user: ", dat[1])

	}
}
