package main

import (
	"bufio"
	"fmt"
	"github.com/fitzr/searcher"
	"os"
)

const (
	entries = "./data/10000entries.txt"
	texts   = "./data/texts/"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	searcher := searcher.NewSearcher(entries, texts)

	fmt.Println("\nsearch >")

	for sc.Scan() {
		t := sc.Text()

		if t == "q" {
			break
		}

		results := searcher.Search(t)

		if len(results) == 0 {
			fmt.Println("nothing")
		} else {
			for _, result := range results {
				fmt.Printf("id :%v title:%v url:%v\n", result.Id, result.Title, result.Url)
			}
		}

		fmt.Println("\n\nsearch >")
	}

	fmt.Println("quit")
}
