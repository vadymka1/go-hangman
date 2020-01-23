package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetWord(variant int) string {

	if variant == 1 {
		return getWordFromApi()
	}

	if variant == 2 {
		return scrapeFromUrl()
	}

	return ""
}

func getWordFromApi() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("api_path")
	key := os.Getenv("rapid_api")

	fmt.Print("This from api")
	fmt.Printf("url is %s and key is %s", url, key)

	var jsonMap map[string]interface{}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "wordsapiv1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &jsonMap); err != nil {
		panic(err)
	}

	fmt.Println(jsonMap["word"].(string))

	return jsonMap["word"].(string)
}

func scrapeFromUrl() string{

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("url_path")

	fmt.Print("This from crawler")
	fmt.Printf("url is %s \n", url)

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