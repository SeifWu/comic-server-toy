package qiman6

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func (q *Qiman6) ComicPageQiman6(url string) {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	// c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 5})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(".mainForm", func(e *colly.HTMLElement) {
		cover := e.ChildAttr(".comicInfo > .ib.cover > .img > img", "src")
		fmt.Println("封面：", cover)
		lastChapter := e.ChildText(".comicInfo > .ib.cover > .img > a")
		fmt.Println("最新话：", lastChapter)
		title := e.ChildText(".comicInfo > .ib.info > .name_mh")
		fmt.Println("漫画名：", title)
		author := e.ChildText(".comicInfo > .ib.info > p:not(.gray) > .ib.l")
		fmt.Println("作者: ", author)
		introduce := e.ChildText(".comicInfo > .ib.info > .content")
		fmt.Println("简介: ", introduce)
	})

	c.Visit(url)
}
