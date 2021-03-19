package main

import (
	"log"
	"regexp"
	"servercheck/shared"
	"strings"
)

var arr []string
var new shared.Newversions

func Aptparse(infoCached shared.Info) map[string]shared.OPackage {
	switch infoCached.OS.Name {
	case "Kali GNU/Linux":
		var t = shared.OS{
			OS:  infoCached.OS.Name,
			URL: "http://mirrors.jevincanders.net/kali/dists/kali-rolling/",
		}
		t.OS = strings.ReplaceAll(t.OS, " ", "")
		t.OS = strings.ReplaceAll(t.OS, "/", "")
		t.Type = []string{"main", "contrib", "non-free"}
		types[t.OS] = t
		err := Read(t.OS)
		if err != nil {
			log.Println(err)
		}
	case "Ubuntu":
		var t = shared.OS{
			OS:  infoCached.OS.Name,
			URL: "http://us.archive.ubuntu.com/ubuntu/dists/focal-updates/",
		}
		t.OS = strings.ReplaceAll(t.OS, " ", "")
		t.OS = strings.ReplaceAll(t.OS, "/", "")
		t.Type = []string{"main", "multiverse", "restricted", "universe"}
		types[t.OS] = t
		err := Read(t.OS)
		if err != nil {
			log.Println(err)
		}
	}
	var new = make(map[string]shared.OPackage)
	for i := range arr {
		txt := arr[i]
		re := regexp.MustCompile(`Package: (.+)[\S\s\n]+Version: (.+)`)
		n := re.FindAllStringSubmatch(txt, i)
		if len(n) > 0 && len(n[0]) > 2 {
			var pack = shared.OPackage{
				Name:    strings.ToLower(strings.TrimSpace(n[0][1])),
				Version: n[0][2],
			}
			new[pack.Name] = pack
		}
	}
	return new
}

func Outdatedrepos(host string) map[string]shared.Outdatedrepos {
	var outdated = make(map[string]shared.Outdatedrepos)
	for _, ar := range pagedata.PageInfo[host].Aptrepos {
		for _, nv := range pagedata.PageInfo[host].Newversions {
			if nv.Name == strings.ToLower(strings.TrimSpace(ar.Name)) {
				if ar.Version != nv.Version {
					var pack = shared.Outdatedrepos{
						Name:       ar.Name,
						Oldversion: ar.Version,
						Newversion: nv.Version,
					}
					outdated[pack.Name] = pack
				}
			}
		}
	}
	return outdated
}
