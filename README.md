# GoTick

A dead simple RSS news ticker for the console. Watch the news go by, and press Enter when something catches your eye.

<img alt="gotick" src="https://i.imgur.com/2LfYwI7.png" />

* `Z` to cycle between several text colors
* `<Up>`, `<Down>`, `j`, `k` to scroll
* `q` or `Ctrl+C` to quit
* `Enter` to open the selected item

Feeds and some options are configured in the `config.json` file

```json
{
	"TextColor": 2,
	"RefreshRate": 60,
	"Feeds": [
		{"title":"DistroWatch", "url":"https://distrowatch.com/news/dw.xml"}
	]
}
```

* `TextColor` (an integer between 0 and 6) defines a text [color option](https://github.com/gizak/termui/blob/f976fe697aa09b747f16fa6e08c36dde5fb16f27/theme.go#L7)
* `RefreshRate` defines how often the feeds will be updated, in seconds
* `Feeds` is an array of JSON objects, each containing a `title` and `url` parameters, self-explanatory