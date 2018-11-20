package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"net/http"
)
var (
	connections = make(map[string]*websocket.Conn)
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func WSHandler(c echo.Context) error {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if ws != nil {
		connections[ws.RemoteAddr().String()] = ws
	}
	for {
		ws.ReadMessage()
	}

}

