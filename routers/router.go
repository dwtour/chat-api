package routers

import (
	"fmt"
	"github.com/dwtour/chat-api/controllers"
	"net/http"
)


func init() {
	fmt.Println("Start serving mux")
	mux := http.NewServeMux()
	mux.HandleFunc("/messages", controllers.GetHandler)
	mux.HandleFunc("/message", controllers.PostHandler)
	fmt.Println("Mux is being served")
}
