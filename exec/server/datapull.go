package main

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Datapull(f string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	gz, _ := gzip.NewReader(resp.Body)
	os, _ := ioutil.ReadAll(gz)
	ioutil.WriteFile(f, os, 0666)
}

func Getdata() {
	for {
		Datapull("kalimain.txt", "http://mirrors.jevincanders.net/kali/dists/kali-rolling/main/binary-amd64/Packages.gz")
		Datapull("kalicontrib.txt", "http://mirrors.jevincanders.net/kali/dists/kali-rolling/contrib/binary-amd64/Packages.gz")
		Datapull("kalinonfree.txt", "http://mirrors.jevincanders.net/kali/dists/kali-rolling/non-free/binary-amd64/Packages.gz")
		time.Sleep(24 * time.Hour)
	}
}
