package main

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
	"servercheck/shared"
	"strconv"
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
	log.Println("Writing file", f)
	return ioutil.WriteFile("txtfiles/"+f, os, 0666)
}

func Getdata(os, repo, url string, num string) {
	log.Println("Getting data")
	err := Datapull(num+os+repo+".txt", url)
	if err != nil {
		log.Println(err)
	}
}

func Read(os string, infoCached shared.Info) ([]string, error) {
	log.Println("reading files")
	var arr []string
	var count int
	for n := range types[os].Type {
		var a []string
		var s string
		c := strconv.Itoa(count)
		out, err := ioutil.ReadFile("txtfiles/" + c + os + n + ".txt")
		log.Println(c + os + n + ".txt")
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
		count++
	}
	return arr, nil
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
			for nu, u := range pi.URL {
				for i := range types[pi.OS.Name].Type {
					var ur = u.URL + "/" + "/dists/" + u.Repo + "/" + i + "/binary-amd64/Packages.gz"
					num := strconv.Itoa(nu)
					Getdata(pi.OS.Name, i, ur, num)
				}
			}
		}
		time.Sleep(24 * time.Hour)
	}
}
