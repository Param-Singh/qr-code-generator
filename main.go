package main

import (
	"fmt"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("response"))
}

func main() {
	http.HandleFunc("/generate", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server Listening at port 8080")
}
