package controller

import (
	"encoding/xml"
	"server/service"
	"server/util/code"
	"server/util/rest"

	"github.com/gin-gonic/gin"
)

type RSSXml struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel string   `xml:"channel"`
}

func GetRss(c *gin.Context) {
	doc, err := service.GetRss(c)
	if err != nil {
		rest.Res(c, code.NotFound, nil)
	}
	v := RSSXml{Version: "2.0", Channel: doc}
	c.XML(200, &v)
}
