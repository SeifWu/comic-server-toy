package qiman6

import (
	"fmt"
	"reflect"

	"github.com/gocolly/colly"
)

type SearchResults struct {
	Data []searchResult
}

type searchResult struct {
	URL           string
	Cover         string
	LatestChapter string
	Title         string
	Tip           string
}

// Search 搜索
func (q *Qiman6) Search(query string) (SearchResults, error) {
	var result SearchResults
	var err error
	var searchList []searchResult
	domain := "www.qiman6.com"

	// Instantiate default collector
	myColly := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		colly.AllowedDomains(domain),
	)

	myColly.OnHTML(".mainForm > .updateList > .bookList_3", func(e *colly.HTMLElement) {

		e.ForEach(".item.ib", func(i int, h *colly.HTMLElement) {
			item := searchResult{
				URL:           domain + h.ChildAttr("a", "href"),
				Cover:         h.ChildAttr(".book > a > .cover", "src"),
				LatestChapter: h.ChildText(".book > a > .msg.op"),
				Title:         h.ChildText(".title"),
				Tip:           h.ChildText(".tip"),
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

	myColly.Visit("http://www.qiman6.com/search.php?keyword=" + query)
	return result, err
}
