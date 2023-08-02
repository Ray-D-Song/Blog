package service

import (
	"context"
	"github.com/beevik/etree"
	"strconv"
)

var m = map[string]string{
	"title":       "Ray-D-Song's Blog",
	"link":        "http://ray-d-song.com",
	"description": "code, meme and tech",
	"language":    "en",
	"copyright":   "Copyright 2023",
}

func GetRss(ctx context.Context) (string, error) {
	list, err := GetBlogList(ctx)
	if err != nil {
		return "", err
	}
	doc := etree.NewDocument()

	m["lastBuildDate"] = list[0].Time
	for key, value := range m {
		doc.CreateElement(key).CreateText(value)
	}
	for _, value := range list {
		item := doc.CreateElement("item")
		item.CreateElement("title").CreateText(value.Title)
		item.CreateElement("description").CreateText(value.Subtitle)
		item.CreateElement("link").CreateText("blog/" + strconv.Itoa(value.ID))
		item.CreateElement("guid").CreateText("blog/" + strconv.Itoa(value.ID))
		item.CreateElement("pubDate").CreateText(value.Time)
	}
	doc.Indent(2)
	value, err := doc.WriteToString()
	if err != nil {
		return "", err
	}
	return value, nil
}
