package services

import (
	"fmt"
	"github.com/icrowley/fake"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type ViewData struct{
	Word, Title string
	CharCounter int
}

type AllData struct {
	View  	 ViewData
	Errors   int64
	ErrorChars string
	WordMask string
	Win bool
}

var (
	title = "Hangman"
	win = false
	errorChars = ""
)

func ShowForm (w http.ResponseWriter, r *http.Request) {


	Word := fake.Word()

	CharNumbers := len(Word)

	WordMask := strings.Repeat("_", CharNumbers)

	data := AllData {
		View : ViewData {
			Word,
			title,
			CharNumbers,
		},
		ErrorChars:errorChars,
		Errors:0,
		WordMask:WordMask,
		Win:win,
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
	errors, _ := strconv.ParseInt(r.FormValue("errors"), 10, 0)
	errorChar := r.FormValue("errorChar")
	respChar := r.FormValue("char")
	rightChar := r.FormValue("rightChar")

	if respChar == "" {
		panic("char mustnt be empty")
	}

	errorExist := true

	for k, char := range word {

		if string(char) == respChar {

			workMask = replaceAtindex(workMask, char, k)

			rightChar = rightChar + string(char)

			errorExist = false
		}
	}

	checkWorkMask := strings.Index(workMask, "_")

	if checkWorkMask == -1 {
		win = true
	}

	if errorExist {
		errorChar = errorChar + respChar
		errors++
	}

	data := AllData {
		View : ViewData {
			word,
			title,
			len(word),
		},
		ErrorChars:errorChar,
		Errors:errors,
		WordMask:workMask,
		Win:win,
	}

	t, err := template.ParseFiles("templates/form.html")

	if err != nil {
		panic(err)
	}

	err = t.Execute(w, data)

}

func replaceAtindex(in string, r rune, i int) string {

	out := []rune(in)

	if out[i] == r {
		return string(out)
	}

	out[i] = r
	return string(out)
}
