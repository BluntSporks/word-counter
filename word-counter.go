// word-counter counts the words in a text file.
package main

import (
	"bufio"
	"fmt"
	"github.com/BluntSporks/mapsort"
	"log"
	"os"
	"regexp"
	"strings"
)

var wordRegExp = regexp.MustCompile(`\pL+('\pL+)*`)
var wordCnts = make(map[string]int)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing filename argument")
	}
	file := os.Args[1]
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
	// Print the substring counts.
	pairs := mapsort.ByVal(wordCnts, false)
	for _, pair := range pairs {
		fmt.Printf("%s,%d\n", pair.Key, pair.Val)
	}
}
