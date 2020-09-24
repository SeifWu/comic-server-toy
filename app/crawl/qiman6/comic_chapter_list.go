// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package qiman6

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func (q *Qiman6) ComicChapterListQiman6(url string) {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > .footer`),
		// find and click "Expand All" link
		chromedp.Click(`#chapterlistload > .moreChapter`, chromedp.NodeVisible),
		// retrieve the value of the textarea
		chromedp.Value(`#example_After .play .input textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
