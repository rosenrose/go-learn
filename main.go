package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rosenrose/go-learn/accounts"
	"github.com/rosenrose/go-learn/dict"
)

var errRequestFailed = errors.New("request failed")

func main() {
	// banking()
	// dictionary()
	urlChecker()
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
	dictionary := dict.Dictionary {"hello": "안녕"}

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

func urlChecker() {
	urls := []string {
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://nomadcoders.co/",
	}
	// results := map[string]error {}
	results := make(map[string]string)
	
	for _, url := range urls {
		result := "Fail"
		err := hitUrl(url)

		if err == nil {
			result = "Success"
		}
	
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking:", url)
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return errRequestFailed
	}

	return nil
}