package services

import (
	"fmt"
	"github.com/icrowley/fake"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Chars struct{
	RightChar, ErrorChar string
}

type ViewData struct{
	Word, Title string
	CharCounter int
}

type AllData struct {
	View  	 ViewData
	Chars 	 Chars
	Errors   int
	WordMask string
	Over     bool
}

var over bool

var title = "Hangman"

func ShowForm (w http.ResponseWriter, r *http.Request) {

	over = false

	Word := fake.Word()

	CharNumbers := len(Word)

	WordMask := strings.Repeat("_", CharNumbers)

	data := AllData {
		View : ViewData {
			Word,
			title,
			CharNumbers,
		},
		Chars:Chars{
			"",
			"",
		},
		Errors:0,
		WordMask:WordMask,
		Over:over,
	}

	fmt.Println(data)

	t, err := template.ParseFiles("templates/form.html")

	if err != nil {
		panic(err)
	}

	err = t.Execute(w, data)

	fmt.Println(err)

}

func CheckData (w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
	}

	word := r.FormValue("word")

	workMask := r.FormValue("wordMask")

	errors, err := strconv.Atoi(r.FormValue("errors"))

	if err != nil {
		errors = 0
	}

	if errors > 5 {
		over = true
	}

	errorChar := r.FormValue("errorChar")
	respChar := r.FormValue("char")
	rightChar := r.FormValue("rightChar")

	fmt.Println(workMask, word, errors, errorChar, respChar, rightChar)



	for k, v := range workMask {
		fmt.Println("key : %d, value : %s\n", k, string(v))
	}

	for k, char := range word {
		if string(char) == respChar {
			ch := string(char)
			workMask[k] = string(char)

			rightChar = rightChar + string(char)
		} else {
			errors++
		}
	}

	for key, value := range r.Form{
		fmt.Printf("%s = %s\n", key, value)
	}

	data := AllData {
		View : ViewData {
			word,
			title,
			len(word),
		},
		Chars:Chars{
			rightChar,
			errorChar,
		},
		Errors:0,
		WordMask:workMask,
		Over:over,
	}
}
