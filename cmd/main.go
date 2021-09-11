package main

import (
	"database/sql"
	"fmt"
	"newsfeed/platform/newsfeed"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./newsfeed.db")
	feed := newsfeed.NewFeed(db)

	// feed.Add(newsfeed.Item{
	// 	Content: "Hello!",
	// })

	items := feed.Get()

	fmt.Println(items)

}
