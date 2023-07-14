package main

import (
	"io"
	"net/http"
)

type handleHello struct {
	content string
}

func (h *handleHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, h.content)
}
func main1() {
	//pattern url path
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello go server"))
	},
	)
	http.Handle("/hello", &handleHello{content: "handle hello"})

	//server listen port
	//http 1.1
	http.ListenAndServe(":8080", nil)
	//https ssl CA
	//http.ListenAndServeTLS()
}
