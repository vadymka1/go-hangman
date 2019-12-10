package services

import (
	"fmt"
	"github.com/icrowley/fake"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var errCounter, rightCounter int

type ErrorChar struct{
	wChar, rChar string
}

type ViewData struct{
	Word, Title string
	CharCounter int
}

type AllData struct {
	view  	ViewData
	chars 	ErrorChar
	errors  int
	word    string
}

func ShowForm (w http.ResponseWriter, r *http.Request) {

	word := fake.Word()

	charNumbers := len(word)

	chars := strings.Repeat("_", charNumbers)

	data := &AllData {
		view : ViewData {
			Word 		: word,
			Title 		: "Hangman",
			CharCounter : charNumbers,
		},
		chars:ErrorChar{
			wChar: "",
			rChar: "",
		},
		errors:0,
		word:chars,
	}

	t, _ := template.ParseFiles("templates/form.html")

	t.Execute(w, data)

}

func CheckData (w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

}
