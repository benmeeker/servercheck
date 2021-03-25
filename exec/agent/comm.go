package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"servercheck/shared"
)

var url string = "http://localhost:8080/upload"

func Jsoncomm() shared.Info {
	var err error
	var info = shared.Info{
		Kernel:   Getkernel(),
		OS:       Getos(),
		Hostname: Gethost(),
	}
	info.Aptrepos = Getrepos(info.OS.Name)
	info.URL, err = Geturl(info.OS.Name)
	if err != nil {
		log.Println(err)
	}
	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(info)
	res, err := http.Post(url, "application/json; charset=utf-8", buff)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return info
}
