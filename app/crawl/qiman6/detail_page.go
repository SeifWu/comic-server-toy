package qiman6

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/robertkrimen/otto"
)

// Picture 具体内容
type Picture struct {
	URL     string `json:"url"`
	Postion int    `json:"postion"`
}

// ComicDetailPage 漫画详情页
func (q *New) ComicDetailPage(url string) ([]Picture, error) {
	var pictures []Picture
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// count links
	c.OnHTML("body script", func(e *colly.HTMLElement) {
		boolean, _ := regexp.MatchString("eval(.*)", e.Text)
		if boolean {
			javaScript := e.Text
			javaScript += " function getImgList() {return newImgs;} getImgList();"

			vm := otto.New()
			value, _ := vm.Run(javaScript)
			imagesString, _ := value.ToString()
			result := strings.Split(imagesString, ",")

			for index, value := range result {
				picture := Picture{
					Postion: index + 1,
					URL:     value,
				}

				pictures = append(pictures, picture)
			}
		}
	})

	c.Visit(fmt.Sprintf("http://%s", url))

	return pictures, nil
}
