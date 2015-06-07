package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
var indexTpl = template.Must(template.ParseFiles("templates/index.html"))

func serveHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	indexTpl.Execute(w, r.Host)
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serve ws")
}

func main() {
	fmt.Println("yaim")
	flag.Parse()
	http.HandleFunc("/", serveHttp)
	http.HandleFunc("/ws", serveWS)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
