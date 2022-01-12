package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

type wordStruct struct {
	Word  string
	Count int
}

type wordArr []wordStruct

var noOfWords int = 10

func fileUpload(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("form.html")
	switch r.Method {
	case "GET":
		tmpl.Execute(w, nil)
	case "POST":
		startCount(w, r)
	}
}

func startCount(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}
	Res := make(wordArr, noOfWords)
	Res = noOfOccurance(fileBytes)
	// Resp := Response{Res}
	tmpl, _ := template.ParseFiles("form2.html")
	tmpl.Execute(w, struct{ Response wordArr }{Res})
}

func noOfOccurance(data []byte) wordArr {

	countMap := make(map[string]int)
	obj := regexp.MustCompile("[^a-z|^A-Z|^0-9]")
	words := obj.Split(strings.ToLower(string(data)), -1) //to make sure words 'india' and 'india,' are considered same
	for _, word := range words {
		if word != "" {
			countMap[word]++
		}
	}

	//code to sort the map
	res := make(wordArr, len(countMap))

	i := 0
	for k, v := range countMap {
		res[i] = wordStruct{k, v}
		i++
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Count > res[j].Count
	})
	if len(res) < 10 {
		noOfWords = len(res)
	}
	res = res[0:noOfWords]
	return res
}

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", fileUpload)
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
