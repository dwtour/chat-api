package TCPserver

import (
	"fmt"
	"io"
	"log"
	"net"
)

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
		go handleRequest(conn)
	}
}
//testing connection
func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf);
	for  err != io.EOF {
		fmt.Println(string(buf[:n]) + "(Message received).")
		n, err = conn.Read(buf);
	}
	fmt.Println("All messages received.")
	conn.Close()
}

