package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/rosenrose/go-learn/accounts"
	"github.com/rosenrose/go-learn/dict"
	"github.com/rosenrose/go-learn/scrapper"
)

var errRequestFailed = errors.New("request failed")

const CSV_NAME string = "jobs.csv"

func main() {
	// banking()
	// dictionary()
	// urlChecker()

	server := echo.New()
	server.GET("/", func(context echo.Context) error {
		return context.File("home.html")
	})
	server.POST("/scrap", func(context echo.Context) error {
		defer os.Remove(CSV_NAME)

		query := context.FormValue("query")
		query = strings.ToLower(scrapper.CleanString(query))

		scrapper.Scrap(query)
		return context.Attachment(CSV_NAME, CSV_NAME)
	})
	server.Logger.Fatal(server.Start(":8000"))
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
