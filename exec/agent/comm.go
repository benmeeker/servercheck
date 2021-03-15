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
	var info = shared.Info{
		Kernel:   Getkernel(),
		OS:       Getos(),
		Hostname: Gethost(),
		Aptrepos: Allaptrepos(),
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
