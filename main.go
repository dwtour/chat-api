package main

import (
	"fmt"
	"github.com/dwtour/chat-api/db"
	"github.com/dwtour/chat-api/routers"
	_ "github.com/dwtour/chat-api/routers"
)


func main(){
	fmt.Println("Start connecting to database...")
	pong, err := db.Connect()
	if err != nil {
		fmt.Println(pong, err)
		panic("error connect")
	}
	fmt.Println("Connection is established")

	routers.Run()
	//TCPserver.Listen()
}
