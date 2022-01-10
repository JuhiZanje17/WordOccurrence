package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func noOfOccurance(fileName string) wordArr {
	mydir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(mydir + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	countMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		obj := regexp.MustCompile("[^a-z|^A-Z|0-9]")
		words := obj.Split(strings.ToLower(scanner.Text()), -1) //to make sure words 'india' and 'india,' are considered same
		for _, word := range words {
			if word != "" {
				countMap[word]++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
	fmt.Println(res)
	return res
}

func main() {

	fileName := "temp1.txt"
	res := noOfOccurance(fileName)
	for i := 0; i < noOfWords; i++ {
		fmt.Println("Word=", res[i].Word, " Count=", res[i].Count)
	}
}
