/*
  feed
    get feed from FlatTrack's blog
*/

package feed

import (
	"log"

	"github.com/mmcdole/gofeed"
)

// GetLatestRSSPost ...
// fetch the feed from FlatTrack's blog
func GetLatestRSSPost () (*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://blog.flattrack.io/rss/")
	log.Println(feed.Items)
	latestPost := feed.Items[0]
	return latestPost, err
}

