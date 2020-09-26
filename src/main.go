package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher is coroutine safe
	searcher = riot.New("zh")
)

func main() {
	//实例化echo对象。
	e := echo.New()
	initRouter(e)

	data := types.DocData{Content: `I wonder how, I wonder why
		, I wonder where they are`}
	data1 := types.DocData{Content: "所以, 你好, 再见"}
	data2 := types.DocData{Content: "没有理由"}

	// Add the document to the index, docId starts at 1
	searcher.Index("1", data)
	searcher.Index("2", data1)
	searcher.Index("3", data2)

	// Wait for the index to refresh
	searcher.Flush()
	// engine.FlushIndex()

	// The search output format is found in the types.SearchResp structure
	log.Print(searcher.Search(types.SearchReq{Text: "google testing"}))

	req := types.SearchReq{Text: "你好"}
	search := searcher.Search(req)
	log.Print("search...", search)

	e.Start(":8080")
}
