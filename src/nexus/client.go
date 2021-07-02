package nexus

import (
	"github.com/mmcdole/gofeed"
	"seeder/src/config"
	"strconv"
)

type Client struct {
	baseURL string
	Rule config.NodeRule
}

type Torrent struct {
	GUID  string
	Title string
	URL   string
	Size  string
}

func NewClient(source string, limit int, passkey string,Rule config.NodeRule) Client {
	var baseURL = "https://" + source + "/torrentrss.php?rows=" + strconv.Itoa(limit) + "&cat410=1&cat429=1&cat424=1&cat430=1&cat426=1&cat437=1&cat431=1&cat432=1&cat436=1&cat425=1&cat433=1&cat411=1&cat412=1&cat413=1&cat440=1&linktype=dl&passkey=" + passkey
	return Client{
		baseURL: baseURL,
		Rule:Rule,
	}
}

func (c *Client) Get() ([]Torrent, error) {
	var ts []Torrent
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(c.baseURL)
	if err == nil {
		for _, value := range feed.Items {
			ts = append(ts, Torrent{
				GUID:  value.GUID,
				Title: value.Title,
				URL:   value.Enclosures[0].URL,
				Size:  value.Enclosures[0].Length,
			})
		}
		return ts, nil
	}

	return nil, err
}
