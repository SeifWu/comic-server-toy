package qiman6

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// ComicPage 漫画详情页
func (q *New) ComicPage(id string) (Comic, error) {

	var moreChapter = map[string]string{"id": id, "id2": "1"}
	var comic Comic
	var chapterList []Chapter
	var otherChapters []OtherChapter
	c := colly.NewCollector()
	visited := false

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// count links
	c.OnHTML(".mainForm", func(e *colly.HTMLElement) {
		title := e.ChildText(".comicInfo .ib.info h1.name_mh")
		log.Println("获取名称: ", title)

		introduction := e.ChildText(".comicInfo .ib.info p.content")
		log.Println("获取简介: ", introduction)

		cover := e.ChildAttr(".comicInfo .ib.cover .img img", "src")
		log.Println("获取封面: ", cover)

		latestChapter := e.ChildText(".comicInfo .ib.cover .img a.op.ib")
		log.Println("获取最新章节: ", latestChapter)

		e.ForEach(".chapterList#chapterList .list#chapter-list1 a.ib", func(i int, h *colly.HTMLElement) {
			chapter := Chapter{
				Title: h.Text,
				URL:   Domain + h.Attr("href"),
			}
			chapterList = append(chapterList, chapter)
		})

		comic.Title = title
		comic.Introduction = introduction
		comic.Cover = cover
		comic.LatestChapter = latestChapter
		comic.ChapterList = chapterList
	})

	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &otherChapters)
		if err != nil {
			log.Println("解析JSON错误: ", err.Error())
		}
		if !visited {
			visited = true
			r.Request.Post(fmt.Sprintf("http://%s/bookchapter/", Domain), moreChapter)
		}
	})

	c.Visit(fmt.Sprintf("http://%s/%s/", Domain, id))
	chapters := handleOtherChapters(otherChapters, id)
	comic.ChapterList = append(comic.ChapterList, chapters...)
	return comic, nil
}

func handleOtherChapters(otherChapters []OtherChapter, id string) []Chapter {
	var chapters []Chapter
	for _, value := range otherChapters {
		url := fmt.Sprintf("%s/%s/", Domain, id)
		chapter := Chapter{
			URL:   url + value.ChapterID + ".html",
			Title: value.ChapterName,
		}
		chapters = append(chapters, chapter)
	}

	return chapters
}
