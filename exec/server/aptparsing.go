package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"servercheck/shared"
)

var mash []byte

func Aptparse() {
	switch infoCached.OS.Name {
	case "Kali GNU/Linux":
		types.OS = "kali"
		Read("kali")
	}
	var pack shared.OPackage
	re := regexp.MustCompile(`Package: (.+)[\S\s\n]+Version: (.+)`)
	n := re.FindAllStringSubmatch(string(mash), 1)
	if len(n) > 0 && len(n[0]) > 2 {
		pack.Name = n[0][1]
		pack.Version = n[0][2]
		log.Println(pack)
	}
}

func Read(os string) error {
	for _, n := range types.Type {
		out, err := ioutil.ReadFile(os + n + ".txt")
		if err != nil {
			log.Println(err)
			continue
		}
		mash = append(mash, out...)
	}
	return nil
}
