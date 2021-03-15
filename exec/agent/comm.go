package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"servercheck/shared"
)

var url string = "http://192.168.1.63:8080/upload"

func Jsoncomm() shared.Info {
	var info = shared.Info{
		Kernel:   Getkernel(),
		OS:       Getos(),
		Repos:    Aptrepos(),
		Hostname: Gethost(),
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
