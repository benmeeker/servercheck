package main

import (
	"io/ioutil"
)

func Aptparse() {
	switch infoCached.OS.Name {
	case "Kali GNU/Linux":
		ioutil.ReadFile("kalimain.txt")
	}

}
