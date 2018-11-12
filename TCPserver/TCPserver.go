package TCPserver

import (
	"fmt"
	"github.com/dwtour/chat-api/db"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"net"
)

var listConnect = make(map[string]net.Conn)

func Listen(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting TCP Listener.")
	}
	defer listener.Close()
	fmt.Println("Listening")
	for {
		conn, err:= listener.Accept()
		if err != nil {
			log.Fatal("Error accepting", err.Error())
		}

		listConnect[conn.RemoteAddr().String()] = conn
		ShowAllMessagesToNewUser(conn)

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Read error - %s\n", err)
			}
			break
		}
		hash := uuid.NewV4().String()
		//fmt.Println(hash)

		db.Conn.RPush("messages", hash)
		key := "message:"+hash+":"+conn.RemoteAddr().String()
		db.Conn.Set(key, buf, 0)
		fmt.Println(db.Conn.Get(key))

		go sendMessage(buf[:n], conn)
	}
	delete(listConnect, conn.RemoteAddr().String())
	fmt.Println("User ", conn.RemoteAddr().String()," disconnected.")

	conn.Close()
}

func sendMessage(mes []byte, conn net.Conn) {
	for _, v := range listConnect {
		if (v != conn) {
			v.Write(mes)
		}
	}
}

func ShowAllMessagesToNewUser(conn net.Conn) {
	listMessage, _ := db.Conn.LRange("messages",0,-1).Result()
	for _, v:= range listMessage {
		key := db.Conn.Keys("message:"+v+"*").Val()[0]
		mes, _:= db.Conn.Get(key).Bytes()
		conn.Write(mes)
		fmt.Println("message ", key, string(mes), "is sent")
	}
}