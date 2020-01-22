package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Parser interface {
	GetWord()
}

func GetWord(variant int) string {
	if variant == 1 {
		return getWordFromApi("https://wordsapiv1.p.rapidapi.com/words/hatchback/typeOf")
	}

	if variant == 2 {
		//return scrapeFromUrl(os.Getenv("url_path"))
		return scrapeFromUrl("http://watchout4snakes.com/wo4snakes/Random/RandomWord")
	}

	return ""
}

func getWordFromApi(url string) string {
	fmt.Println(url)
		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-host", "wordsapiv1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "1dab3637c5msh61ac6090e769afcp12dcb8jsnf9245d848de4")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))

		return  string(body)
}

func scrapeFromUrl(url string) string{

	fmt.Println(url)

	data := []byte(`{"LastWord":""}`)
	r := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, r)

	req.Header.Set("X-Custom-Header", "test")

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("something wrond : %s\n", err)
	}

	defer resp.Body.Close()

	return string(body)
}