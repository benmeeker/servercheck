package main

import (
	"log"
	"regexp"
	"servercheck/shared"
	"strings"
)

var new shared.Newversions

func Aptparse(infoCached shared.Info) map[string]shared.OPackage {
	var err error
	var arr []string
	var ext = make(map[string]string)
	var t = shared.OS{
		OS: infoCached.OS.Name,
	}
	for _, pack := range infoCached.URL {
		str := strings.Split(pack.Extensions, " ")
		for _, s := range str {
			ext[s] = ""
		}
	}
	t.Type = ext
	t.OS = strings.ReplaceAll(t.OS, " ", "")
	t.OS = strings.ReplaceAll(t.OS, "/", "")
	types[t.OS] = t
	arr, err = Read(t.OS, infoCached)
	if err != nil {
		log.Println(err)
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

func Outdatedrepos(host string) (map[string]shared.Outdatedrepos, int) {
	var outdated = make(map[string]shared.Outdatedrepos)
	var total int
	for _, ar := range pagedata.PageInfo[host].Aptrepos {
		for _, nv := range pagedata.PageInfo[host].Newversions {
			if strings.ToLower(strings.TrimSpace(nv.Name)) == strings.ToLower(strings.TrimSpace(ar.Name)) {
				if ar.Version != nv.Version {
					var pack = shared.Outdatedrepos{
						Name:       ar.Name,
						Oldversion: ar.Version,
						Newversion: nv.Version,
					}
					log.Println(pack.Newversion)
					outdated[pack.Name] = pack
					total++
				}
			}
		}
	}
	return outdated, total
}
