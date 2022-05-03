package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rosenrose/go-learn/accounts"
	"github.com/rosenrose/go-learn/dict"
)

var errRequestFailed = errors.New("request failed")

func main() {
	// banking()
	// dictionary()
	// urlChecker()
	jobScrapper()
}

func banking() {
	account := accounts.NewAccount("Hans")
	fmt.Println(*account)

	account.Deposit(300)
	fmt.Println(account.Balance())

	err := account.Withdraw(500)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	err = account.Withdraw(200)

	account.ChangeOwner("Alice")
	fmt.Println(account.Owner(), account.Balance(), err)
}

func dictionary() {
	dictionary := dict.Dictionary{"hello": "안녕"}

	definition, err := dictionary.Search("bye")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	definition, err = dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("hello", "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}
	err = dictionary.Add("good", "bad")
	if err != nil {
		fmt.Println(err)
	} else {
		definition, _ = dictionary.Search("good")
		fmt.Println(definition)
	}

	err = dictionary.Update("what", "the")
	if err != nil {
		fmt.Println(err)
	} else {
		definition, _ = dictionary.Search("what")
		fmt.Println(definition)
	}
	err = dictionary.Update("hello", "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		definition, _ = dictionary.Search("hello")
		fmt.Println(definition)
	}

	fmt.Println(dictionary)
	dictionary.Delete("ㅋㅋ")
	fmt.Println(dictionary)
	dictionary.Delete("good")
	fmt.Println(dictionary)
}

type hitResult struct {
	err    error
	status string
}

func urlChecker() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://nomadcoders.co/",
		"https://go.dev/",
		"https://code.visualstudio.com/",
	}
	results := map[string]hitResult{}
	channel := make(chan hitResult)

	for _, url := range urls {
		go hitUrl(url, channel)
	}

	for _, url := range urls {
		results[url] = <-channel
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string, channel chan<- hitResult) { // send only / <-chan receive only
	fmt.Println("Checking:", url)
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		channel <- hitResult{err: errRequestFailed, status: "Fail"}
	} else {
		channel <- hitResult{err: nil, status: "Success"}
	}
}

var baseUrl = "https://www.indeed.com/jobs?q=golang"

type extractedJob struct {
	title    string
	company  string
	id       string
	location string
	summary  string
	salary   string
}

func jobScrapper() {
	// totalPages := getPages()
	var totalJobs []extractedJob

	for i := 0; i < 2; i++ {
		extractedJobs := getPage(i)
		totalJobs = append(totalJobs, extractedJobs...)
	}

	fmt.Println(len(totalJobs), totalJobs)
}

func getPages() int {
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

func getPage(page int) []extractedJob {
	pageUrl := fmt.Sprintf("%v&start=%v", baseUrl, page*10)
	fmt.Println("Requesting ", pageUrl)
	var jobs []extractedJob

	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobCards := doc.Find("div#mosaic-provider-jobcards > a")
	jobCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	title, _ := card.Find("h2.jobTitle span[title]").Attr("title")
	company := cleanString(card.Find("span.companyName").Text())
	id, _ := card.Attr("data-jk")
	location := cleanString(card.Find("div.companyLocation").Text())
	summary := cleanString(card.Find("div.job-snippet").Text())

	salary := ""
	salaryDiv := card.Find(".salary-snippet-container, .estimated-salary-container")
	if salaryDiv != nil {
		salary = cleanString(salaryDiv.Text())
	}
	// fmt.Println(title, company, id, location, summary, salary)

	return extractedJob{title: title, company: company, id: id, location: location, summary: summary, salary: salary}
}

func cleanString(str string) string {
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
