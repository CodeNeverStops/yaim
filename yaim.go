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
	"github.com/lokizone/yaim/models"
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
		isStatic bool = true
	)

	path := r.URL.Path[1:]

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/html"
		isStatic = false
	}

	w.Header().Set("Content-Type", contentType + "; charset=utf-8")
	if isStatic {
		data, err := ioutil.ReadFile(PublicPath + "/" + string(path))
		if err != nil {
			w.WriteHeader(404)
		} else {
			w.Write(data)
		}
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		var ret bool
		var action string
		switch path {
			case "signin":
				ret = models.SignIn(r.PostForm["name"][0], r.PostForm["password"][0])
				action = "login"
			case "signup":
				ret = models.SignUp(r.PostForm["name"][0], r.PostForm["password"][0], r.PostForm["password_again"][0])
				action = "register"
		}
		var result string
		if (ret) {
			result = action + " success"
		} else {
			result = action + " failed"
		}
		w.WriteHeader(200)
		w.Write([]byte(result))
		return
	}

	if path == "" {
		path = "index"
	}
	template := templates.Lookup(path + ".html")
	if template == nil {
		w.WriteHeader(404)
		return
	}
	template.Execute(w, nil)
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

