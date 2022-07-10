package main

import (
	"container/list"
	"log"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/Aggregator", handler)
	log.Printf("Started listening on %s. URL: https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
