package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)

const (
	TemplatePath string = "templates"
	PublicPath string = "public"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	templates *template.Template
)

func main() {
	flag.Parse()

	templates = parseTemplates()

	http.HandleFunc("/", serveHttp)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHttp(w http.ResponseWriter, r *http.Request) {
	var (
		contentType string
		isTpl bool
	)

	path := r.URL.Path[1:]

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/html"
		isTpl = true
	}

	w.Header().Set("Content-Type", contentType + "; charset=utf-8")
	if isTpl {
		if path == "" {
			path = "index"
		}
		template := templates.Lookup(path + ".html")
		if template == nil {
			w.WriteHeader(404)
			return
		}
		template.Execute(w, nil)
	} else {
		data, err := ioutil.ReadFile(PublicPath + "/" + string(path))
		if err != nil {
			w.WriteHeader(404)
			return
		}
		w.Write(data)
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serve ws")
}

func parseTemplates() *template.Template {
	result := template.New("templates")

	templateFolder, _ := os.Open(TemplatePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, TemplatePath + "/" + pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)
	return result
}

