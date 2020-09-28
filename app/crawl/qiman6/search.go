package qiman6

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gocolly/colly"
)

// SearchResults 搜索结果
type SearchResults struct {
	Data []SearchResult `json:"data"`
}

// SearchResult 单条搜索结果
type SearchResult struct {
	ID            string `json:"id"`
	URL           string `json:"url"`
	Cover         string `json:"cover"`
	LatestChapter string `json:"latestChapter"`
	Title         string `json:"title"`
	Author        string `json:"author"`
}

// Search 搜索
func (q *New) Search(query string) (SearchResults, error) {
	var result SearchResults
	var err error
	var searchList []SearchResult
	// Instantiate default collector
	myColly := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		colly.AllowedDomains(Domain),
	)

	myColly.OnHTML(".mainForm > .updateList > .bookList_3", func(e *colly.HTMLElement) {

		e.ForEach(".item.ib", func(i int, h *colly.HTMLElement) {
			item := SearchResult{
				ID:            strings.Split(h.ChildAttr("a", "href"), "/")[1],
				URL:           Domain + h.ChildAttr("a", "href"),
				Cover:         h.ChildAttr(".book > a > .cover", "src"),
				LatestChapter: h.ChildText(".book > a > .msg.op"),
				Title:         h.ChildText(".title"),
				Author:        h.ChildText(".tip"),
			}

			searchList = append(searchList, item)
		})

		result.Data = searchList
	})

	// Set error handler
	myColly.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Before making a request print "Visiting ..."
	myColly.OnRequest(func(r *colly.Request) {
		fmt.Println("OnRequest", r.URL.String())
	})

	myColly.OnResponse(func(resp *colly.Response) {
		a := reflect.TypeOf(resp.Body)
		fmt.Println("OnResponse:", a)
	})

	visitURL := fmt.Sprintf("http://www.qiman6.com/search.php?keyword=%s", query)
	myColly.Visit(visitURL)
	return result, err
}
