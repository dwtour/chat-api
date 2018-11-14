package TCPserver

import (
	"encoding/json"
	"fmt"
	"github.com/dwtour/chat-api/db"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"net"
	"strings"
)

var connections = make(map[string]net.Conn)

type Message struct {
	Body string
	IP string
	Hash string
}

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

		connections[conn.RemoteAddr().String()] = conn
		getMessages(conn)

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
		key := fmt.Sprintf("message:%s:%s", hash, conn.RemoteAddr().String())

		db.Conn.RPush("messages", key)
		db.Conn.Set(key, buf[:n], 0)
		fmt.Println(db.Conn.Get(key))

		go sendMessage(buf[:n], conn)
	}
	delete(connections, conn.RemoteAddr().String())
	fmt.Printf("User %s disconnected.\n", conn.RemoteAddr().String())

	conn.Close()
}

func sendMessage(mes []byte, conn net.Conn) {
	for _, v := range connections {
		if v != conn {
			v.Write(mes)
		}
	}
}

func getMessages(conn net.Conn) {
	messages, _ := db.Conn.LRange("messages",0,-1).Result()
	output := make([]Message, 0)
	for _, key:= range messages {
		mes:= db.Conn.Get(key).Val()
		partsKey := strings.Split(key, ":")
		fmt.Println("parts of key", partsKey)
		output = append(output, Message{Body: mes, IP: partsKey[2], Hash: partsKey[1]})
		fmt.Printf("%s %s is sent\n", key, mes)
	}
	fmt.Println(output)
	messagesJSON, _ := json.Marshal(output)
	fmt.Println(string(messagesJSON))
	conn.Write(messagesJSON)
}