package routers

import (
	"github.com/dwtour/chat-api/controllers"
	"github.com/labstack/echo"
)

func Run() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/messages", controllers.GetHandler)
	//mux.HandleFunc("/message", controllers.PostHandler)
	//go http.ListenAndServe(":3000", mux)

	e := echo.New()
	e.GET("/messages", controllers.GetHandler)
	e.POST("/message", controllers.PostHandler)
	e.Start(":1323")
}
