package feeds

import (
	"html"
	"sort"

	"github.com/mmcdole/gofeed"
	"github.com/mwgg/gotick/helpers"
)

func getFeed(url string, title string, ch_items chan []string) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		//
	} else {
		sort.Sort(feed)
		for _, item := range feed.Items {
			ch_items <- []string{title + ": " + html.UnescapeString(item.Title), helpers.HashMD5(item.Title), item.Link}
		}
	}
}

func UpdateAllFeeds(feeds []map[string]string, ch_items chan []string) {
	for _, feedData := range feeds {
		go func(url string, title string, ch_items chan []string) {
			getFeed(url, title, ch_items)
		}(feedData["url"], feedData["title"], ch_items)
	}
}
