package routers

import (
	"github.com/dwtour/chat-api/controllers"
	"net/http"
)


func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/messages", controllers.GetHandler)
	mux.HandleFunc("/message", controllers.PostHandler)
	go http.ListenAndServe(":3000", mux)
}
