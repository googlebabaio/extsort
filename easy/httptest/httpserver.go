package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"hhhhhh,hello,world%s",request.FormValue("name"))
	})
	http.ListenAndServe(":8888",nil)
}