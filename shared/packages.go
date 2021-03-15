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
	Kernel   string     `json:"kernel"`
	OS       OPackage   `json:"os"`
	Repos    []RPackage `json:"repos"`
	Hostname string     `json:"hostname"`
}

type PageData struct {
	PageInfo map[string]Info `json:"pageinfo"`
	Path     string          `json:"path"`
	Pagename string          `json:"pagename"`
}
