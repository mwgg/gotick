package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mwgg/gotick/feeds"
	"github.com/mwgg/gotick/ui"
)

type Config struct {
	TextColor   int
	RefreshRate int
	Feeds       []map[string]string
}

func loadConfig() Config {
	var conf Config
	confContent, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Error reading config.json: %v", err)
	}
	err = json.Unmarshal(confContent, &conf)
	if err != nil {
		log.Fatalf("Error processing config.json: %v", err)
	}
	return conf
}

func main() {
	var conf Config = loadConfig()
	FeedItemUrls := make(map[int]string)

	// UI things
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()
	l := widgets.NewList()
	if conf.TextColor > len(termui.StandardStyles)-1 {
		conf.TextColor = len(termui.StandardStyles) - 1
	}
	go ui.DrawUI(len(conf.Feeds), l, FeedItemUrls, conf.TextColor)

	// Feed things
	ch_items := make(chan []string)
	go feeds.ProcessItems(ch_items, l, FeedItemUrls)
	for {
		feeds.UpdateAllFeeds(conf.Feeds, ch_items)
		time.Sleep(time.Duration(conf.RefreshRate * 1000000000))
	}
}
