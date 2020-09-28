package qiman6

// New 爬取 http://www.qiman6.com 内容
type New struct{}

// Domain 目标站点
var Domain = "www.qiman6.com"

// OtherChapter 请求来的章节
type OtherChapter struct {
	ChapterID   string `json:"chapterid"`
	ChapterName string `json:"chaptername"`
}

// Chapter 章节
type Chapter struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

// Comic 漫画结构体
type Comic struct {
	Title         string    `json:"title"`
	Introduction  string    `json:"introduction"`
	Cover         string    `json:"cover"`
	Author        string    `json:"author"`
	LatestChapter string    `json:"latestChapter"`
	ChapterList   []Chapter `json:"chapterList"`
}
