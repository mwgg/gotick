package feeds

import (
	"time"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mwgg/gotick/helpers"
)

func ProcessItems(ch_items chan []string, l *widgets.List, FeedItemUrls map[int]string) {
	var displayedItems []string

	for {
		item := <-ch_items
		if !helpers.Contains(displayedItems, item[1]) {
			l.Rows = append(l.Rows, item[0])
			FeedItemUrls[len(l.Rows)-1] = item[2]
			l.ScrollDown()
			termui.Render(l)
			displayedItems = append(displayedItems, item[1])
		}
		time.Sleep(time.Millisecond * 10)
	}
}
