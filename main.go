package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mottet-dev/medium-go-colly-basics/utils"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var productName, imageUrl, description, price, totalReviews,stars string

			productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				return
			}

			imageUrl = e.ChildAttr("img.s-image","src")

			totalReviews = e.ChildText("span.a-size-base")

			description = e.ChildText("span.a-size-base.a-color-secondary")

			description = strings.ReplaceAll(description,totalReviews,"")

			stars = e.ChildText("span.a-icon-alt")
			utils.FormatStars(&stars)

			price = e.ChildText("span.a-price > span.a-offscreen")
			utils.FormatPrice(&price)

			fmt.Printf("Product Name: %s \nImage URL :%s \nDescription:%s \nTotal Reviews:%s \nStars: %s \nPrice: %s \n", productName, imageUrl,description, totalReviews,stars, price)
		})
	})

	c.Visit("https://www.amazon.in/s?k=ps2&ref=nb_sb_noss_2")
}
