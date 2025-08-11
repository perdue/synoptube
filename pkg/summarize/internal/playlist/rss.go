package playlist

import (
	"encoding/xml"
	"time"
)

// RSSFeed represents the structure of a YouTube RSS feed for a playlist.
type RSSFeed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

// Entry represents a single video entry in the RSS feed.
type Entry struct {
	XMLName xml.Name  `xml:"entry"`
	VideoID string    `xml:"videoId"` // YouTube video ID
	Title   string    `xml:"title"`
	Link    Link      `xml:"link"`
	Updated time.Time `xml:"updated"`
}

// Link represents the link element in an RSS entry.
type Link struct {
	Href string `xml:"href,attr"`
}
