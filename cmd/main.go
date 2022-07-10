package main

import (
	"container/list"
	"net/url"
	"os"

	"github.com/ungerik/go-rss"
)

type MyFeedItem struct {
	Title    string
	Link     string
	PubDate  rss.Date
	Category []string
}

func log_err(err error) {
	os.Stderr.WriteString("ERR: " + err.Error())
}

func main() {
	test_feed_rss := "https://news.ycombinator.com/rss"
	feeds := make(map[string]interface{})

	u, err := url.Parse(test_feed_rss)
	if err != nil {
		panic(err)
	}

	hackernews := false
	if u.Host == "ycombinator.com" {
		hackernews = true
	}

	resp, err := rss.Read(test_feed_rss, hackernews)
	if err != nil {
		log_err(err)
	}

	channel, err := rss.Regular(resp)
	if err != nil {
		log_err(err)
	}

	my_items := list.New() //MyFeedItem

	for _, item := range channel.Item {
		my_items.PushBack(MyFeedItem{
			Title:    item.Title,
			PubDate:  item.PubDate,
			Link:     item.Link,
			Category: item.Category,
		})
	}

	feeds[channel.Title] = my_items
}
