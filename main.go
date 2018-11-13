package main

import (
	"fmt"
	"github.com/dwtour/chat-api/TCPserver"
	"github.com/dwtour/chat-api/db"
)


func main(){
	fmt.Println("Start connecting to database...")
	pong, err := db.Connect()
	if err != nil {
		fmt.Println(pong, err)
		panic("error connect")
	}
	fmt.Println("Connection is established")
	TCPserver.Listen()
}
