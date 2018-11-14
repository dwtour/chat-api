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
	Body string `json:"body"`
	IP string `json:"ip"`
	Hash string `json:"hash"`
}

type MessageJSON struct {
	Data []Message `json:"data"`
}

type PostMessage struct {
	Message string `json:"message"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		messages, _ := db.Conn.LRange("messages", 0, -1).Result()
		m := make([]Message, 0)
		for _, key := range messages {

			m = append(m, Message{
				Body: db.Conn.Get(key).Val(),
				IP: strings.Split(key, ":")[2],
				Hash: strings.Split(key, ":")[1]})
		}

		resp, _ := json.Marshal(&MessageJSON{Data: m})

		w.Write(resp)
	}  else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		req, _ := ioutil.ReadAll(r.Body)
		m := PostMessage{}
		json.Unmarshal(req, &m)
		fmt.Println(m)

		hash := uuid.NewV4().String()
		key := fmt.Sprintf("message:%s:%s", hash, r.Host)
		fmt.Println(r.Host)

		db.Conn.RPush("messages", key)
		db.Conn.Set(key, m.Message, 0)

		resp, _ := json.Marshal(&Message{
			Body: db.Conn.Get(key).Val(),
			IP: strings.Split(key, ":")[2],
			Hash: strings.Split(key, ":")[1]})

		w.Write(resp)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}


