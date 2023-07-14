package main

import (
	"encoding/json"
	"net/http"
)

type user struct {
	Id   int
	Name string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write([]byte("hello"))
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("author", "jingyang")
		//io.WriteString(writer, "{\"name\":\"zhangsan\"}")

		//struct --> json  json.encoder
		u := user{Id: 1, Name: "zhangsan"}
		json.NewEncoder(writer).Encode(u) //struct-->json
	})
	http.ListenAndServe("localhost:8080", nil)
}
