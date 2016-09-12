package main

import (
	"bufio"
	"fmt"
	"github.com/fitzr/searcher"
	"os"
	"time"
)

const (
	entries = "./data/10000entries.txt"
	texts   = "./data/texts/"
)

func main() {

	fmt.Println("load data ...")

	t1 := time.Now()
	searcher := searcher.NewSearcher(entries, texts)
	t2 := time.Now()

	fmt.Printf("load finished ... %.3f(s)", t2.Sub(t1).Seconds())

	fmt.Println("\n\nsearch >")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		t := sc.Text()

		if t == "" {
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
