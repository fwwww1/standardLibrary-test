package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main2() {
	//url path Methods HTTP1.1 GET POST
	//user?id=1&name=""
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "query"+request.URL.Query().Get("id"))
	})

	//add user post entype JSON
	http.HandleFunc("/user/add", func(writer http.ResponseWriter, request *http.Request) {
		//post data entype json
		var params map[string]string
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&params)
		io.WriteString(writer, "post json:"+params["name"])

	})
	//post data form-encoded
	http.HandleFunc("/user/del", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		io.WriteString(writer, "form:"+request.Form.Get("name"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
