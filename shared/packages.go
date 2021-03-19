package shared

type RPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
}

type OPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Info struct {
	Kernel        string                   `json:"kernel"`
	OS            OPackage                 `json:"os"`
	Hostname      string                   `json:"hostname"`
	Aptrepos      []RPackage               `json:"aptrepos"`
	Newversions   map[string]OPackage      `json:"newversions"`
	Outdated      bool                     `json:"outdated"`
	Outdatedrepos map[string]Outdatedrepos `json:"outdatedrepos"`
}

type PageData struct {
	PageInfo map[string]Info `json:"pageinfo"`
	Path     string          `json:"path"`
	Pagename string          `json:"pagename"`
	Hostname string          `json:"hostname"`
}

type OS struct {
	OS   string   `json:"os"`
	Type []string `json:"type"`
	URL  string   `json:"url"`
}

type Newversions struct {
	Aptversions []OPackage `json:"aptversions"`
}

type Outdatedrepos struct {
	Name       string `json:"name"`
	Oldversion string `json:"oldversion"`
	Newversion string `json:"newversion"`
}
