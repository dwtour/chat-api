package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)
var (
	connections = make(map[string]*websocket.Conn)
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func WSHandler(c echo.Context) error {
	//loop or return?
	for {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		//defer ws.Close()

		connections[ws.RemoteAddr().String()] = ws
	}
}

