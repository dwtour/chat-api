package routers

import (
	"github.com/dwtour/chat-api/controllers"
	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()
	e.GET("/messages", controllers.GetHandler)
	e.POST("/message", controllers.PostHandler)
	e.GET("/ws", controllers.WSHandler)
	e.Start(":1323")
}
