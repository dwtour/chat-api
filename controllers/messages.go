package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dwtour/chat-api/db"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

type Message struct {
	Body string
	IP string
	Hash string
}


func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		messages, _ := db.Conn.LRange("messages", 0, -1).Result()
		output := make([]Message, 0)
		for _, key := range messages {
			mes := db.Conn.Get(key).Val()
			partsKey := strings.Split(key, ":")
			fmt.Println("parts of key", partsKey)

			output = append(output, Message{Body: mes, IP: partsKey[2], Hash: partsKey[1]})
			fmt.Printf("%s %s is sent\n", key, mes)
		}
		messagesJSON, _ := json.Marshal(output)

		w.Write(messagesJSON)
	}  else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		message, _ := ioutil.ReadAll(r.Body)

		hash := uuid.NewV4().String()
		key := fmt.Sprintf("message:%s:%s", hash, r.URL.Host)
		fmt.Println("host", r.URL.Host)

		db.Conn.RPush("messages", key)
		db.Conn.Set(key, message, 0)

		mes := db.Conn.Get(key).Val()
		partsKey := strings.Split(key, ":")

		output:=  &Message{Body: mes, IP: partsKey[2], Hash: partsKey[1]}
		MessageJSON, _ := json.Marshal(output)
		w.Write(MessageJSON)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}


