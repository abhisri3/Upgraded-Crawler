package main

import (
	dbconfig "crawl/dbConfig"
	model "crawl/models"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)


func main() {


	fmt.Println("Started search")
	for {
		searches := dbconfig.GetAllSeraches()
		for _, search := range searches {
			if !search.Completed {
				getInsuranceDetails(search)
			}
		}
	}

}

func getInsuranceDetails(entry model.Search) {

	if !entry.Completed {
		// fmt.Println(url)
		a := colly.NewCollector()

		c := a.Clone()

		cnt := 0
		// containerTags := []string{
		// 	"header", "div.header", "body",
		// }
		// for _, tag := range containerTags {
		a.OnHTML(entry.Container, func(e *colly.HTMLElement) {
			// fmt.Println("------")
			dbconfig.ChangeStatus(entry.Id, true)
			if cnt == 0 {
				e.ForEach("a[href]", func(_ int, kf *colly.HTMLElement) {
					cnt++
					link := kf.Attr("href")
					planUrl := kf.Request.AbsoluteURL(link)
					for _, v := range entry.Patterns {
						if strings.Contains(strings.ToLower(planUrl), v) {
							// fmt.Println(planUrl)

							c.Visit(planUrl)
						}
					}

				})

			}

		})
		// }

		plan_url := ""

		// tempdata := InsuranceData{}
		c.OnHTML("body", func(e *colly.HTMLElement) {
			tempdata := model.InsuranceData{}
			tempdata.Url = plan_url
			for _, title := range entry.Title {
				tempdata.Title = e.ChildText(title)
				if tempdata.Title != "" {
					break
				}
			}
			for _, desc := range entry.Desc {
				tempdata.Desc = e.ChildText(desc)
				if tempdata.Desc != "" {
					break
				}
			}

			for _, benefits := range entry.Benefits {
				e.ForEach(benefits, func(_ int, elem *colly.HTMLElement) {
					tempdata.Benefits = append(tempdata.Benefits, elem.Text)
				})
				if len(tempdata.Benefits) > 0 {
					break
				}

			}

			for _, eligibility := range entry.Eligibility {
				e.ForEach(eligibility, func(_ int, elem *colly.HTMLElement) {
					tempdata.Eligibility = append(tempdata.Eligibility, elem.Text)
				})
				if len(tempdata.Eligibility) > 0 {
					break
				}
			}

			dbconfig.InsertData(tempdata)
		})

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		c.OnRequest(func(r *colly.Request) {
			plan_url = r.URL.String()
			fmt.Println("Visiting : ", r.URL.String())
		})

		a.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting Webiste : ", r.URL.String())
		})
		
		a.Visit(entry.Url)
	}

}
