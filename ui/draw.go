package ui

import (
	"fmt"
	"os"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/pkg/browser"
)

func DrawUI(feedCount int, l *widgets.List, FeedItemUrls map[int]string, TextColor int) {

	feedCountWord := " feed"
	if feedCount > 1 {
		feedCountWord = " feeds"
	}
	l.Title = "Monitoring " + fmt.Sprint(feedCount) + feedCountWord
	l.Rows = []string{}
	l.TextStyle = termui.NewStyle(termui.StandardColors[TextColor])
	l.WrapText = false
	l.SetRect(0, 0, 50, 16)

	termui.Render(l)

	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			os.Exit(0)
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "z":
			l.TextStyle = NextTheme(l.TextStyle)
		case "<Enter>":
			browser.OpenURL(FeedItemUrls[l.SelectedRow])
		case "<Resize>":
			payload := e.Payload.(termui.Resize)
			l.SetRect(0, 0, payload.Width, payload.Height)
		}
		termui.Render(l)
	}
}

func NextTheme(currentStyle termui.Style) termui.Style {
	var index int
	for i, item := range termui.StandardStyles {
		if item == currentStyle {
			index = i
			break
		}
	}
	if len(termui.StandardStyles) > index+1 {
		return termui.StandardStyles[index+1]
	} else {
		return termui.StandardStyles[0]
	}
}
