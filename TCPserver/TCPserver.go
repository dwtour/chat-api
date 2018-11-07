package TCPserver

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type SafeConnect struct{
	list map[string]net.Conn
	mux sync.Mutex
}

var listConnect = SafeConnect{list:make(map[string]net.Conn)}

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

		listConnect.list[conn.RemoteAddr().String()] = conn
		go handleRequest(conn)
	}
}
//testing connection
func handleRequest(conn net.Conn) {
	chMessage := make(chan []byte, 100)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Read error - %s\n", err)
			}
			break
		}
		chMessage <- buf[:n]

		go sendMessage(chMessage, conn)
		//fmt.Println(string(buf[:n]) + "(Message received).")
	}
	delete(listConnect.list, conn.RemoteAddr().String())
	fmt.Println("All messages received.")

	conn.Close()
}

func sendMessage(ch chan []byte, conn net.Conn) {
	listConnect.mux.Lock()
	mes := <- ch
	for _, v := range listConnect.list {
		if (v != conn) {
			v.Write(mes)
		}
	}
	listConnect.mux.Unlock()
}