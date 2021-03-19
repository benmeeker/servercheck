package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"servercheck/shared"
	"strings"
	"time"
)

var types = make(map[string]shared.OS)

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
	f = strings.ReplaceAll(f, " ", "")
	f = strings.ReplaceAll(f, "/", "")
	return ioutil.WriteFile(f, os, 0666)
}

func Getdata(os, repo, url string) {
	err := Datapull(os+repo+".txt", url)
	if err != nil {
		log.Println(err)
	}
}

func Read(os string) error {
	for _, n := range types[os].Type {
		var a []string
		var s string
		out, err := ioutil.ReadFile(os + n + ".txt")
		if err != nil {
			log.Println(err)
			continue
		}
		scanner := bufio.NewScanner(strings.NewReader(string(out)))
		for scanner.Scan() {
			if scanner.Text() == "" {
				s = strings.Join(a, "\n")
				arr = append(arr, s)
				s = ""
				a = nil
				continue
			}
			a = append(a, string(scanner.Text()))
		}
	}
	return nil
}

func Outdated(infoCached shared.Info) bool {
	var out bool
	for n := range infoCached.Aptrepos {
		if key, ok := infoCached.Newversions[infoCached.Aptrepos[n].Name]; ok {
			if key.Version != infoCached.Aptrepos[n].Version {
				out = true
			}
		}
	}
	return out
}

func Links() {
	for {
		time.Sleep(10 * time.Second)
		for _, pi := range pagedata.PageInfo {
			pi.OS.Name = strings.ReplaceAll(pi.OS.Name, " ", "")
			pi.OS.Name = strings.ReplaceAll(pi.OS.Name, "/", "")
			fmt.Printf("%#v\n", types)
			for _, n := range types[pi.OS.Name].Type {
				var u = types[pi.OS.Name].URL + n + "/binary-amd64/Packages.gz"
				fmt.Println(u)
				Getdata(pi.OS.Name, n, u)
			}
		}
		time.Sleep(24 * time.Hour)
	}
}
