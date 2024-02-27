package main

import (
	"log"
	"net/http"
)

func main() {  //handles excecutables
	http.HandleFunc("/",func(http.ResponseWriter,*http.Request){ 
		log.Println("hello world")
	})
	http.ListenAndServe(":9090",nil) // defaultservemux (http handler- handler is an interface) - nil
}