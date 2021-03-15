package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"servercheck/shared"
)

var pagedata = shared.PageData{
	PageInfo: make(map[string]shared.Info),
}

func Server() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/upload", Upload)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("connection from %s\n", r.RemoteAddr)
	var infoCached shared.Info
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &infoCached)
	if err != nil {
		log.Fatal(err)
	}
	pagedata.PageInfo[infoCached.Hostname] = infoCached
}

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		pagedata.Path = "/"
		pagedata.Pagename = "Home"
	case "/allservers":
		pagedata.Path = "/allservers"
		pagedata.Pagename = "All Servers"
	default:
		http.NotFound(w, r)
		return
	}
	fmt.Println(pagedata.Path)
	tem, err := template.ParseGlob("*.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = tem.ExecuteTemplate(w, "index", pagedata)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
