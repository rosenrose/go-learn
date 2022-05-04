package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	title    string
	company  string
	link     string
	location string
	summary  string
	salary   string
}

func Scrap(query string) {
	baseUrl := fmt.Sprintf("https://www.indeed.com/jobs?q=%v", query)
	totalPages := getPages(baseUrl)
	var totalJobs []extractedJob
	channel := make(chan []extractedJob)

	for i := 0; i < totalPages; i++ {
		go getPage(baseUrl, i, channel)
	}
	for i := 0; i < totalPages; i++ {
		totalJobs = append(totalJobs, <-channel...)
	}

	fmt.Printf("Extracted %v jobs", len(totalJobs))
	writeJobs(totalJobs)
}

func getPages(baseUrl string) int {
	var pages int

	res, err := http.Get(baseUrl)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("div.pagination ul.pagination-list").Each(func(i int, ul *goquery.Selection) {
		pages = ul.Find("a").Length()
	})

	return pages
}

func getPage(baseUrl string, page int, mainChannel chan<- []extractedJob) {
	pageUrl := fmt.Sprintf("%v&start=%v", baseUrl, page*10)
	fmt.Println("Requesting ", pageUrl)

	req, err := http.NewRequest("GET", pageUrl, nil)
	checkErr(err)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	var jobs []extractedJob
	jobCards := doc.Find("div#mosaic-provider-jobcards > a")
	channel := make(chan extractedJob)

	jobCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, channel)
	})
	for i := 0; i < jobCards.Length(); i++ {
		jobs = append(jobs, <-channel)
	}

	mainChannel <- jobs
}

func extractJob(card *goquery.Selection, channel chan<- extractedJob) {
	title, _ := card.Find("h2.jobTitle span[title]").Attr("title")
	company := CleanString(card.Find("span.companyName").Text())
	id, _ := card.Attr("data-jk")
	link := fmt.Sprintf("https://www.indeed.com/viewjob?jk=%v&vjs=3", id)

	location := CleanString(card.Find("div.companyLocation").Text())
	regex := regexp.MustCompile(`\+(1 location|1?\d+ locations)`)
	location = regex.ReplaceAllString(location, "")

	summary := CleanString(card.Find("div.job-snippet").Text())

	salary := ""
	salaryDiv := card.Find(".salary-snippet-container, .estimated-salary-container")
	if salaryDiv != nil {
		salary = CleanString(salaryDiv.Text())
	}
	// fmt.Println(title, company, id, location, summary, salary)

	channel <- extractedJob{title: title, company: company, link: link, location: location, summary: summary, salary: salary}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	columnTitle := []string{"Title", "Company", "Link", "Location", "Summary", "Salary"}
	err = writer.Write(columnTitle)
	checkErr(err)

	for _, job := range jobs {
		record := []string{job.title, job.company, job.link, job.location, job.summary, job.salary}
		err := writer.Write(record)
		checkErr(err)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %v %v", res.StatusCode, res.Status)
	}
}
