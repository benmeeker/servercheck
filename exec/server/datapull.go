package main

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"servercheck/shared"
	"time"
)

var types = shared.OS{
	Type: []string{"main", "contrib", "nonfree"},
}

func Datapull(f string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	os, err := ioutil.ReadAll(gz)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f, os, 0666)
}

func Getdata() {
	for {
		for _, n := range types.Type {
			Datapull(types.OS+n+".txt", "http://mirrors.jevincanders.net/kali/dists/kali-rolling/"+n+"/binary-amd64/Packages.gz")
		}
		time.Sleep(24 * time.Hour)
	}
}
