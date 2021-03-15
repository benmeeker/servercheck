package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"servercheck/shared"
	"strings"
)

func Getkernel() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func Getos() shared.OPackage {
	out, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	var opackages shared.OPackage
	transfer := string(out)
	os := regexp.MustCompile(`\bNAME="(.+)"[\S\s\n]+VERSION="(.+)"`)
	n := os.FindAllStringSubmatch(transfer, 1)
	if len(n) > 0 && len(n[0]) > 2 {
		opackages.Name = n[0][1]
		opackages.Version = n[0][2]
	}
	return opackages
}

func Gethost() string {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return host
}

func Allaptrepos() []shared.RPackage {
	var array []string
	out, err := exec.Command("dpkg", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	var rpackages []shared.RPackage
	for n := range array {
		re := regexp.MustCompile(`  (\S+)\s+(\S+)\s+(\S+)\s+`)
		txt := array[n]
		each := re.FindAllStringSubmatch(txt, n)
		if len(each) == 0 {
			continue
		}
		var pkg = shared.RPackage{
			Name:    each[0][1],
			Version: each[0][2],
			Arch:    each[0][3],
		}
		rpackages = append(rpackages, pkg)
	}
	return rpackages
}
