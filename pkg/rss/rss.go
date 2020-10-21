// Copyright @tsoonjin
// All Rights Reserved
//
// Process RSS Feed
package rss

import (
    "fmt"
    "database/sql"
    "github.com/mmcdole/gofeed"
    _ "github.com/mattn/go-sqlite3"
    "strings"
)

const USER = "Jin"
const DB_PATH = "../../db/rss.db"

// List all urls given a RSS Feed
func List(feedUrl string) string {
    database, _ := sql.Open("sqlite3", "./db/rss.db")
    rows, _ := database.Query(fmt.Sprintf("SELECT feeds FROM Users WHERE Name = \"%s\" ", USER))
    var feeds string
    for rows.Next() {
        rows.Scan(&feeds)
        fmt.Println(feeds)
    }
    fp := gofeed.NewParser()
    feed, _ := fp.ParseURL(strings.Split(feeds, ",")[0])
    fmt.Println("Title: " + feed.Title)
    return feeds
}
