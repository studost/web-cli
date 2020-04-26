package main

import (
    "fmt"
    "github.com/mmcdole/gofeed"
)

func main() {
    fmt.Printf("Eine Zeile Text")
    fp := gofeed.NewParser()
    feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
    for i := 0; i < 5; i++ {
            fmt.Println(feedItems[i].Title)
    }
}


