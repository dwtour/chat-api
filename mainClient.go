package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//net.Dial("tcp", ":8080")
	//buf := make([]byte, 1024)
	//defer conn.Close()
	//conn.Write([]byte("hello guys"))
	resp, err := http.Get("localhost:3000/messages/")
	if err != nil {
		fmt.Println(err.Error())
	}
	temp, _ := ioutil.ReadAll(resp.Body)
	var dat []map[string]interface{}
	json.Unmarshal(temp, &dat)
	fmt.Println(resp)
}
