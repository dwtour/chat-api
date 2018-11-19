package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dwtour/chat-api/db"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
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

func GetHandler(c echo.Context) error {
	messages, _ := db.Conn.LRange("messages", 0, -1).Result()
	m := make([]Message, 0)
	for _, key := range messages {
		m = append(m, Message{
			Body: db.Conn.Get(key).Val(),
			IP: strings.Split(key, ":")[2],
			Hash: strings.Split(key, ":")[1]})
		}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&MessageJSON{Data: m})
}

func PostHandler(c echo.Context) error {
	m := new(PostMessage)
	if err := c.Bind(m); err != nil {
		delete(connections, c.Request().RemoteAddr)
		return err
	}

	hash := uuid.NewV4().String()
	key := fmt.Sprintf("message:%s:%s", hash, c.Request().Host)

	db.Conn.RPush("messages", key)
	db.Conn.Set(key, m.Message, 0)
	sendMessage(&Message{
		Body: m.Message,
		IP: c.Request().Host,
		Hash: hash})
	return nil
}

func sendMessage(resp *Message) {
	// error handler?
	for _, conn := range connections {
		err := conn.WriteJSON(resp)
		if err != nil {
			delete(connections, conn.RemoteAddr().String())
			continue
		}
	}
}

