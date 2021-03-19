package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"servercheck/shared"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var pagedata = shared.PageData{
	PageInfo: make(map[string]shared.Info),
}

func Server() {
	r = gin.Default()
	r.POST("/upload", Upload)
	r.GET("/", Home)
	r.GET("/allservers", Home)
	r.GET("/outdated/:hostname", Home)
	r.Run(":8080")
}

func Upload(c *gin.Context) {
	var infoCached shared.Info
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &infoCached)
	if err != nil {
		log.Fatal(err)
	}
	infoCached.Newversions = Aptparse(infoCached)
	infoCached.Outdated = Outdated(infoCached)
	pagedata.PageInfo[infoCached.Hostname] = infoCached
}

func Home(c *gin.Context) {
	r.LoadHTMLGlob("*.html")
	pagedata.Hostname = c.Param("hostname")
	switch c.Request.URL.Path {
	case "/":
		pagedata.Path = "/"
		pagedata.Pagename = "Home"
	case "/allservers":
		pagedata.Path = "allservers"
		pagedata.Pagename = "All Servers"
	case "/outdated/" + pagedata.Hostname:
		pagedata.Path = "outdated"
		pagedata.Pagename = pagedata.Hostname
		pi := pagedata.PageInfo[pagedata.Hostname]
		pi.Outdatedrepos = Outdatedrepos(pagedata.Hostname)
		pagedata.PageInfo[pagedata.Hostname] = pi
	default:
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.HTML(http.StatusOK, "index", pagedata)
}
