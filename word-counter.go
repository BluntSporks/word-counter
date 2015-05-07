// word-counter counts the words in a text file.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/BluntSporks/mapsort"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var wordRegExp = regexp.MustCompile(`\pL+('\pL+)*`)
var wordCnts = make(map[string]int)

func main() {
	lookup := flag.String("lookup", "", "Existing word count file to use for lookup")
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Missing filename argument")
	}
	file := flag.Arg(0)

	// Choose program function.
	if *lookup != "" {
		lookUpCnts(file, *lookup)
	} else {
		cntWords(file)
	}

	// Print the word counts.
	pairs := mapsort.ByVal(wordCnts, false)
	for _, pair := range pairs {
		fmt.Printf("%s,%d\n", pair.Key, pair.Val)
	}
}

// cntWords counts the words in a file.
func cntWords(file string) {
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()
	scanner := bufio.NewScanner(hdl)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := strings.ToLower(scanner.Text())
		words := wordRegExp.FindAllString(line, -1)
		for _, word := range words {
			wordCnts[word]++
		}
	}
}

// lookUpCnts looks up the word counts for a file in an existing word count file.
func lookUpCnts(file string, lookup string) {
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()
	scanner := bufio.NewScanner(hdl)
	wordsToLookUp := make(map[string]bool)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := strings.ToLower(scanner.Text())
		words := wordRegExp.FindAllString(line, -1)
		for _, word := range words {
			wordsToLookUp[word] = true
		}
	}
	lookupHdl, err := os.Open(lookup)
	if err != nil {
		log.Fatal(err)
	}
	defer lookupHdl.Close()
	scanner = bufio.NewScanner(lookupHdl)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := strings.ToLower(scanner.Text())
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			continue
		}
		word := fields[0]
		if wordsToLookUp[word] {
			cnt, err := strconv.Atoi(fields[1])
			if err == nil {
				wordCnts[word] = cnt
			}
		}
	}
}
